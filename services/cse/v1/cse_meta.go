package v1

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cse/v1/model"
	"net/http"
)

func GenReqDefForCreateEngine() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v2/{project_id}/enginemgr/engines").
		WithResponse(new(model.CreateEngineResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("XEnterpriseProjectID").
		WithJsonTag("X-Enterprise-Project-ID").
		WithLocationType(def.Header))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDeleteEngine() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v2/{project_id}/enginemgr/engines/{engine_id}").
		WithResponse(new(model.DeleteEngineResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("EngineId").
		WithJsonTag("engine_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("XEnterpriseProjectID").
		WithJsonTag("X-Enterprise-Project-ID").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDownloadKie() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v1/{project_id}/kie/download").
		WithResponse(new(model.DownloadKieResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Label").
		WithJsonTag("label").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Match").
		WithJsonTag("match").
		WithLocationType(def.Query))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("XEnterpriseProjectID").
		WithJsonTag("X-Enterprise-Project-ID").
		WithLocationType(def.Header))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("XEngineId").
		WithJsonTag("x-engine-id").
		WithLocationType(def.Header))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListEngines() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/enginemgr/engines").
		WithResponse(new(model.ListEnginesResponse)).
		WithContentType("application/json")

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

func GenReqDefForListFlavors() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/enginemgr/flavors").
		WithResponse(new(model.ListFlavorsResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("SpecType").
		WithJsonTag("spec_type").
		WithLocationType(def.Query))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowEngine() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/enginemgr/engines/{engine_id}").
		WithResponse(new(model.ShowEngineResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("EngineId").
		WithJsonTag("engine_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("XEnterpriseProjectID").
		WithJsonTag("X-Enterprise-Project-ID").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowEngineJob() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v2/{project_id}/enginemgr/engines/{engine_id}/jobs/{job_id}").
		WithResponse(new(model.ShowEngineJobResponse)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("EngineId").
		WithJsonTag("engine_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("JobId").
		WithJsonTag("job_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("XEnterpriseProjectID").
		WithJsonTag("X-Enterprise-Project-ID").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForUploadKie() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v1/{project_id}/kie/file").
		WithResponse(new(model.UploadKieResponse)).
		WithContentType("multipart/form-data")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Override").
		WithJsonTag("override").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Label").
		WithJsonTag("label").
		WithLocationType(def.Query))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("XEnterpriseProjectID").
		WithJsonTag("X-Enterprise-Project-ID").
		WithLocationType(def.Header))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("XEngineId").
		WithJsonTag("x-engine-id").
		WithLocationType(def.Header))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	requestDef := reqDefBuilder.Build()
	return requestDef
}
