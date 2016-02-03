#!/bin/bash
godep restore -v
rm -rf Godeps
find . -iname '*.go' | xargs sed -i 's:".*Godeps/_workspace/src/:":g'
godep save -r
