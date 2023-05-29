if [[ $# -eq 0 ]]; then
    echo 'example usage: sh run.sh 100000'
    exit 0
fi

cd test/
go run gen-test-data.go $1
cd ..
go test -v -args $1
