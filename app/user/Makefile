GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

APP="user"
INTERNAL_PROTO_FILES="internal/conf/conf.proto"
API_PROTO_FILES="../../api/${APP}/service/v1/${APP}.proto"
ERR_PROTO_FILES="../../api/${APP}/service/v1/${APP}_error.proto"

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=./internal \
	       --proto_path=../../third_party \
 	       --go_out=paths=source_relative:./internal \
	       $(INTERNAL_PROTO_FILES)

.PHONY: api
# generate api proto
api:
	protoc --proto_path=../../api/${APP} \
	       --proto_path=../../third_party \
 	       --go_out=paths=source_relative:../../api/${APP} \
 	       --go-http_out=paths=source_relative:../../api/${APP} \
 	       --go-grpc_out=paths=source_relative:../../api/${APP} \
	       --openapi_out=fq_schema_naming=true,default_response=false:. \
	       $(API_PROTO_FILES)    && \
	protoc --proto_path=../../api/${APP} \
		   --proto_path=../../third_party \
		   --go_out=paths=source_relative:../../api/${APP} \
		   --go-http_out=paths=source_relative:../../api/${APP} \
		   --go-grpc_out=paths=source_relative:../../api/${APP} \
		   --openapi_out=fq_schema_naming=true,default_response=false:. \
		   $(ERR_PROTO_FILES)

.PHONY: wire
# wire
wire:
	cd cmd && wire