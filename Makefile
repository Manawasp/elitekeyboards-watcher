BINARY=elitekeyboards-watcher
VERSION=`git describe --abbrev=0 --tags`
BUILD=`git rev-parse HEAD`

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD} -s -w"

.DEFAULT_GOAL: $(BINARY)

$(BINARY):
	go build ${LDFLAGS} -o ${BINARY}

configure:
	glide install

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

re: clean $(BINARY)

.PHONY: clean install re
