SOURCEDIR=.

SOURCES := $(shell find $(SOURCEDIR) -name '*.go')


BINARY=rivermq

VERSION=1.0.0
BUILD_TIME=`date +%FT%T%z`

.DEFAULT_GOAL := integration


.PHONY: install
install:
	go install $(SOURCEDIR)/...

.PHONY: clean
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

test: $(SOURCES)
	go test ${SOURCEDIR}
	go test ${SOURCEDIR}/route
	go test ${SOURCEDIR}/model

.PHONY: integration
integration:
	go test ${SOURCEDIR}
	go test ${SOURCEDIR}/route -tags integration
	go test ${SOURCEDIR}/model -tags integration


build: $(SOURCES)
	go build -o ${BINARY} $(SOURCEDIR)/*.go

run:
	go build -o ${BINARY} $(SOURCEDIR)/*.go
	$(SOURCEDIR)/$(BINARY)
