tarserver
=========

tarserver can serve files directly from tar archives over HTTP.

How to build
------------

tarserver is written in Go. You need golang 1.2+ to build it.

```
git clone https://github.com/Babazka/tarserver.git
cd tarserver
make
```

How to run
----------

```
bin/tarserver --listen=:8089 --base-location=/files --root=/www
```

Files located in `/www` directory will be served from `http://localhost:8089/files/<filename.tar>/<path inside the archive>`.
