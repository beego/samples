# WebDAV

This sample is about to build a webdav service based on beego.

## Installation

```
cd $GOPATH/src/samples/WebDAV
go get golang.org/x/net/webdav
go install
cp $GOPATH/bin/WebDAV ./
./WebDAV
```

## Usage

* Get a file
```
curl 'http://127.0.0.1:8080/test.txt'
```

* Creating a new foder
```
curl -X MKCOL 'http://127.0.0.1:8080/test/'
```

* Upload a new file
```
curl -T '{localfile}' 'http://127.0.0.1:8080/test/'
```

* Delete foder
```
curl -X DELETE 'http://127.0.0.1:8080/test/'
```
