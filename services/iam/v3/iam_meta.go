package v3

import (
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
    "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3/model"
    "net/http"
)

func GenReqDefForAssociateAgencyWithDomainPermission(request *model.AssociateAgencyWithDomainPermissionRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPut).
    WithPath("/v3.0/OS-AGENCY/domains/{domain_id}/agencies/{agency_id}/roles/{role_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("agency_id").
    WithLocationType(def.Path))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("role_id").
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

func GenRespForAssociateAgencyWithDomainPermission() (*model.AssociateAgencyWithDomainPermissionResponse, *def.HttpResponseDef) {
    resp := new(model.AssociateAgencyWithDomainPermissionResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForAssociateAgencyWithProjectPermission(request *model.AssociateAgencyWithProjectPermissionRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPut).
    WithPath("/v3.0/OS-AGENCY/projects/{project_id}/agencies/{agency_id}/roles/{role_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("agency_id").
    WithLocationType(def.Path))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("role_id").
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

func GenRespForAssociateAgencyWithProjectPermission() (*model.AssociateAgencyWithProjectPermissionResponse, *def.HttpResponseDef) {
    resp := new(model.AssociateAgencyWithProjectPermissionResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForCheckDomainPermissionForAgency(request *model.CheckDomainPermissionForAgencyRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodHead).
    WithPath("/v3.0/OS-AGENCY/domains/{domain_id}/agencies/{agency_id}/roles/{role_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("agency_id").
    WithLocationType(def.Path))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("role_id").
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

func GenRespForCheckDomainPermissionForAgency() (*model.CheckDomainPermissionForAgencyResponse, *def.HttpResponseDef) {
    resp := new(model.CheckDomainPermissionForAgencyResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForCheckProjectPermissionForAgency(request *model.CheckProjectPermissionForAgencyRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodHead).
    WithPath("/v3.0/OS-AGENCY/projects/{project_id}/agencies/{agency_id}/roles/{role_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("agency_id").
    WithLocationType(def.Path))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("role_id").
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

func GenRespForCheckProjectPermissionForAgency() (*model.CheckProjectPermissionForAgencyResponse, *def.HttpResponseDef) {
    resp := new(model.CheckProjectPermissionForAgencyResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForCreateAgency(request *model.CreateAgencyRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v3.0/OS-AGENCY/agencies").
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

func GenRespForCreateAgency() (*model.CreateAgencyResponse, *def.HttpResponseDef) {
    resp := new(model.CreateAgencyResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForCreateAgencyCustomPolicy(request *model.CreateAgencyCustomPolicyRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v3.0/OS-ROLE/roles").
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

func GenRespForCreateAgencyCustomPolicy() (*model.CreateAgencyCustomPolicyResponse, *def.HttpResponseDef) {
    resp := new(model.CreateAgencyCustomPolicyResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForCreateCloudServiceCustomPolicy(request *model.CreateCloudServiceCustomPolicyRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v3.0/OS-ROLE/roles").
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

func GenRespForCreateCloudServiceCustomPolicy() (*model.CreateCloudServiceCustomPolicyResponse, *def.HttpResponseDef) {
    resp := new(model.CreateCloudServiceCustomPolicyResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForDeleteAgency(request *model.DeleteAgencyRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodDelete).
    WithPath("/v3.0/OS-AGENCY/agencies/{agency_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("agency_id").
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

func GenRespForDeleteAgency() (*model.DeleteAgencyResponse, *def.HttpResponseDef) {
    resp := new(model.DeleteAgencyResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForDeleteCustomPolicy(request *model.DeleteCustomPolicyRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodDelete).
    WithPath("/v3.0/OS-ROLE/roles/{role_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("role_id").
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

func GenRespForDeleteCustomPolicy() (*model.DeleteCustomPolicyResponse, *def.HttpResponseDef) {
    resp := new(model.DeleteCustomPolicyResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneAddUserToGroup(request *model.KeystoneAddUserToGroupRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPut).
    WithPath("/v3/groups/{group_id}/users/{user_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("group_id").
    WithLocationType(def.Path))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("user_id").
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

func GenRespForKeystoneAddUserToGroup() (*model.KeystoneAddUserToGroupResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneAddUserToGroupResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneAssociateGroupWithAllProjectPermission(request *model.KeystoneAssociateGroupWithAllProjectPermissionRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPut).
    WithPath("/v3/OS-INHERIT/domains/{domain_id}/groups/{group_id}/roles/{role_id}/inherited_to_projects")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("group_id").
    WithLocationType(def.Path))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("role_id").
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

func GenRespForKeystoneAssociateGroupWithAllProjectPermission() (*model.KeystoneAssociateGroupWithAllProjectPermissionResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneAssociateGroupWithAllProjectPermissionResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneAssociateGroupWithDomainPermission(request *model.KeystoneAssociateGroupWithDomainPermissionRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPut).
    WithPath("/v3/domains/{domain_id}/groups/{group_id}/roles/{role_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("group_id").
    WithLocationType(def.Path))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("role_id").
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

func GenRespForKeystoneAssociateGroupWithDomainPermission() (*model.KeystoneAssociateGroupWithDomainPermissionResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneAssociateGroupWithDomainPermissionResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneAssociateGroupWithProjectPermission(request *model.KeystoneAssociateGroupWithProjectPermissionRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPut).
    WithPath("/v3/projects/{project_id}/groups/{group_id}/roles/{role_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("group_id").
    WithLocationType(def.Path))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("role_id").
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

func GenRespForKeystoneAssociateGroupWithProjectPermission() (*model.KeystoneAssociateGroupWithProjectPermissionResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneAssociateGroupWithProjectPermissionResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneCheckDomainPermissionForGroup(request *model.KeystoneCheckDomainPermissionForGroupRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodHead).
    WithPath("/v3/domains/{domain_id}/groups/{group_id}/roles/{role_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("group_id").
    WithLocationType(def.Path))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("role_id").
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

func GenRespForKeystoneCheckDomainPermissionForGroup() (*model.KeystoneCheckDomainPermissionForGroupResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneCheckDomainPermissionForGroupResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneCheckProjectPermissionForGroup(request *model.KeystoneCheckProjectPermissionForGroupRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodHead).
    WithPath("/v3/projects/{project_id}/groups/{group_id}/roles/{role_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("group_id").
    WithLocationType(def.Path))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("role_id").
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

func GenRespForKeystoneCheckProjectPermissionForGroup() (*model.KeystoneCheckProjectPermissionForGroupResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneCheckProjectPermissionForGroupResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneCheckUserInGroup(request *model.KeystoneCheckUserInGroupRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodHead).
    WithPath("/v3/groups/{group_id}/users/{user_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("group_id").
    WithLocationType(def.Path))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("user_id").
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

func GenRespForKeystoneCheckUserInGroup() (*model.KeystoneCheckUserInGroupResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneCheckUserInGroupResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneCreateGroup(request *model.KeystoneCreateGroupRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v3/groups").
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

func GenRespForKeystoneCreateGroup() (*model.KeystoneCreateGroupResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneCreateGroupResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneCreateProject(request *model.KeystoneCreateProjectRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v3/projects").
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

func GenRespForKeystoneCreateProject() (*model.KeystoneCreateProjectResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneCreateProjectResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneDeleteGroup(request *model.KeystoneDeleteGroupRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodDelete).
    WithPath("/v3/groups/{group_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("group_id").
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

func GenRespForKeystoneDeleteGroup() (*model.KeystoneDeleteGroupResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneDeleteGroupResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneListAuthDomains(request *model.KeystoneListAuthDomainsRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3/auth/domains")





    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForKeystoneListAuthDomains() (*model.KeystoneListAuthDomainsResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneListAuthDomainsResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneListAuthProjects(request *model.KeystoneListAuthProjectsRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3/auth/projects")





    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForKeystoneListAuthProjects() (*model.KeystoneListAuthProjectsResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneListAuthProjectsResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneListDomainPermissionsForGroup(request *model.KeystoneListDomainPermissionsForGroupRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3/domains/{domain_id}/groups/{group_id}/roles")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("group_id").
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

func GenRespForKeystoneListDomainPermissionsForGroup() (*model.KeystoneListDomainPermissionsForGroupResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneListDomainPermissionsForGroupResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneListEndpoints(request *model.KeystoneListEndpointsRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3/endpoints")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("interface").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("service_id").
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

func GenRespForKeystoneListEndpoints() (*model.KeystoneListEndpointsResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneListEndpointsResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneListGroups(request *model.KeystoneListGroupsRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3/groups")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("name").
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

func GenRespForKeystoneListGroups() (*model.KeystoneListGroupsResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneListGroupsResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneListPermissions(request *model.KeystoneListPermissionsRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3/roles")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("name").
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

func GenRespForKeystoneListPermissions() (*model.KeystoneListPermissionsResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneListPermissionsResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneListProjectPermissionsForGroup(request *model.KeystoneListProjectPermissionsForGroupRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3/projects/{project_id}/groups/{group_id}/roles")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("group_id").
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

func GenRespForKeystoneListProjectPermissionsForGroup() (*model.KeystoneListProjectPermissionsForGroupResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneListProjectPermissionsForGroupResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneListProjects(request *model.KeystoneListProjectsRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3/projects")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("name").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("parent_id").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("enabled").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("is_domain").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("page").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("per_page").
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

func GenRespForKeystoneListProjects() (*model.KeystoneListProjectsResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneListProjectsResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneListProjectsForUser(request *model.KeystoneListProjectsForUserRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3/users/{user_id}/projects")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("user_id").
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

func GenRespForKeystoneListProjectsForUser() (*model.KeystoneListProjectsForUserResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneListProjectsForUserResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneListRegions(request *model.KeystoneListRegionsRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3/regions")





    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForKeystoneListRegions() (*model.KeystoneListRegionsResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneListRegionsResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneListServices(request *model.KeystoneListServicesRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3/services")


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

func GenRespForKeystoneListServices() (*model.KeystoneListServicesResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneListServicesResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneListUsersForGroupByAdmin(request *model.KeystoneListUsersForGroupByAdminRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3/groups/{group_id}/users")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("group_id").
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

func GenRespForKeystoneListUsersForGroupByAdmin() (*model.KeystoneListUsersForGroupByAdminResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneListUsersForGroupByAdminResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneListVersions(request *model.KeystoneListVersionsRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/")





    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForKeystoneListVersions() (*model.KeystoneListVersionsResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneListVersionsResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneRemoveDomainPermissionFromGroup(request *model.KeystoneRemoveDomainPermissionFromGroupRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodDelete).
    WithPath("/v3/domains/{domain_id}/groups/{group_id}/roles/{role_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("group_id").
    WithLocationType(def.Path))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("role_id").
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

func GenRespForKeystoneRemoveDomainPermissionFromGroup() (*model.KeystoneRemoveDomainPermissionFromGroupResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneRemoveDomainPermissionFromGroupResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneRemoveProjectPermissionFromGroup(request *model.KeystoneRemoveProjectPermissionFromGroupRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodDelete).
    WithPath("/v3/projects/{project_id}/groups/{group_id}/roles/{role_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("group_id").
    WithLocationType(def.Path))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("role_id").
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

func GenRespForKeystoneRemoveProjectPermissionFromGroup() (*model.KeystoneRemoveProjectPermissionFromGroupResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneRemoveProjectPermissionFromGroupResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneRemoveUserFromGroup(request *model.KeystoneRemoveUserFromGroupRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodDelete).
    WithPath("/v3/groups/{group_id}/users/{user_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("group_id").
    WithLocationType(def.Path))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("user_id").
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

func GenRespForKeystoneRemoveUserFromGroup() (*model.KeystoneRemoveUserFromGroupResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneRemoveUserFromGroupResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneShowCatalog(request *model.KeystoneShowCatalogRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3/auth/catalog")





    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForKeystoneShowCatalog() (*model.KeystoneShowCatalogResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneShowCatalogResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneShowEndpoint(request *model.KeystoneShowEndpointRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3/endpoints/{endpoint_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("endpoint_id").
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

func GenRespForKeystoneShowEndpoint() (*model.KeystoneShowEndpointResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneShowEndpointResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneShowGroup(request *model.KeystoneShowGroupRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3/groups/{group_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("group_id").
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

func GenRespForKeystoneShowGroup() (*model.KeystoneShowGroupResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneShowGroupResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneShowPermission(request *model.KeystoneShowPermissionRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3/roles/{role_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("role_id").
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

func GenRespForKeystoneShowPermission() (*model.KeystoneShowPermissionResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneShowPermissionResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneShowProject(request *model.KeystoneShowProjectRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3/projects/{project_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
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

func GenRespForKeystoneShowProject() (*model.KeystoneShowProjectResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneShowProjectResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneShowRegion(request *model.KeystoneShowRegionRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3/regions/{region_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("region_id").
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

func GenRespForKeystoneShowRegion() (*model.KeystoneShowRegionResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneShowRegionResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneShowSecurityCompliance(request *model.KeystoneShowSecurityComplianceRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3/domains/{domain_id}/config/security_compliance")





    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForKeystoneShowSecurityCompliance() (*model.KeystoneShowSecurityComplianceResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneShowSecurityComplianceResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneShowSecurityComplianceByOption(request *model.KeystoneShowSecurityComplianceByOptionRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3/domains/{domain_id}/config/security_compliance/{option}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("option").
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

func GenRespForKeystoneShowSecurityComplianceByOption() (*model.KeystoneShowSecurityComplianceByOptionResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneShowSecurityComplianceByOptionResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneShowService(request *model.KeystoneShowServiceRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3/services/{service_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("service_id").
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

func GenRespForKeystoneShowService() (*model.KeystoneShowServiceResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneShowServiceResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneShowVersion(request *model.KeystoneShowVersionRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3")





    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForKeystoneShowVersion() (*model.KeystoneShowVersionResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneShowVersionResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneUpdateGroup(request *model.KeystoneUpdateGroupRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPatch).
    WithPath("/v3/groups/{group_id}").
    WithContentType("application/json;charset=UTF-8")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("group_id").
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

func GenRespForKeystoneUpdateGroup() (*model.KeystoneUpdateGroupResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneUpdateGroupResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneUpdateProject(request *model.KeystoneUpdateProjectRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPatch).
    WithPath("/v3/projects/{project_id}").
    WithContentType("application/json;charset=UTF-8")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
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

func GenRespForKeystoneUpdateProject() (*model.KeystoneUpdateProjectResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneUpdateProjectResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForListAgencies(request *model.ListAgenciesRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3.0/OS-AGENCY/agencies")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("trust_domain_id").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("name").
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

func GenRespForListAgencies() (*model.ListAgenciesResponse, *def.HttpResponseDef) {
    resp := new(model.ListAgenciesResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForListCustomPolicies(request *model.ListCustomPoliciesRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3.0/OS-ROLE/roles")





    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("domain_id").
    WithLocationType(def.Path))

    requestDef := reqDefBuilder.Build()
    return requestDef
}

func GenRespForListCustomPolicies() (*model.ListCustomPoliciesResponse, *def.HttpResponseDef) {
    resp := new(model.ListCustomPoliciesResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForListDomainPermissionsForAgency(request *model.ListDomainPermissionsForAgencyRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3.0/OS-AGENCY/domains/{domain_id}/agencies/{agency_id}/roles")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("agency_id").
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

func GenRespForListDomainPermissionsForAgency() (*model.ListDomainPermissionsForAgencyResponse, *def.HttpResponseDef) {
    resp := new(model.ListDomainPermissionsForAgencyResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForListProjectPermissionsForAgency(request *model.ListProjectPermissionsForAgencyRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3.0/OS-AGENCY/projects/{project_id}/agencies/{agency_id}/roles")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("agency_id").
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

func GenRespForListProjectPermissionsForAgency() (*model.ListProjectPermissionsForAgencyResponse, *def.HttpResponseDef) {
    resp := new(model.ListProjectPermissionsForAgencyResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForRemoveDomainPermissionFromAgency(request *model.RemoveDomainPermissionFromAgencyRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodDelete).
    WithPath("/v3.0/OS-AGENCY/domains/{domain_id}/agencies/{agency_id}/roles/{role_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("agency_id").
    WithLocationType(def.Path))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("role_id").
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

func GenRespForRemoveDomainPermissionFromAgency() (*model.RemoveDomainPermissionFromAgencyResponse, *def.HttpResponseDef) {
    resp := new(model.RemoveDomainPermissionFromAgencyResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForRemoveProjectPermissionFromAgency(request *model.RemoveProjectPermissionFromAgencyRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodDelete).
    WithPath("/v3.0/OS-AGENCY/projects/{project_id}/agencies/{agency_id}/roles/{role_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
    WithLocationType(def.Path))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("agency_id").
    WithLocationType(def.Path))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("role_id").
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

func GenRespForRemoveProjectPermissionFromAgency() (*model.RemoveProjectPermissionFromAgencyResponse, *def.HttpResponseDef) {
    resp := new(model.RemoveProjectPermissionFromAgencyResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForShowAgency(request *model.ShowAgencyRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3.0/OS-AGENCY/agencies/{agency_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("agency_id").
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

func GenRespForShowAgency() (*model.ShowAgencyResponse, *def.HttpResponseDef) {
    resp := new(model.ShowAgencyResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForShowCustomPolicy(request *model.ShowCustomPolicyRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3.0/OS-ROLE/roles/{role_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("role_id").
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

func GenRespForShowCustomPolicy() (*model.ShowCustomPolicyResponse, *def.HttpResponseDef) {
    resp := new(model.ShowCustomPolicyResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForShowProjectDetailsAndStatus(request *model.ShowProjectDetailsAndStatusRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3-ext/projects/{project_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
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

func GenRespForShowProjectDetailsAndStatus() (*model.ShowProjectDetailsAndStatusResponse, *def.HttpResponseDef) {
    resp := new(model.ShowProjectDetailsAndStatusResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForUpdateAgency(request *model.UpdateAgencyRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPut).
    WithPath("/v3.0/OS-AGENCY/agencies/{agency_id}").
    WithContentType("application/json;charset=UTF-8")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("agency_id").
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

func GenRespForUpdateAgency() (*model.UpdateAgencyResponse, *def.HttpResponseDef) {
    resp := new(model.UpdateAgencyResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForUpdateAgencyCustomPolicy(request *model.UpdateAgencyCustomPolicyRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPatch).
    WithPath("/v3.0/OS-ROLE/roles/{role_id}").
    WithContentType("application/json;charset=UTF-8")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("role_id").
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

func GenRespForUpdateAgencyCustomPolicy() (*model.UpdateAgencyCustomPolicyResponse, *def.HttpResponseDef) {
    resp := new(model.UpdateAgencyCustomPolicyResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForUpdateCloudServiceCustomPolicy(request *model.UpdateCloudServiceCustomPolicyRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPatch).
    WithPath("/v3.0/OS-ROLE/roles/{role_id}").
    WithContentType("application/json;charset=UTF-8")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("role_id").
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

func GenRespForUpdateCloudServiceCustomPolicy() (*model.UpdateCloudServiceCustomPolicyResponse, *def.HttpResponseDef) {
    resp := new(model.UpdateCloudServiceCustomPolicyResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForUpdateProjectStatus(request *model.UpdateProjectStatusRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPut).
    WithPath("/v3-ext/projects/{project_id}").
    WithContentType("application/json;charset=UTF-8")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("project_id").
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

func GenRespForUpdateProjectStatus() (*model.UpdateProjectStatusResponse, *def.HttpResponseDef) {
    resp := new(model.UpdateProjectStatusResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForCreatePermanentAccessKey(request *model.CreatePermanentAccessKeyRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v3.0/OS-CREDENTIAL/credentials").
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

func GenRespForCreatePermanentAccessKey() (*model.CreatePermanentAccessKeyResponse, *def.HttpResponseDef) {
    resp := new(model.CreatePermanentAccessKeyResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForCreateTemporaryAccessKeyByAgency(request *model.CreateTemporaryAccessKeyByAgencyRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v3.0/OS-CREDENTIAL/securitytokens").
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

func GenRespForCreateTemporaryAccessKeyByAgency() (*model.CreateTemporaryAccessKeyByAgencyResponse, *def.HttpResponseDef) {
    resp := new(model.CreateTemporaryAccessKeyByAgencyResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForCreateTemporaryAccessKeyByToken(request *model.CreateTemporaryAccessKeyByTokenRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v3.0/OS-CREDENTIAL/securitytokens").
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

func GenRespForCreateTemporaryAccessKeyByToken() (*model.CreateTemporaryAccessKeyByTokenResponse, *def.HttpResponseDef) {
    resp := new(model.CreateTemporaryAccessKeyByTokenResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForDeletePermanentAccessKey(request *model.DeletePermanentAccessKeyRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodDelete).
    WithPath("/v3.0/OS-CREDENTIAL/credentials/{access_key}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("access_key").
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

func GenRespForDeletePermanentAccessKey() (*model.DeletePermanentAccessKeyResponse, *def.HttpResponseDef) {
    resp := new(model.DeletePermanentAccessKeyResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForListPermanentAccessKeys(request *model.ListPermanentAccessKeysRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3.0/OS-CREDENTIAL/credentials")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("user_id").
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

func GenRespForListPermanentAccessKeys() (*model.ListPermanentAccessKeysResponse, *def.HttpResponseDef) {
    resp := new(model.ListPermanentAccessKeysResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForShowPermanentAccessKey(request *model.ShowPermanentAccessKeyRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3.0/OS-CREDENTIAL/credentials/{access_key}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("access_key").
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

func GenRespForShowPermanentAccessKey() (*model.ShowPermanentAccessKeyResponse, *def.HttpResponseDef) {
    resp := new(model.ShowPermanentAccessKeyResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForUpdatePermanentAccessKey(request *model.UpdatePermanentAccessKeyRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPut).
    WithPath("/v3.0/OS-CREDENTIAL/credentials/{access_key}").
    WithContentType("application/json;charset=UTF-8")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("access_key").
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

func GenRespForUpdatePermanentAccessKey() (*model.UpdatePermanentAccessKeyResponse, *def.HttpResponseDef) {
    resp := new(model.UpdatePermanentAccessKeyResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForCreateUser(request *model.CreateUserRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v3.0/OS-USER/users").
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

func GenRespForCreateUser() (*model.CreateUserResponse, *def.HttpResponseDef) {
    resp := new(model.CreateUserResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneCreateUser(request *model.KeystoneCreateUserRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v3/users").
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

func GenRespForKeystoneCreateUser() (*model.KeystoneCreateUserResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneCreateUserResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneDeleteUser(request *model.KeystoneDeleteUserRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodDelete).
    WithPath("/v3/users/{user_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("user_id").
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

func GenRespForKeystoneDeleteUser() (*model.KeystoneDeleteUserResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneDeleteUserResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneListGroupsForUser(request *model.KeystoneListGroupsForUserRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3/users/{user_id}/groups")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("user_id").
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

func GenRespForKeystoneListGroupsForUser() (*model.KeystoneListGroupsForUserResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneListGroupsForUserResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneListUsers(request *model.KeystoneListUsersRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3/users")


    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("enabled").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("name").
    WithLocationType(def.Query))
    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("password_expires_at").
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

func GenRespForKeystoneListUsers() (*model.KeystoneListUsersResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneListUsersResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneShowUser(request *model.KeystoneShowUserRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3/users/{user_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("user_id").
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

func GenRespForKeystoneShowUser() (*model.KeystoneShowUserResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneShowUserResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneUpdateUserByAdmin(request *model.KeystoneUpdateUserByAdminRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPatch).
    WithPath("/v3/users/{user_id}").
    WithContentType("application/json;charset=UTF-8")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("user_id").
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

func GenRespForKeystoneUpdateUserByAdmin() (*model.KeystoneUpdateUserByAdminResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneUpdateUserByAdminResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForKeystoneUpdateUserPassword(request *model.KeystoneUpdateUserPasswordRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPost).
    WithPath("/v3/users/{user_id}/password").
    WithContentType("application/json;charset=UTF-8")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("user_id").
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

func GenRespForKeystoneUpdateUserPassword() (*model.KeystoneUpdateUserPasswordResponse, *def.HttpResponseDef) {
    resp := new(model.KeystoneUpdateUserPasswordResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForShowUser(request *model.ShowUserRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodGet).
    WithPath("/v3.0/OS-USER/users/{user_id}")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("user_id").
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

func GenRespForShowUser() (*model.ShowUserResponse, *def.HttpResponseDef) {
    resp := new(model.ShowUserResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForUpdateUser(request *model.UpdateUserRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPut).
    WithPath("/v3.0/OS-USER/users/{user_id}").
    WithContentType("application/json;charset=UTF-8")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("user_id").
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

func GenRespForUpdateUser() (*model.UpdateUserResponse, *def.HttpResponseDef) {
    resp := new(model.UpdateUserResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

func GenReqDefForUpdateUserInformation(request *model.UpdateUserInformationRequest) *def.HttpRequestDef {
    reqDefBuilder := def.NewHttpRequestDefBuilder().
    WithMethod(http.MethodPut).
    WithPath("/v3.0/OS-USER/users/{user_id}/info").
    WithContentType("application/json;charset=UTF-8")

    reqDefBuilder.WithRequestField(def.NewFieldDef().
    WithName("user_id").
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

func GenRespForUpdateUserInformation() (*model.UpdateUserInformationResponse, *def.HttpResponseDef) {
    resp := new(model.UpdateUserInformationResponse)
    respDefBuilder := def.NewHttpResponseDefBuilder().WithBodyJson(resp)
    responseDef := respDefBuilder.Build()
    return resp, responseDef
}

