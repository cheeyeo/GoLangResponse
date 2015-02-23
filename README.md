## GoLang http server

Example project of building a simple go lang http server using table cloth and GOM.

The server returns different content types and its an exercise in building and implementing something we take for granted in Rails in a golang app.

It's also an exercise in using a package manager similar to Bundler but for Golang.

## Build

```
go get github.com/mattn/gom

gom install

gom build
```

## Run

```
./httptestapp
```

## Using Makefile

```
make clean

make build

make run
```

