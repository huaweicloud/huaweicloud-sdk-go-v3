package v3

import (
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v3/model"
    "net/http"
)

func GenReqDefForBatchCreateSubNetworkInterfaceV3(request *model.BatchCreateSubNetworkInterfaceV3Request) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v3/{project_id}/vpc/sub-network-interfaces/batch-create").
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

func GenRespForBatchCreateSubNetworkInterfaceV3() (*model.BatchCreateSubNetworkInterfaceV3Response, *def.HttpResponseDef) {
    resp := new(model.BatchCreateSubNetworkInterfaceV3Response)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForCreateSubNetworkInterface(request *model.CreateSubNetworkInterfaceRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v3/{project_id}/vpc/sub-network-interfaces").
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

func GenRespForCreateSubNetworkInterface() (*model.CreateSubNetworkInterfaceResponse, *def.HttpResponseDef) {
    resp := new(model.CreateSubNetworkInterfaceResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForDeleteSubNetworkInterface(request *model.DeleteSubNetworkInterfaceRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodDelete).
    WithPath("/v3/{project_id}/vpc/sub-network-interfaces/{sub_network_interface_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("sub_network_interface_id").
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

func GenRespForDeleteSubNetworkInterface() (*model.DeleteSubNetworkInterfaceResponse, *def.HttpResponseDef) {
    resp := new(model.DeleteSubNetworkInterfaceResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForListSubNetworkInterfaces(request *model.ListSubNetworkInterfacesRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3/{project_id}/vpc/sub-network-interfaces")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("limit").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("marker").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("id").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("virsubnet_id").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("private_ip_address").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("mac_address").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("vpc_id").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("description").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("parent_id").
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

func GenRespForListSubNetworkInterfaces() (*model.ListSubNetworkInterfacesResponse, *def.HttpResponseDef) {
    resp := new(model.ListSubNetworkInterfacesResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForShowSubNetworkInterface(request *model.ShowSubNetworkInterfaceRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3/{project_id}/vpc/sub-network-interfaces/{sub_network_interface_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("sub_network_interface_id").
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

func GenRespForShowSubNetworkInterface() (*model.ShowSubNetworkInterfaceResponse, *def.HttpResponseDef) {
    resp := new(model.ShowSubNetworkInterfaceResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForShowSubNetworkInterfacesQuantity(request *model.ShowSubNetworkInterfacesQuantityRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3/{project_id}/vpc/sub-network-interfaces/count")





    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForShowSubNetworkInterfacesQuantity() (*model.ShowSubNetworkInterfacesQuantityResponse, *def.HttpResponseDef) {
    resp := new(model.ShowSubNetworkInterfacesQuantityResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForUpdateSubNetworkInterface(request *model.UpdateSubNetworkInterfaceRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPut).
    WithPath("/v3/{project_id}/vpc/sub-network-interfaces/{sub_network_interface_id}").
    WithContentType("application/json")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("sub_network_interface_id").
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

func GenRespForUpdateSubNetworkInterface() (*model.UpdateSubNetworkInterfaceResponse, *def.HttpResponseDef) {
    resp := new(model.UpdateSubNetworkInterfaceResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

