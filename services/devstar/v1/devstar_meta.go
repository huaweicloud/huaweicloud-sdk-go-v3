package v1

import (
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/devstar/v1/model"
    "net/http"
)

func GenReqDefForRunTemplateJobV2(request *model.RunTemplateJobV2Request) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v1/jobs/template").
    WithContentType("application/json;charset=UTF-8")



    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("X-Language").
    WithLocationType(def.Header))

    reqDefBuilder.WithBodyJson(request.Body)

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForRunTemplateJobV2() (*model.RunTemplateJobV2Response, *def.HttpResponseDef) {
    resp := new(model.RunTemplateJobV2Response)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForShowJobDetail(request *model.ShowJobDetailRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v1/jobs/{job_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("job_id").
    WithLocationType(def.Path))


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("X-Language").
    WithLocationType(def.Header))


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForShowJobDetail() (*model.ShowJobDetailResponse, *def.HttpResponseDef) {
    resp := new(model.ShowJobDetailResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForListPublishedTemplates(request *model.ListPublishedTemplatesRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v1/templates")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("keyword").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("offset").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("limit").
    WithLocationType(def.Query))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("X-Language").
    WithLocationType(def.Header))


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForListPublishedTemplates() (*model.ListPublishedTemplatesResponse, *def.HttpResponseDef) {
    resp := new(model.ListPublishedTemplatesResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForShowTemplateDetail(request *model.ShowTemplateDetailRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v1/templates/{template_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("template_id").
    WithLocationType(def.Path))


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("X-Language").
    WithLocationType(def.Header))


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForShowTemplateDetail() (*model.ShowTemplateDetailResponse, *def.HttpResponseDef) {
    resp := new(model.ShowTemplateDetailResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

