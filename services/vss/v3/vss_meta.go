package v3

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/def"

	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/services/vss/v3/model"
	"net/http"
)

func GenReqDefForCreateDomains() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v3/{project_id}/webscan/domains").
		WithResponse(new(model.CreateDomainsResponse)).
		WithContentType("application/json; charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDeleteDomains() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v3/{project_id}/webscan/domains").
		WithResponse(new(model.DeleteDomainsResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainName").
		WithJsonTag("domain_name").
		WithLocationType(def.Query))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListDomains() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/{project_id}/webscan/domains").
		WithResponse(new(model.ListDomainsResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("AuthStatus").
		WithJsonTag("auth_status").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Offset").
		WithJsonTag("offset").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Limit").
		WithJsonTag("limit").
		WithLocationType(def.Query))

	requestDef := reqDefBuilder.Build()
	return requestDef
}
