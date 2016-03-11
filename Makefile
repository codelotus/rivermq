SOURCEDIR=.

SOURCES := $(shell find $(SOURCEDIR) -name '*.go')


BINARY=rivermq

VERSION=1.0.0
BUILD_TIME=`date +%FT%T%z`

.DEFAULT_GOAL: $(BINARY)


$(BINARY): $(SOURCES)
	go build -o ${BINARY} $(SOURCEDIR)/*.go

.PHONY: install
install:
	go install ./...

.PHONY: clean
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: tests
test:
	go test ${SOURCEDIR}
	go test ${SOURCEDIR}/model

build:
	go build -o ${BINARY} $(SOURCEDIR)/*.go

run:
	go build -o ${BINARY} $(SOURCEDIR)/*.go
	$(SOURCEDIR)/$(BINARY)
