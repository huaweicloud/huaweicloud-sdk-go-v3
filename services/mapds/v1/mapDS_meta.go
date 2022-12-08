package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/mapds/v1/model"
	"net/http"
)

func GenReqDefForCreateCredential() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v1/map/credentials").
		WithResponse(new(model.CreateCredentialResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("XAuthToken").
		WithJsonTag("X-Auth-Token").
		WithLocationType(def.Header))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ContentType").
		WithJsonTag("Content-Type").
		WithLocationType(def.Header))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreateSasToken() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v1/map/sastoken").
		WithResponse(new(model.CreateSasTokenResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("XAuthToken").
		WithJsonTag("X-Auth-Token").
		WithLocationType(def.Header))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ContentType").
		WithJsonTag("Content-Type").
		WithLocationType(def.Header))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDeleteCedential() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v1/map/credentials/{clientid}").
		WithResponse(new(model.DeleteCedentialResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Clientid").
		WithJsonTag("clientid").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("XAuthToken").
		WithJsonTag("X-Auth-Token").
		WithLocationType(def.Header))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ContentType").
		WithJsonTag("Content-Type").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowCredential() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v1/map/credentials").
		WithResponse(new(model.ShowCredentialResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("XAuthToken").
		WithJsonTag("X-Auth-Token").
		WithLocationType(def.Header))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ContentType").
		WithJsonTag("Content-Type").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowMapTile() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v1/map/tile/{z}/{x}/{y}").
		WithResponse(new(model.ShowMapTileResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Z").
		WithJsonTag("z").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("X").
		WithJsonTag("x").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Y").
		WithJsonTag("y").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Authorization").
		WithJsonTag("Authorization").
		WithLocationType(def.Header))

	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("ContentType").
		WithJsonTag("Content-Type").
		WithKindName("string").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}
