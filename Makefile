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


