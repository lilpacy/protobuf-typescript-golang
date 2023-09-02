```sh
# go mod init github.com/lilpacy/protobuf-typescript-golang
go mod tidy
protoc --go_out=. common.proto
go run main.go

# yarn init -y
yarn
protoc --ts_out=common common.proto
npx ts-node main.ts
```
