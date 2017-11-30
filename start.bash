CMD=$(ls | grep -i -e .*.go$ | grep -v _test | grep -v __)
go run $CMD $1

# go run main.go chunk.go settings.go test.log
