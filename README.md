imgdedup
========

Introduction
------------

Simple program which scans directory with images (jpg, png) and computes their histograms. Next compares their Manhattan distance and shows pairs which distance falls below treshold, indicating potential duplicates.

How to build?
------------

```
$ cd $GOPATH
$ mkdir -p src/github.com/Erdk && cd src/github.com/Erdk
$ git clone https://github.com/Erdk/imgdedup && cd imgdedup
$ go build .
```

How to use?
-----------

```
$ ./imgdedup -p $DIR [-v] [-t $NUM] [-d $DISTANCE_FUNCTION]
```

* -p

  directory with images to check, *required*
* -v

  verbose output
* -t $NUM

  set tolerance to $NUM, distance below this value indicates that images are similar, default: 100000
* -d $DISTANCE_FUNCTION

  set distance function to one of the following:
  
  * manhattan (default),
  * chisquare,
  * correlation,
  * intersection,
  * bhattacharyya

TODO
----

- [ ] More comparison algorithms than Manhatan,
- [x] Instead of relying on hardcoded distance threshold allow to provide one,
- [ ] Tests,
- [ ] Http frontend,
- [ ] Recursively search directory,
- [ ] Allow specyfying multiple directories
