# Packages to install
Protoc

protoc-gen-go and protoc-gen-go-

grpc-web

# Install protoc
```
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go

go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

# Generate Go and Javascript code
```
protoc ./proto/service.proto --js_out=import_style=commonjs,binary:. --grpc-web_out=import_style=commonjs,mode=grpcwebtext:. --go-grpc_out=./proto --go_out=./proto
```

# Install Javascript dependencies
```
npm install -g npx
npm install grpc-web
npm install google-protobuf
```

#  Build frontend in webapp/hwmonitor
```
npm run build
```

# Run Go server
```
go run *.go
```