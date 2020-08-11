package v2

import (
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eip/v2/model"
    "net/http"
)

func GenReqDefForAddPublicipsIntoSharedBandwidth(request *model.AddPublicipsIntoSharedBandwidthRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v2.0/{project_id}/bandwidths/{bandwidth_id}/insert").
    WithContentType("application/json;charset=UTF-8")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("bandwidth_id").
    WithLocationType(def.Path))



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

func GenRespForAddPublicipsIntoSharedBandwidth() (*model.AddPublicipsIntoSharedBandwidthResponse, *def.HttpResponseDef) {
    resp := new(model.AddPublicipsIntoSharedBandwidthResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForBatchCreateSharedBandwidths(request *model.BatchCreateSharedBandwidthsRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v2.0/{project_id}/batch-bandwidths").
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

func GenRespForBatchCreateSharedBandwidths() (*model.BatchCreateSharedBandwidthsResponse, *def.HttpResponseDef) {
    resp := new(model.BatchCreateSharedBandwidthsResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForCreateSharedBandwidth(request *model.CreateSharedBandwidthRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v2.0/{project_id}/bandwidths").
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

func GenRespForCreateSharedBandwidth() (*model.CreateSharedBandwidthResponse, *def.HttpResponseDef) {
    resp := new(model.CreateSharedBandwidthResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForDeleteSharedBandwidth(request *model.DeleteSharedBandwidthRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodDelete).
    WithPath("/v2.0/{project_id}/bandwidths/{bandwidth_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("bandwidth_id").
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

func GenRespForDeleteSharedBandwidth() (*model.DeleteSharedBandwidthResponse, *def.HttpResponseDef) {
    resp := new(model.DeleteSharedBandwidthResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForListBandwidths(request *model.ListBandwidthsRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v1/{project_id}/bandwidths")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("marker").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("limit").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("enterprise_project_id").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("share_type").
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

func GenRespForListBandwidths() (*model.ListBandwidthsResponse, *def.HttpResponseDef) {
    resp := new(model.ListBandwidthsResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForListQuotas(request *model.ListQuotasRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v1/{project_id}/quotas")


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

func GenRespForListQuotas() (*model.ListQuotasResponse, *def.HttpResponseDef) {
    resp := new(model.ListQuotasResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForRemovePublicipsFromSharedBandwidth(request *model.RemovePublicipsFromSharedBandwidthRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v2.0/{project_id}/bandwidths/{bandwidth_id}/remove").
    WithContentType("application/json;charset=UTF-8")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("bandwidth_id").
    WithLocationType(def.Path))



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

func GenRespForRemovePublicipsFromSharedBandwidth() (*model.RemovePublicipsFromSharedBandwidthResponse, *def.HttpResponseDef) {
    resp := new(model.RemovePublicipsFromSharedBandwidthResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForShowBandwidth(request *model.ShowBandwidthRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v1/{project_id}/bandwidths/{bandwidth_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("bandwidth_id").
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

func GenRespForShowBandwidth() (*model.ShowBandwidthResponse, *def.HttpResponseDef) {
    resp := new(model.ShowBandwidthResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForUpdateBandwidth(request *model.UpdateBandwidthRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPut).
    WithPath("/v1/{project_id}/bandwidths/{bandwidth_id}").
    WithContentType("application/json;charset=UTF-8")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("bandwidth_id").
    WithLocationType(def.Path))



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

func GenRespForUpdateBandwidth() (*model.UpdateBandwidthResponse, *def.HttpResponseDef) {
    resp := new(model.UpdateBandwidthResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForUpdatePrePaidBandwidth(request *model.UpdatePrePaidBandwidthRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPut).
    WithPath("/v2.0/{project_id}/bandwidths/{bandwidth_id}").
    WithContentType("application/json;charset=UTF-8")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("bandwidth_id").
    WithLocationType(def.Path))



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

func GenRespForUpdatePrePaidBandwidth() (*model.UpdatePrePaidBandwidthResponse, *def.HttpResponseDef) {
    resp := new(model.UpdatePrePaidBandwidthResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForBatchCreatePublicipTags(request *model.BatchCreatePublicipTagsRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v2.0/{project_id}/publicips/{publicip_id}/tags/action").
    WithContentType("application/json;charset=UTF-8")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("publicip_id").
    WithLocationType(def.Path))



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

func GenRespForBatchCreatePublicipTags() (*model.BatchCreatePublicipTagsResponse, *def.HttpResponseDef) {
    resp := new(model.BatchCreatePublicipTagsResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForBatchDeletePublicipTags(request *model.BatchDeletePublicipTagsRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v2.0/{project_id}/publicips/{publicip_id}/tags/action").
    WithContentType("application/json;charset=UTF-8")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("publicip_id").
    WithLocationType(def.Path))



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

func GenRespForBatchDeletePublicipTags() (*model.BatchDeletePublicipTagsResponse, *def.HttpResponseDef) {
    resp := new(model.BatchDeletePublicipTagsResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForCreatePrePaidPublicip(request *model.CreatePrePaidPublicipRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v2.0/{project_id}/publicips").
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

func GenRespForCreatePrePaidPublicip() (*model.CreatePrePaidPublicipResponse, *def.HttpResponseDef) {
    resp := new(model.CreatePrePaidPublicipResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForCreatePublicip(request *model.CreatePublicipRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v1/{project_id}/publicips").
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

func GenRespForCreatePublicip() (*model.CreatePublicipResponse, *def.HttpResponseDef) {
    resp := new(model.CreatePublicipResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForCreatePublicipTag(request *model.CreatePublicipTagRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v2.0/{project_id}/publicips/{publicip_id}/tags").
    WithContentType("application/json;charset=UTF-8")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("publicip_id").
    WithLocationType(def.Path))



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

func GenRespForCreatePublicipTag() (*model.CreatePublicipTagResponse, *def.HttpResponseDef) {
    resp := new(model.CreatePublicipTagResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForDeletePublicip(request *model.DeletePublicipRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodDelete).
    WithPath("/v1/{project_id}/publicips/{publicip_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("publicip_id").
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

func GenRespForDeletePublicip() (*model.DeletePublicipResponse, *def.HttpResponseDef) {
    resp := new(model.DeletePublicipResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForDeletePublicipTag(request *model.DeletePublicipTagRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodDelete).
    WithPath("/v2.0/{project_id}/publicips/{publicip_id}/tags/{key}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("publicip_id").
    WithLocationType(def.Path))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("key").
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

func GenRespForDeletePublicipTag() (*model.DeletePublicipTagResponse, *def.HttpResponseDef) {
    resp := new(model.DeletePublicipTagResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForListPublicipTags(request *model.ListPublicipTagsRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v2.0/{project_id}/publicips/tags")





    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForListPublicipTags() (*model.ListPublicipTagsResponse, *def.HttpResponseDef) {
    resp := new(model.ListPublicipTagsResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForListPublicips(request *model.ListPublicipsRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v1/{project_id}/publicips")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("marker").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("limit").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("ip_version").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("enterprise_project_id").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("port_id").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("public_ip_address").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("private_ip_address").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("id").
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

func GenRespForListPublicips() (*model.ListPublicipsResponse, *def.HttpResponseDef) {
    resp := new(model.ListPublicipsResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForListPublicipsByTags(request *model.ListPublicipsByTagsRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v2.0/{project_id}/publicips/resource_instances/action").
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

func GenRespForListPublicipsByTags() (*model.ListPublicipsByTagsResponse, *def.HttpResponseDef) {
    resp := new(model.ListPublicipsByTagsResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForShowPublicip(request *model.ShowPublicipRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v1/{project_id}/publicips/{publicip_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("publicip_id").
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

func GenRespForShowPublicip() (*model.ShowPublicipResponse, *def.HttpResponseDef) {
    resp := new(model.ShowPublicipResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForShowPublicipTags(request *model.ShowPublicipTagsRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v2.0/{project_id}/publicips/{publicip_id}/tags")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("publicip_id").
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

func GenRespForShowPublicipTags() (*model.ShowPublicipTagsResponse, *def.HttpResponseDef) {
    resp := new(model.ShowPublicipTagsResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForUpdatePublicip(request *model.UpdatePublicipRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPut).
    WithPath("/v1/{project_id}/publicips/{publicip_id}").
    WithContentType("application/json;charset=UTF-8")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("publicip_id").
    WithLocationType(def.Path))



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

func GenRespForUpdatePublicip() (*model.UpdatePublicipResponse, *def.HttpResponseDef) {
    resp := new(model.UpdatePublicipResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForNeutronCreateFloatingIp(request *model.NeutronCreateFloatingIpRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v2.0/floatingips").
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

func GenRespForNeutronCreateFloatingIp() (*model.NeutronCreateFloatingIpResponse, *def.HttpResponseDef) {
    resp := new(model.NeutronCreateFloatingIpResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForNeutronDeleteFloatingIp(request *model.NeutronDeleteFloatingIpRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodDelete).
    WithPath("/v2.0/floatingips/{floatingip_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("floatingip_id").
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

func GenRespForNeutronDeleteFloatingIp() (*model.NeutronDeleteFloatingIpResponse, *def.HttpResponseDef) {
    resp := new(model.NeutronDeleteFloatingIpResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForNeutronListFloatingIps(request *model.NeutronListFloatingIpsRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v2.0/floatingips")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("limit").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("marker").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("page_reverse").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("id").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("floating_ip_address").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("router_id").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("port_id").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("fixed_ip_address").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("tenant_id").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("floating_network_id").
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

func GenRespForNeutronListFloatingIps() (*model.NeutronListFloatingIpsResponse, *def.HttpResponseDef) {
    resp := new(model.NeutronListFloatingIpsResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForNeutronShowFloatingIp(request *model.NeutronShowFloatingIpRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v2.0/floatingips/{floatingip_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("floatingip_id").
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

func GenRespForNeutronShowFloatingIp() (*model.NeutronShowFloatingIpResponse, *def.HttpResponseDef) {
    resp := new(model.NeutronShowFloatingIpResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForNeutronUpdateFloatingIp(request *model.NeutronUpdateFloatingIpRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPut).
    WithPath("/v2.0/floatingips/{floatingip_id}").
    WithContentType("application/json;charset=UTF-8")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("floatingip_id").
    WithLocationType(def.Path))



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

func GenRespForNeutronUpdateFloatingIp() (*model.NeutronUpdateFloatingIpResponse, *def.HttpResponseDef) {
    resp := new(model.NeutronUpdateFloatingIpResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

