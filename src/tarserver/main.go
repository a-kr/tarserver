package main

/*
Copyright 2014 Alexey Kryuchkov

This file is part of tarserver.

tarserver is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

tarserver is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with tarserver.  If not, see <http://www.gnu.org/licenses/>.
*/

import (
	"archive/tar"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

var (
	Root         = flag.String("root", "/www/", "root directory to serve tar files from")
	BaseLocation = flag.String("base-location", "/contents", "base HTTP location to strip from URI")
	ServeAddr    = flag.String("listen", ":80", "address to listen on")
	Blocksize    = flag.Int64("blocksize", 1024*1024, "file reading blocksize, bytes")
)

const (
	TarSuffix = ".tar/"
)

func getContentTypeByFilename(filename string) string {
	if strings.HasSuffix(filename, ".mp4") {
		return "video/mp4"
	}
	if strings.HasSuffix(filename, ".m3u8") {
		return "application/vnd.apple.mpegurl"
	}
	if strings.HasSuffix(filename, ".m3u") {
		return "audio/mpegurl"
	}
	return "application/octet-stream"
}

func tarHandler(w http.ResponseWriter, r *http.Request) {
	urlPath := strings.TrimPrefix(r.URL.Path, *BaseLocation)
	log.Printf("Requested %s", urlPath)

	uriParts := strings.SplitN(urlPath, TarSuffix, 2)
	if len(uriParts) != 2 {
		http.Error(w, "Request URI must contain "+TarSuffix+" sequence", http.StatusBadRequest)
		return
	}
	tarPath := uriParts[0] + ".tar"
	insideTarPath := uriParts[1]

	fullTarPath := path.Join(*Root, tarPath)

	file, err := os.Open(fullTarPath)
	if err != nil {
		log.Printf("Could not open %s: %s", fullTarPath, err)
		http.Error(w, fmt.Sprintf("Cannot open file: %s", err), http.StatusNotFound)
		return
	}
	defer file.Close()

	tarReader := tar.NewReader(file)
	for {
		hdr, err := tarReader.Next()
		if err == io.EOF {
			// end of tar archive, and we still haven't found our file
			log.Printf("Cannot find %s inside %s", insideTarPath, fullTarPath)
			http.Error(w, fmt.Sprintf("No such file inside the archive: %s", insideTarPath), http.StatusNotFound)
			return
		}
		if err != nil {
			log.Printf("Error while iterating over files in %s: %s", fullTarPath, err)
			http.Error(w, fmt.Sprintf("Error while reading the archive: %s", err), http.StatusInternalServerError)
			return
		}
		if hdr.Name == insideTarPath {
			w.Header().Set("Content-Type", getContentTypeByFilename(insideTarPath))
			w.Header().Set("Content-Length", fmt.Sprintf("%d", hdr.Size))
			var err error
			for err = nil; err == nil; _, err = io.CopyN(w, tarReader, *Blocksize) {
			}
			if err != io.EOF {
				log.Printf("Error while serving %s/%s: %s", fullTarPath, insideTarPath, err)
				return
			}
			log.Printf("Finished serving %s/%s", fullTarPath, insideTarPath)
			return
		}
	}

}

func main() {
	flag.Parse()
	log.Printf("Serving contents of %s on %s...", *Root, *ServeAddr)
	http.HandleFunc("/", tarHandler)
	if err := http.ListenAndServe(*ServeAddr, nil); err != nil {
		log.Fatalf("Cannot serve: %s", err)
	}
}
