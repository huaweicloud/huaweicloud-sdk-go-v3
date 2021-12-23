package v1

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/def"

	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/services/sfsturbo/v1/model"
	"net/http"
)

func GenReqDefForBatchAddSharedTags() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v1/{project_id}/sfs-turbo/{share_id}/tags/action").
		WithResponse(new(model.BatchAddSharedTagsResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ShareId").
		WithJsonTag("share_id").
		WithLocationType(def.Path))

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

func GenReqDefForChangeSecurityGroup() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v1/{project_id}/sfs-turbo/shares/{share_id}/action").
		WithResponse(new(model.ChangeSecurityGroupResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ShareId").
		WithJsonTag("share_id").
		WithLocationType(def.Path))

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

func GenReqDefForCreateShare() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v1/{project_id}/sfs-turbo/shares").
		WithResponse(new(model.CreateShareResponse)).
		WithContentType("application/json;charset=UTF-8")

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

func GenReqDefForCreateSharedTag() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v1/{project_id}/sfs-turbo/{share_id}/tags").
		WithResponse(new(model.CreateSharedTagResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ShareId").
		WithJsonTag("share_id").
		WithLocationType(def.Path))

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

func GenReqDefForDeleteShare() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v1/{project_id}/sfs-turbo/shares/{share_id}").
		WithResponse(new(model.DeleteShareResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ShareId").
		WithJsonTag("share_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ContentType").
		WithJsonTag("Content-Type").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDeleteSharedTag() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v1/{project_id}/sfs-turbo/{share_id}/tags/{key}").
		WithResponse(new(model.DeleteSharedTagResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ShareId").
		WithJsonTag("share_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Key").
		WithJsonTag("key").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ContentType").
		WithJsonTag("Content-Type").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForExpandShare() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v1/{project_id}/sfs-turbo/shares/{share_id}/action").
		WithResponse(new(model.ExpandShareResponse)).
		WithContentType("application/json;charset=UTF-8")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ShareId").
		WithJsonTag("share_id").
		WithLocationType(def.Path))

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

func GenReqDefForListSharedTags() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v1/{project_id}/sfs-turbo/tags").
		WithResponse(new(model.ListSharedTagsResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ContentType").
		WithJsonTag("Content-Type").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListShares() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v1/{project_id}/sfs-turbo/shares/detail").
		WithResponse(new(model.ListSharesResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Limit").
		WithJsonTag("limit").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Offset").
		WithJsonTag("offset").
		WithLocationType(def.Query))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ContentType").
		WithJsonTag("Content-Type").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowShare() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v1/{project_id}/sfs-turbo/shares/{share_id}").
		WithResponse(new(model.ShowShareResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ShareId").
		WithJsonTag("share_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ContentType").
		WithJsonTag("Content-Type").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowSharedTags() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v1/{project_id}/sfs-turbo/{share_id}/tags").
		WithResponse(new(model.ShowSharedTagsResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ShareId").
		WithJsonTag("share_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ContentType").
		WithJsonTag("Content-Type").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}
