package v3

import (
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cts/v3/model"
    "net/http"
)

func GenReqDefForCreateTracker(request *model.CreateTrackerRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v3/{project_id}/tracker").
    WithContentType("application/json;charset=UTF-8")




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

func GenRespForCreateTracker() (*model.CreateTrackerResponse, *def.HttpResponseDef) {
    resp := new(model.CreateTrackerResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForDeleteTracker(request *model.DeleteTrackerRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodDelete).
    WithPath("/v3/{project_id}/trackers")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("tracker_name").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("tracker_type").
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

func GenRespForDeleteTracker() (*model.DeleteTrackerResponse, *def.HttpResponseDef) {
    resp := new(model.DeleteTrackerResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForListTraces(request *model.ListTracesRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3/{project_id}/traces")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("trace_type").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("limit").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("from").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("next").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("to").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("tracker_name").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("service_type").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("user").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("resource_id").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("resource_name").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("resource_type").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("trace_id").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("trace_name").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("trace_rating").
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

func GenRespForListTraces() (*model.ListTracesResponse, *def.HttpResponseDef) {
    resp := new(model.ListTracesResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForListTrackers(request *model.ListTrackersRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3/{project_id}/trackers")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("tracker_name").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("tracker_type").
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

func GenRespForListTrackers() (*model.ListTrackersResponse, *def.HttpResponseDef) {
    resp := new(model.ListTrackersResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForUpdateTracker(request *model.UpdateTrackerRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPut).
    WithPath("/v3/{project_id}/tracker").
    WithContentType("application/json;charset=UTF-8")




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

func GenRespForUpdateTracker() (*model.UpdateTrackerResponse, *def.HttpResponseDef) {
    resp := new(model.UpdateTrackerResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

