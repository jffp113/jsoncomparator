**JSONComparator** is a command line tool that compares two files with an array of json objects.

## Features
- Compare arrays of json objects (wihout nested complex data types)
- Reads both file in parallel to speed up IO
## Table of Contents

- [Build and run Comparator](#build-and-run-comparator)
- [Build docker images](#build-docker-images)
- [Critical thinking](#critical-thinking)

## Build and run Comparator

Requirements:

- **[Build]** [Golang](https://golang.org/) - [go1.18](https://golang.org/dl/)

Compile and install:

```bash
$ go clean . && go build .
```

Run Unit Tests:

```bash
$ go clean . && go test . -v
```

Run comparator tool:

```bash
$ ./comparator <file1> <file2>
```
## Build docker images (NOT COMPLETED)

This builds the docker image jffp113/jsoncomparator:latest using a multi-stage build

```bash
$ docker build -t jffp113/jsoncomparator:latest .
```

Launching containerized tool

```bash
$ docker run jffp113/jsoncomparator:latest file1 file2
```

## Critical thinking

The comparator has an interface that accepts two readers to allow an extra level of flexibility when passing data. For example, a file implements the reader interface, so it can be easily passed into the comparator definition. We can also use it with byte slices, and so on.

Reading a file from a disk is an IO-bound operation. To speed up the process, the function "parseFile" starts a go-routine and returns a channel to store the result of the go-routine. This is done for both files in parallel.

If we have two gigantic arrays with JSON objects, the complexity could explode exponentially. To lower the complexity to O(n + m), or O(n) if both files are equal, each JSON object is hashed and stored in a hashmap for the first file to posteriorly be compared with the second in linear complexity. 