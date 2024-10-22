.PHONY: gen-demo-thrift
gen-demo-thrift:
	@cd demo/demo_thrift && cwgo server --type RPC --module github.com/showyquasar88/go-mall/demo/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift

.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demo_proto && cwgo server -I ../../idl --type RPC --module github.com/showyquasar88/go-mall/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto 
	