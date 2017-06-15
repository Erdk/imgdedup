imgdedup
========

Introduction
------------

Simple program which scans directory with images (jpg, png) and computes their histograms. Next compares their Manhatan distance and shows pairs which distance falls below treshold, indicating potential duplicates.

How to build?
------------

```
$ cd $IMGDEDUP_SRC_DIR
$ go build .
```

How to use?
-----------

```
$ ./imgdedup $DIR
```

TODO
----

- [ ] More comparison algorithms than Manhatan,
- [ ] Instead of relying on hardcoded distance threshold allow to provide one,
- [ ] Some tests?
