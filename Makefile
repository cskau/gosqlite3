include $(GOROOT)/src/Make.inc

TARG=sqlite3

CGOFILES=\
	sqlite3.go

CGO_LDFLAGS=-lsqlite3

include $(GOROOT)/src/Make.pkg
