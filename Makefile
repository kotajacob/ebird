# ebird
# See LICENSE for copyright and license details.
.POSIX:

include config.mk

all: clean build

build:
	go build -ldflags "-X main.Version=$(VERSION)"

clean:
	rm -f ebird

install: build
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp -f ebird $(DESTDIR)$(PREFIX)/bin
	chmod 755 $(DESTDIR)$(PREFIX)/bin/ebird

uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/ebird

.PHONY: all build clean install uninstall
