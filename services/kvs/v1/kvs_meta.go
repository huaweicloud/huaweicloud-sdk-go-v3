package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/kvs/v1/model"
	"net/http"
)

func GenReqDefForCreateTable() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v1/create-table").
		WithResponse(new(model.CreateTableResponse)).
		WithContentType("application/bson")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("StoreName").
		WithJsonTag("store_name").
		WithLocationType(def.Cname))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDescribeTable() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v1/describe-table").
		WithResponse(new(model.DescribeTableResponse)).
		WithContentType("application/bson")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("StoreName").
		WithJsonTag("store_name").
		WithLocationType(def.Cname))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListStore() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v1/list-store").
		WithResponse(new(model.ListStoreResponse)).
		WithContentType("application/bson")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListTable() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v1/list-table").
		WithResponse(new(model.ListTableResponse)).
		WithContentType("application/bson")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("StoreName").
		WithJsonTag("store_name").
		WithLocationType(def.Cname))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForBatchWriteKv() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v1/batch-write-kv").
		WithResponse(new(model.BatchWriteKvResponse)).
		WithContentType("application/bson")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("StoreName").
		WithJsonTag("store_name").
		WithLocationType(def.Cname))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDeleteKv() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v1/delete-kv").
		WithResponse(new(model.DeleteKvResponse)).
		WithContentType("application/bson")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("StoreName").
		WithJsonTag("store_name").
		WithLocationType(def.Cname))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForGetKv() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v1/get-kv").
		WithResponse(new(model.GetKvResponse)).
		WithContentType("application/bson")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("StoreName").
		WithJsonTag("store_name").
		WithLocationType(def.Cname))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForPutKv() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v1/put-kv").
		WithResponse(new(model.PutKvResponse)).
		WithContentType("application/bson")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("StoreName").
		WithJsonTag("store_name").
		WithLocationType(def.Cname))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForScanKv() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v1/scan-kv").
		WithResponse(new(model.ScanKvResponse)).
		WithContentType("application/bson")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("StoreName").
		WithJsonTag("store_name").
		WithLocationType(def.Cname))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForScanSkeyKv() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v1/scan-skey-kv").
		WithResponse(new(model.ScanSkeyKvResponse)).
		WithContentType("application/bson")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("StoreName").
		WithJsonTag("store_name").
		WithLocationType(def.Cname))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForUpdateKv() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v1/update-kv").
		WithResponse(new(model.UpdateKvResponse)).
		WithContentType("application/bson")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("StoreName").
		WithJsonTag("store_name").
		WithLocationType(def.Cname))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}
