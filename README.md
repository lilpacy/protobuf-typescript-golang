```sh
go mod tidy
protoc --go_out=. common.proto
go run main.go

yarn
protoc --ts_out=common common.proto
npx ts-node main.ts
```
