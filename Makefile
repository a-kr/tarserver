# Copyright 2014 Alexey Kryuchkov
#
# This file is part of tarserver.
#
# tarserver is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# tarserver is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with tarserver.  If not, see <http://www.gnu.org/licenses/>.
#
BIN=tarserver
APP_PORT=6112
DOCKERTAG=tarserver

.PHONY: $(BIN) test doc

export GOPATH:=$(shell pwd)
export CWD:=$(shell pwd)

$(BIN):
	go install $(BIN)

run: $(BIN)
	bin/$(BIN) $(ARGS)

test:
	go test $(BIN)

doc:
	godoc -http=:6060 -goroot=`pwd`

dockerbuild: $(BIN)
	docker build -t $(DOCKERTAG) .


dockerrun: dockerbuild
	-docker stop $(DOCKERTAG)
	-docker rm $(DOCKERTAG)
	docker run --name $(DOCKERTAG) -d -p $(APP_PORT):$(APP_PORT) $(DOCKERTAG) --listen=":$(APP_PORT)"


