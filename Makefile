.POSIX:
PREFIX = /usr/local

.SUFFIXES:
all:
	go build
install:
	mkdir -p $(PREFIX)/bin
	cp bitter $(PREFIX)/bin
uninstall:
	rm $(PREFIX)/bin/bitter
