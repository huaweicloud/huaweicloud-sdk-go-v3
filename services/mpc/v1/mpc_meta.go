package v1

import (
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/mpc/v1/model"
    "net/http"
)

func GenReqDefForCreateAnimatedGraphicsTask(request *model.CreateAnimatedGraphicsTaskRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v1/{project_id}/animated-graphics").
    WithContentType("application/json")




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

func GenRespForCreateAnimatedGraphicsTask() (*model.CreateAnimatedGraphicsTaskResponse, *def.HttpResponseDef) {
    resp := new(model.CreateAnimatedGraphicsTaskResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForDeleteAnimatedGraphicsTask(request *model.DeleteAnimatedGraphicsTaskRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodDelete).
    WithPath("/v1/{project_id}/animated-graphics")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("task_id").
    WithLocationType(def.Query))



    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForDeleteAnimatedGraphicsTask() (*model.DeleteAnimatedGraphicsTaskResponse, *def.HttpResponseDef) {
    resp := new(model.DeleteAnimatedGraphicsTaskResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForListAnimatedGraphicsTask(request *model.ListAnimatedGraphicsTaskRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v1/{project_id}/animated-graphics")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("task_id").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("status").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("start_time").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("end_time").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("page").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("size").
    WithLocationType(def.Query))



    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForListAnimatedGraphicsTask() (*model.ListAnimatedGraphicsTaskResponse, *def.HttpResponseDef) {
    resp := new(model.ListAnimatedGraphicsTaskResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForCreateEncryptTask(request *model.CreateEncryptTaskRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v1/{project_id}/encryptions").
    WithContentType("application/json")




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

func GenRespForCreateEncryptTask() (*model.CreateEncryptTaskResponse, *def.HttpResponseDef) {
    resp := new(model.CreateEncryptTaskResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForDeleteEncryptTask(request *model.DeleteEncryptTaskRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodDelete).
    WithPath("/v1/{project_id}/encryptions")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("task_id").
    WithLocationType(def.Query))



    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForDeleteEncryptTask() (*model.DeleteEncryptTaskResponse, *def.HttpResponseDef) {
    resp := new(model.DeleteEncryptTaskResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForListEncryptTask(request *model.ListEncryptTaskRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v1/{project_id}/encryptions")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("task_id").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("status").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("start_time").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("end_time").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("page").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("size").
    WithLocationType(def.Query))



    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForListEncryptTask() (*model.ListEncryptTaskResponse, *def.HttpResponseDef) {
    resp := new(model.ListEncryptTaskResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForCreateExtractTask(request *model.CreateExtractTaskRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v1/{project_id}/extract-metadata").
    WithContentType("application/json")




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

func GenRespForCreateExtractTask() (*model.CreateExtractTaskResponse, *def.HttpResponseDef) {
    resp := new(model.CreateExtractTaskResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForDeleteExtractTask(request *model.DeleteExtractTaskRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodDelete).
    WithPath("/v1/{project_id}/extract-metadata")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("task_id").
    WithLocationType(def.Query))



    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForDeleteExtractTask() (*model.DeleteExtractTaskResponse, *def.HttpResponseDef) {
    resp := new(model.DeleteExtractTaskResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForListExtractTask(request *model.ListExtractTaskRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v1/{project_id}/extract-metadata")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("task_id").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("status").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("start_time").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("end_time").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("page").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("size").
    WithLocationType(def.Query))



    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForListExtractTask() (*model.ListExtractTaskResponse, *def.HttpResponseDef) {
    resp := new(model.ListExtractTaskResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForShowJobStatus(request *model.ShowJobStatusRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v1/{tenantId}/jobs/{jobId}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("tenantId").
    WithLocationType(def.Path))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("jobId").
    WithLocationType(def.Path))




    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForShowJobStatus() (*model.ShowJobStatusResponse, *def.HttpResponseDef) {
    resp := new(model.ShowJobStatusResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForListAllBuckets(request *model.ListAllBucketsRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v1/{project_id}/buckets")





    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForListAllBuckets() (*model.ListAllBucketsResponse, *def.HttpResponseDef) {
    resp := new(model.ListAllBucketsResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForUpdateBucketAuthorized(request *model.UpdateBucketAuthorizedRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPut).
    WithPath("/v1/{project_id}/authority").
    WithContentType("application/json")




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

func GenRespForUpdateBucketAuthorized() (*model.UpdateBucketAuthorizedResponse, *def.HttpResponseDef) {
    resp := new(model.UpdateBucketAuthorizedResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForListAllObsObjList(request *model.ListAllObsObjListRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v1.0-ext/{project_id}/objects")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("bucket").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("prefix").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("type").
    WithLocationType(def.Query))



    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForListAllObsObjList() (*model.ListAllObsObjListResponse, *def.HttpResponseDef) {
    resp := new(model.ListAllObsObjListResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForCreateQualityEnhanceTemplate(request *model.CreateQualityEnhanceTemplateRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v1/{project_id}/template/qualityenhance").
    WithContentType("application/json")




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

func GenRespForCreateQualityEnhanceTemplate() (*model.CreateQualityEnhanceTemplateResponse, *def.HttpResponseDef) {
    resp := new(model.CreateQualityEnhanceTemplateResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForDeleteQualityEnhanceTemplate(request *model.DeleteQualityEnhanceTemplateRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodDelete).
    WithPath("/v1/{project_id}/template/qualityenhance")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("template_id").
    WithLocationType(def.Query))



    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForDeleteQualityEnhanceTemplate() (*model.DeleteQualityEnhanceTemplateResponse, *def.HttpResponseDef) {
    resp := new(model.DeleteQualityEnhanceTemplateResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForListQualityEnhanceDefaultTemplate(request *model.ListQualityEnhanceDefaultTemplateRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v1/{project_id}/template/qualityenhance/default")





    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForListQualityEnhanceDefaultTemplate() (*model.ListQualityEnhanceDefaultTemplateResponse, *def.HttpResponseDef) {
    resp := new(model.ListQualityEnhanceDefaultTemplateResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForListQualityEnhanceTemplate(request *model.ListQualityEnhanceTemplateRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v1/{project_id}/template/qualityenhance")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("template_id").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("page").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("size").
    WithLocationType(def.Query))



    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForListQualityEnhanceTemplate() (*model.ListQualityEnhanceTemplateResponse, *def.HttpResponseDef) {
    resp := new(model.ListQualityEnhanceTemplateResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForUpdateQualityEnhanceTemplate(request *model.UpdateQualityEnhanceTemplateRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPut).
    WithPath("/v1/{project_id}/template/qualityenhance").
    WithContentType("application/json")




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

func GenRespForUpdateQualityEnhanceTemplate() (*model.UpdateQualityEnhanceTemplateResponse, *def.HttpResponseDef) {
    resp := new(model.UpdateQualityEnhanceTemplateResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForListTranscodeDetail(request *model.ListTranscodeDetailRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v1/{project_id}/transcodings/detail")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("task_id").
    WithLocationType(def.Query))



    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForListTranscodeDetail() (*model.ListTranscodeDetailResponse, *def.HttpResponseDef) {
    resp := new(model.ListTranscodeDetailResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForCancelRemuxTask(request *model.CancelRemuxTaskRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodDelete).
    WithPath("/v1/{project_id}/remux")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("task_id").
    WithLocationType(def.Query))



    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForCancelRemuxTask() (*model.CancelRemuxTaskResponse, *def.HttpResponseDef) {
    resp := new(model.CancelRemuxTaskResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForCreateRemuxTask(request *model.CreateRemuxTaskRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v1/{project_id}/remux").
    WithContentType("application/json")




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

func GenRespForCreateRemuxTask() (*model.CreateRemuxTaskResponse, *def.HttpResponseDef) {
    resp := new(model.CreateRemuxTaskResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForCreateRetryRemuxTask(request *model.CreateRetryRemuxTaskRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPut).
    WithPath("/v1/{project_id}/remux").
    WithContentType("application/json")




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

func GenRespForCreateRetryRemuxTask() (*model.CreateRetryRemuxTaskResponse, *def.HttpResponseDef) {
    resp := new(model.CreateRetryRemuxTaskResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForDeleteRemuxTask(request *model.DeleteRemuxTaskRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodDelete).
    WithPath("/v1/{project_id}/remux/task")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("task_id").
    WithLocationType(def.Query))



    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForDeleteRemuxTask() (*model.DeleteRemuxTaskResponse, *def.HttpResponseDef) {
    resp := new(model.DeleteRemuxTaskResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForListRemuxTask(request *model.ListRemuxTaskRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v1/{project_id}/remux")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("task_id").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("status").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("start_time").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("end_time").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("input_bucket").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("input_object").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("page").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("size").
    WithLocationType(def.Query))



    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForListRemuxTask() (*model.ListRemuxTaskResponse, *def.HttpResponseDef) {
    resp := new(model.ListRemuxTaskResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForListNotifyEvent(request *model.ListNotifyEventRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v1/{project_id}/notification/event")





    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForListNotifyEvent() (*model.ListNotifyEventResponse, *def.HttpResponseDef) {
    resp := new(model.ListNotifyEventResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForListNotifySmnTopicConfig(request *model.ListNotifySmnTopicConfigRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v1/{project_id}/notification")





    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForListNotifySmnTopicConfig() (*model.ListNotifySmnTopicConfigResponse, *def.HttpResponseDef) {
    resp := new(model.ListNotifySmnTopicConfigResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForListNotifyTemplate(request *model.ListNotifyTemplateRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v1/{project_id}/notification/template")





    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForListNotifyTemplate() (*model.ListNotifyTemplateResponse, *def.HttpResponseDef) {
    resp := new(model.ListNotifyTemplateResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForNotifySmnTopicConfig(request *model.NotifySmnTopicConfigRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPut).
    WithPath("/v1/{project_id}/notification").
    WithContentType("application/json")




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

func GenRespForNotifySmnTopicConfig() (*model.NotifySmnTopicConfigResponse, *def.HttpResponseDef) {
    resp := new(model.NotifySmnTopicConfigResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForListStatSummary(request *model.ListStatSummaryRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v1/{project_id}/transcodings/summaries")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("start_time").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("end_time").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("stat_type").
    WithLocationType(def.Query))



    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForListStatSummary() (*model.ListStatSummaryResponse, *def.HttpResponseDef) {
    resp := new(model.ListStatSummaryResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForListSummary(request *model.ListSummaryRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v1/{project_id}/transcodings/summary")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("range").
    WithLocationType(def.Query))



    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForListSummary() (*model.ListSummaryResponse, *def.HttpResponseDef) {
    resp := new(model.ListSummaryResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForCreateTemplateGroup(request *model.CreateTemplateGroupRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v1/{project_id}/template_group/transcodings").
    WithContentType("application/json")




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

func GenRespForCreateTemplateGroup() (*model.CreateTemplateGroupResponse, *def.HttpResponseDef) {
    resp := new(model.CreateTemplateGroupResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForDeleteTemplateGroup(request *model.DeleteTemplateGroupRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodDelete).
    WithPath("/v1/{project_id}/template_group/transcodings")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("group_id").
    WithLocationType(def.Query))



    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForDeleteTemplateGroup() (*model.DeleteTemplateGroupResponse, *def.HttpResponseDef) {
    resp := new(model.DeleteTemplateGroupResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForListTemplateGroup(request *model.ListTemplateGroupRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v1/{project_id}/template_group/transcodings")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("group_id").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("group_name").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("page").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("size").
    WithLocationType(def.Query))



    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForListTemplateGroup() (*model.ListTemplateGroupResponse, *def.HttpResponseDef) {
    resp := new(model.ListTemplateGroupResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForUpdateTemplateGroup(request *model.UpdateTemplateGroupRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPut).
    WithPath("/v1/{project_id}/template_group/transcodings").
    WithContentType("application/json")




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

func GenRespForUpdateTemplateGroup() (*model.UpdateTemplateGroupResponse, *def.HttpResponseDef) {
    resp := new(model.UpdateTemplateGroupResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForUpdateTenantStatus(request *model.UpdateTenantStatusRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPut).
    WithPath("/v1/services/status").
    WithContentType("application/json")




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

func GenRespForUpdateTenantStatus() (*model.UpdateTenantStatusResponse, *def.HttpResponseDef) {
    resp := new(model.UpdateTenantStatusResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForShowTenantInfo(request *model.ShowTenantInfoRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v1/{project_id}/tenant_info")





    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForShowTenantInfo() (*model.ShowTenantInfoResponse, *def.HttpResponseDef) {
    resp := new(model.ShowTenantInfoResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForUpdateTenantInfo(request *model.UpdateTenantInfoRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPut).
    WithPath("/v1/{project_id}/tenant_info").
    WithContentType("application/json")




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

func GenRespForUpdateTenantInfo() (*model.UpdateTenantInfoResponse, *def.HttpResponseDef) {
    resp := new(model.UpdateTenantInfoResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForCreateThumbnailsTask(request *model.CreateThumbnailsTaskRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v1/{project_id}/thumbnails").
    WithContentType("application/json")




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

func GenRespForCreateThumbnailsTask() (*model.CreateThumbnailsTaskResponse, *def.HttpResponseDef) {
    resp := new(model.CreateThumbnailsTaskResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForDeleteThumbnailsTask(request *model.DeleteThumbnailsTaskRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodDelete).
    WithPath("/v1/{project_id}/thumbnails")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("task_id").
    WithLocationType(def.Query))



    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForDeleteThumbnailsTask() (*model.DeleteThumbnailsTaskResponse, *def.HttpResponseDef) {
    resp := new(model.DeleteThumbnailsTaskResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForListThumbnailsTask(request *model.ListThumbnailsTaskRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v1/{project_id}/thumbnails")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("task_id").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("status").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("start_time").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("end_time").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("page").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("size").
    WithLocationType(def.Query))



    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForListThumbnailsTask() (*model.ListThumbnailsTaskResponse, *def.HttpResponseDef) {
    resp := new(model.ListThumbnailsTaskResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForCreateTranscodingTask(request *model.CreateTranscodingTaskRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v1/{project_id}/transcodings").
    WithContentType("application/json")




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

func GenRespForCreateTranscodingTask() (*model.CreateTranscodingTaskResponse, *def.HttpResponseDef) {
    resp := new(model.CreateTranscodingTaskResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForDeleteTranscodingTask(request *model.DeleteTranscodingTaskRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodDelete).
    WithPath("/v1/{project_id}/transcodings")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("task_id").
    WithLocationType(def.Query))



    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForDeleteTranscodingTask() (*model.DeleteTranscodingTaskResponse, *def.HttpResponseDef) {
    resp := new(model.DeleteTranscodingTaskResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForDeleteTranscodingTaskByConsole(request *model.DeleteTranscodingTaskByConsoleRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodDelete).
    WithPath("/v1/{project_id}/transcodings/task")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("task_id").
    WithLocationType(def.Query))



    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForDeleteTranscodingTaskByConsole() (*model.DeleteTranscodingTaskByConsoleResponse, *def.HttpResponseDef) {
    resp := new(model.DeleteTranscodingTaskByConsoleResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForListTranscodingTask(request *model.ListTranscodingTaskRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v1/{project_id}/transcodings")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("task_id").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("status").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("start_time").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("end_time").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("page").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("size").
    WithLocationType(def.Query))



    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForListTranscodingTask() (*model.ListTranscodingTaskResponse, *def.HttpResponseDef) {
    resp := new(model.ListTranscodingTaskResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForCreateTransTemplate(request *model.CreateTransTemplateRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v1/{project_id}/template/transcodings").
    WithContentType("application/json")




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

func GenRespForCreateTransTemplate() (*model.CreateTransTemplateResponse, *def.HttpResponseDef) {
    resp := new(model.CreateTransTemplateResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForDeleteTemplate(request *model.DeleteTemplateRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodDelete).
    WithPath("/v1/{project_id}/template/transcodings")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("template_id").
    WithLocationType(def.Query))



    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForDeleteTemplate() (*model.DeleteTemplateResponse, *def.HttpResponseDef) {
    resp := new(model.DeleteTemplateResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForListTemplate(request *model.ListTemplateRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v1/{project_id}/template/transcodings")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("template_id").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("page").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("size").
    WithLocationType(def.Query))



    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForListTemplate() (*model.ListTemplateResponse, *def.HttpResponseDef) {
    resp := new(model.ListTemplateResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForUpdateTransTemplate(request *model.UpdateTransTemplateRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPut).
    WithPath("/v1/{project_id}/template/transcodings").
    WithContentType("application/json")




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

func GenRespForUpdateTransTemplate() (*model.UpdateTransTemplateResponse, *def.HttpResponseDef) {
    resp := new(model.UpdateTransTemplateResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForCreateWatermarkTemplate(request *model.CreateWatermarkTemplateRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v1/{project_id}/template/watermark").
    WithContentType("application/json")




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

func GenRespForCreateWatermarkTemplate() (*model.CreateWatermarkTemplateResponse, *def.HttpResponseDef) {
    resp := new(model.CreateWatermarkTemplateResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForDeleteWatermarkTemplate(request *model.DeleteWatermarkTemplateRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodDelete).
    WithPath("/v1/{project_id}/template/watermark")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("template_id").
    WithLocationType(def.Query))



    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForDeleteWatermarkTemplate() (*model.DeleteWatermarkTemplateResponse, *def.HttpResponseDef) {
    resp := new(model.DeleteWatermarkTemplateResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForListWatermarkTemplate(request *model.ListWatermarkTemplateRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v1/{project_id}/template/watermark")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("template_id").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("page").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("size").
    WithLocationType(def.Query))



    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForListWatermarkTemplate() (*model.ListWatermarkTemplateResponse, *def.HttpResponseDef) {
    resp := new(model.ListWatermarkTemplateResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForUpdateWatermarkTemplate(request *model.UpdateWatermarkTemplateRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPut).
    WithPath("/v1/{project_id}/template/watermark").
    WithContentType("application/json")




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

func GenRespForUpdateWatermarkTemplate() (*model.UpdateWatermarkTemplateResponse, *def.HttpResponseDef) {
    resp := new(model.UpdateWatermarkTemplateResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

