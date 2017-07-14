#!/bin/bash
check()
{
    [[ $1 -ne 0 ]] && exit 1
}

MODELS=models
rm -rf $MODELS/*.xo.go 
CMDLINE="xo --single-file mysql://emindsoftwarecenter:1@localhost/emind_software_center --suffix .xo.go -o $MODELS"
$CMDLINE
check $?

QUERY="$CMDLINE -a -N -M -B -T"

$QUERY ScProduct << ENDSQL
SELECT
    *
    FROM sc_product
ENDSQL
check $?

$QUERY ScCategory << ENDSQL
SELECT
    *
    FROM sc_category
ENDSQL
check $?

pushd models &>/dev/null
rm -f *_easyjson.go
sed 's/`json:\(.*\)`/`json:\1 form:\1`/g' -i *.go
sed 's/`json:\(.*\)" /`json:\1,omitempty" /g' -i *.go
easyjson -all *.go
popd &>/dev/null
gofmt -w models

