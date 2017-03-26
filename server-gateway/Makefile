GOARCH  = amd64
GOOS    = $(OS)

ifeq ($(GOOS),)
  ifeq ($(shell  uname -s), Darwin)
    GOOS	= darwin
  else
    GOOS	= linux
  endif
endif

APP			= Gateway
MAIN		= main.go
DIST_DIR	= dist
VENDOR_DIR	= vendor
BIN			= ${DIST_DIR}/${GOOS}-${GOARCH}/${APP}

RM		= rm -rf
GOCMD	= go
GOGET	= $(GOCMD) get -u -v
GOBUILD		= GOOS=$(GOOS) GOARCH=$(GOARCH) $(GOCMD) build
GOINSTALL	= $(GOCMD) install
GODEP	= glide
GOTEST	= ginkgo
GOTEST_OPT	= -r -p -v -cover

# build flags
IMPORTPATH	= "github.com/a-trium/gipeline/server-gateway"
VERSION		= "0.0.1"
GITSHA		= $(shell git rev-parse HEAD)
BUILDTIME	= $(shell TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ')
LDFLAGS		= -ldflags="-X ${IMPORTPATH}/config.GitHash=${GITSHA} -X ${IMPORTPATH}/config.BuildTime=${BUILDTIME} -X ${IMPORTPATH}/config.Version=${VERSION}"

# build commands
.PHONY: all install build build-all test run run-cont clean fclean

all: fclean install lint build-all test

lint:
	$(GOINSTALL) . # https://github.com/alecthomas/gometalinter/issues/9
	gometalinter --vendor --errors

install:
	$(GOGET) github.com/Masterminds/glide # for dep
	$(GOGET) github.com/onsi/ginkgo/ginkgo # for test
	$(GOGET) github.com/franciscocpg/gox # for cross compile
	$(GOGET) github.com/githubnemo/CompileDaemon # for reload
	$(GOGET) github.com/alecthomas/gometalinter # for lint
	gometalinter --install --update --force
	$(GODEP) install

build: clean
	$(GOBUILD) $(LDFLAGS) -i -o ${BIN} -v ${MAIN}

build-all: clean
	gox -verbose \
	$(LDFLAGS) \
	-os="linux darwin windows freebsd openbsd netbsd" \
	-arch="amd64 386 armv5 armv6 armv7 arm64" \
	-osarch="!darwin/arm64" \
	-output="${DIST_DIR}/{{.OS}}-{{.Arch}}/${APP}" .

test:
	$(GOTEST) $(GOTEST_OPT)

test-cont:
	$(GOTEST) watch $(GOTEST_OPT)

run:
	${BIN}

run-cont: build
	CompileDaemon -exclude-dir="${VENDOR_DIR}" -build="make build" -command="make run" -graceful-kill=true

clean:
	find . -type f -name '*.coverprofile' -exec rm {} +
	$(RM) ${DIST_DIR}
	$(RM) vendor/$(IMPORTPATH)

fclean: clean
	$(RM) vendor
