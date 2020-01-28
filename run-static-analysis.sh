#!/bin/bash

go get -u github.com/tkrajina/golongfuncs/...

PATH="$PATH:$(go env GOPATH)/bin"

function report_violations () {
	maximum=$1
	violation_type=$2
	violations=$3

	if [[ ! -z "$violations" ]]
	then
		echo "$violation_type violations found (maximum is $maximum):"
		echo "$violations"
	fi
}

MAXIMUM_FUNCTION_LENGTH="20"
MAXIMUM_FILE_LENGTH="100"
MAXIMUM_COMPLEXITY="5"

FUNCTION_LENGTH_VIOLATIONS=$(golongfuncs -top 100  -min-lines 1 +total_lines -threshold $MAXIMUM_FUNCTION_LENGTH)
FILE_LENGTH_VIOLATIONS=$(find . -name '*.go' | xargs wc -l | egrep -v "test|total" | sort | awk -v max=$MAXIMUM_FILE_LENGTH '$1>max')
COMPLEXITY_VIOLATIONS=$(golongfuncs -top 100 -min-lines 1 +complexity -threshold $MAXIMUM_COMPLEXITY)

report_violations "$MAXIMUM_FUNCTION_LENGTH" "Function length" "$FUNCTION_LENGTH_VIOLATIONS"
report_violations "$MAXIMUM_FILE_LENGTH" "File length" "$FILE_LENGTH_VIOLATIONS"
report_violations "$MAXIMUM_COMPLEXITY" "Complexity" "$COMPLEXITY_VIOLATIONS"

if [[ ! -z "$FUNCTION_LENGTH_VIOLATIONS" || ! -z "$COMPLEXITY_VIOLATIONS" || ! -z "$FILE_LENGTH_VIOLATIONS" ]]
then
	exit 1
fi
