#!/bin/sh
./privateer run --binaries-path . --config /.privateer/config.yml > /dev/null 2>&1

for file in evaluation_results/**/*.log; do echo $file; cat $file; echo; done
