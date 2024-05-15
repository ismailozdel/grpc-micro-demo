module github.com/ismailozdel/micro/invoice

go 1.22.3

replace github.com/ismailozdel/micro/common/proto v0.0.0 => ../common/proto

require (
	github.com/ismailozdel/micro/common/proto v0.0.0
	github.com/kelseyhightower/envconfig v1.4.0
	google.golang.org/grpc v1.64.0
)

require (
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240513163218-0867130af1f8 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
)
