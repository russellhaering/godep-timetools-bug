# Godep Timetools Bug

This repository demonstrates an issue I'm having with Godep.

To reproduce the issue:

1. `go get github.com/russellhaering/godep-timetools-bug`
2. `cd $GOPATH/src/github.com/russellhaering/godep-timetools-bug`
3. `./reproduce.sh`

After step 3 the repository should be clean. Now note that the [timetools import of bson](Godeps/_workspace/src/github.com/mailgun/timetools/rfc2822time.go) has not been rewritten. Also note that `gopkg.in/mgo.v2/bson`
has not been vendored.
