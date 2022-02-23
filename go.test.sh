#!/usr/bin/env zsh

set -e
local DELETE_PRIOR_COVERAGE_DATA=
local SET_DEBUG=

# Comment this line to append coverage changes instead of
# clearing the coverage.txt file each time.
DELETE_PRIOR_COVERAGE_DATA='true'

_runtest() {
	# unique filename for each run - allows parallel execution
	local tmpfile="profile$(date '+%s%N').out"
    go test -race -coverprofile=$tmpfile -covermode=atomic "$1"
    if [ -f $tmpfile ]; then
        cat $tmpfile >> coverage.txt
        rm -rf $tmpfile
    fi
}

if [ -z $DELETE_PRIOR_COVERAGE_DATA ]; then
	echo "--------------------------------------------------------------------------" >> coverage.txt
else
	: >| coverage.txt
fi

echo "--------------------------- $(date +'%b %d %Y: %R') ---------------------------" >> coverage.txt

for d in $(go list ./... | grep -v vendor); do
	_runtest "$d"
done
