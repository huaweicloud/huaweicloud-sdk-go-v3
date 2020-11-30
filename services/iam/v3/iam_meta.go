package v3

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3/model"
	"net/http"
)

func GenReqDefForAssociateAgencyWithAllProjectsPermission() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v3.0/OS-INHERIT/domains/{domain_id}/agencies/{agency_id}/roles/{role_id}/inherited_to_projects").
		WithResponse(new(model.AssociateAgencyWithAllProjectsPermissionResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("AgencyId").
		WithJsonTag("agency_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainId").
		WithJsonTag("domain_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("RoleId").
		WithJsonTag("role_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForAssociateAgencyWithDomainPermission() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v3.0/OS-AGENCY/domains/{domain_id}/agencies/{agency_id}/roles/{role_id}").
		WithResponse(new(model.AssociateAgencyWithDomainPermissionResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainId").
		WithJsonTag("domain_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("AgencyId").
		WithJsonTag("agency_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("RoleId").
		WithJsonTag("role_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForAssociateAgencyWithProjectPermission() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v3.0/OS-AGENCY/projects/{project_id}/agencies/{agency_id}/roles/{role_id}").
		WithResponse(new(model.AssociateAgencyWithProjectPermissionResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ProjectId").
		WithJsonTag("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("AgencyId").
		WithJsonTag("agency_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("RoleId").
		WithJsonTag("role_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCheckAllProjectsPermissionForAgency() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodHead).
		WithPath("/v3.0/OS-INHERIT/domains/{domain_id}/agencies/{agency_id}/roles/{role_id}/inherited_to_projects").
		WithResponse(new(model.CheckAllProjectsPermissionForAgencyResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("AgencyId").
		WithJsonTag("agency_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainId").
		WithJsonTag("domain_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("RoleId").
		WithJsonTag("role_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCheckDomainPermissionForAgency() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodHead).
		WithPath("/v3.0/OS-AGENCY/domains/{domain_id}/agencies/{agency_id}/roles/{role_id}").
		WithResponse(new(model.CheckDomainPermissionForAgencyResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainId").
		WithJsonTag("domain_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("AgencyId").
		WithJsonTag("agency_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("RoleId").
		WithJsonTag("role_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCheckProjectPermissionForAgency() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodHead).
		WithPath("/v3.0/OS-AGENCY/projects/{project_id}/agencies/{agency_id}/roles/{role_id}").
		WithResponse(new(model.CheckProjectPermissionForAgencyResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ProjectId").
		WithJsonTag("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("AgencyId").
		WithJsonTag("agency_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("RoleId").
		WithJsonTag("role_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreateAgency() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v3.0/OS-AGENCY/agencies").
		WithResponse(new(model.CreateAgencyResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreateAgencyCustomPolicy() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v3.0/OS-ROLE/roles").
		WithResponse(new(model.CreateAgencyCustomPolicyResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreateCloudServiceCustomPolicy() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v3.0/OS-ROLE/roles").
		WithResponse(new(model.CreateCloudServiceCustomPolicyResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreateLoginToken() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v3.0/OS-AUTH/securitytoken/logintokens").
		WithResponse(new(model.CreateLoginTokenResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("XSubjectLoginToken").
		WithJsonTag("X-Subject-LoginToken").
		WithKindName("string").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreateMetadata() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v3-ext/OS-FEDERATION/identity_providers/{idp_id}/protocols/{protocol_id}/metadata").
		WithResponse(new(model.CreateMetadataResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("IdpId").
		WithJsonTag("idp_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ProtocolId").
		WithJsonTag("protocol_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreateUnscopeTokenByIdpInitiated() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v3.0/OS-FEDERATION/tokens").
		WithResponse(new(model.CreateUnscopeTokenByIdpInitiatedResponse)).
		WithContentType("application/x-www-form-urlencoded")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("XIdpId").
		WithJsonTag("X-Idp-Id").
		WithLocationType(def.Header))

	// response

	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("XSubjectToken").
		WithJsonTag("X-Subject-Token").
		WithKindName("string").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDeleteAgency() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v3.0/OS-AGENCY/agencies/{agency_id}").
		WithResponse(new(model.DeleteAgencyResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("AgencyId").
		WithJsonTag("agency_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDeleteCustomPolicy() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v3.0/OS-ROLE/roles/{role_id}").
		WithResponse(new(model.DeleteCustomPolicyResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("RoleId").
		WithJsonTag("role_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDeleteDomainGroupInheritedRole() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v3/OS-INHERIT/domains/{domain_id}/groups/{group_id}/roles/{role_id}/inherited_to_projects").
		WithResponse(new(model.DeleteDomainGroupInheritedRoleResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainId").
		WithJsonTag("domain_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("GroupId").
		WithJsonTag("group_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("RoleId").
		WithJsonTag("role_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneAddUserToGroup() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v3/groups/{group_id}/users/{user_id}").
		WithResponse(new(model.KeystoneAddUserToGroupResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("GroupId").
		WithJsonTag("group_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("UserId").
		WithJsonTag("user_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneAssociateGroupWithDomainPermission() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v3/domains/{domain_id}/groups/{group_id}/roles/{role_id}").
		WithResponse(new(model.KeystoneAssociateGroupWithDomainPermissionResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainId").
		WithJsonTag("domain_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("GroupId").
		WithJsonTag("group_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("RoleId").
		WithJsonTag("role_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneAssociateGroupWithProjectPermission() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v3/projects/{project_id}/groups/{group_id}/roles/{role_id}").
		WithResponse(new(model.KeystoneAssociateGroupWithProjectPermissionResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ProjectId").
		WithJsonTag("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("GroupId").
		WithJsonTag("group_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("RoleId").
		WithJsonTag("role_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneCheckDomainPermissionForGroup() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodHead).
		WithPath("/v3/domains/{domain_id}/groups/{group_id}/roles/{role_id}").
		WithResponse(new(model.KeystoneCheckDomainPermissionForGroupResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainId").
		WithJsonTag("domain_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("GroupId").
		WithJsonTag("group_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("RoleId").
		WithJsonTag("role_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneCheckProjectPermissionForGroup() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodHead).
		WithPath("/v3/projects/{project_id}/groups/{group_id}/roles/{role_id}").
		WithResponse(new(model.KeystoneCheckProjectPermissionForGroupResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ProjectId").
		WithJsonTag("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("GroupId").
		WithJsonTag("group_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("RoleId").
		WithJsonTag("role_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneCheckUserInGroup() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodHead).
		WithPath("/v3/groups/{group_id}/users/{user_id}").
		WithResponse(new(model.KeystoneCheckUserInGroupResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("GroupId").
		WithJsonTag("group_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("UserId").
		WithJsonTag("user_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneCheckroleForGroup() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodHead).
		WithPath("/v3/OS-INHERIT/domains/{domain_id}/groups/{group_id}/roles/{role_id}/inherited_to_projects").
		WithResponse(new(model.KeystoneCheckroleForGroupResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainId").
		WithJsonTag("domain_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("GroupId").
		WithJsonTag("group_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("RoleId").
		WithJsonTag("role_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneCreateGroup() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v3/groups").
		WithResponse(new(model.KeystoneCreateGroupResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneCreateIdentityProvider() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v3/OS-FEDERATION/identity_providers/{id}").
		WithResponse(new(model.KeystoneCreateIdentityProviderResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Id").
		WithJsonTag("id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneCreateMapping() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v3/OS-FEDERATION/mappings/{id}").
		WithResponse(new(model.KeystoneCreateMappingResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Id").
		WithJsonTag("id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneCreateProject() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v3/projects").
		WithResponse(new(model.KeystoneCreateProjectResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneCreateProtocol() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v3/OS-FEDERATION/identity_providers/{idp_id}/protocols/{protocol_id}").
		WithResponse(new(model.KeystoneCreateProtocolResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("IdpId").
		WithJsonTag("idp_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ProtocolId").
		WithJsonTag("protocol_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneCreateScopedToken() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v3/auth/tokens").
		WithResponse(new(model.KeystoneCreateScopedTokenResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("XSubjectToken").
		WithJsonTag("X-Subject-Token").
		WithKindName("string").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneDeleteGroup() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v3/groups/{group_id}").
		WithResponse(new(model.KeystoneDeleteGroupResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("GroupId").
		WithJsonTag("group_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneDeleteIdentityProvider() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v3/OS-FEDERATION/identity_providers/{id}").
		WithResponse(new(model.KeystoneDeleteIdentityProviderResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Id").
		WithJsonTag("id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneDeleteMapping() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v3/OS-FEDERATION/mappings/{id}").
		WithResponse(new(model.KeystoneDeleteMappingResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Id").
		WithJsonTag("id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneDeleteProtocol() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v3/OS-FEDERATION/identity_providers/{idp_id}/protocols/{protocol_id}").
		WithResponse(new(model.KeystoneDeleteProtocolResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("IdpId").
		WithJsonTag("idp_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ProtocolId").
		WithJsonTag("protocol_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneListAllProjectPermissionsForGroup() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/OS-INHERIT/domains/{domain_id}/groups/{group_id}/roles/inherited_to_projects").
		WithResponse(new(model.KeystoneListAllProjectPermissionsForGroupResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainId").
		WithJsonTag("domain_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("GroupId").
		WithJsonTag("group_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneListAuthDomains() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/auth/domains").
		WithResponse(new(model.KeystoneListAuthDomainsResponse))

	// request

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneListAuthProjects() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/auth/projects").
		WithResponse(new(model.KeystoneListAuthProjectsResponse))

	// request

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneListDomainPermissionsForGroup() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/domains/{domain_id}/groups/{group_id}/roles").
		WithResponse(new(model.KeystoneListDomainPermissionsForGroupResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainId").
		WithJsonTag("domain_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("GroupId").
		WithJsonTag("group_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneListEndpoints() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/endpoints").
		WithResponse(new(model.KeystoneListEndpointsResponse))

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Interface").
		WithJsonTag("interface").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ServiceId").
		WithJsonTag("service_id").
		WithLocationType(def.Query))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneListGroups() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/groups").
		WithResponse(new(model.KeystoneListGroupsResponse))

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainId").
		WithJsonTag("domain_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Name").
		WithJsonTag("name").
		WithLocationType(def.Query))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneListIdentityProviders() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/OS-FEDERATION/identity_providers").
		WithResponse(new(model.KeystoneListIdentityProvidersResponse))

	// request

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneListMappings() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/OS-FEDERATION/mappings").
		WithResponse(new(model.KeystoneListMappingsResponse))

	// request

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneListPermissions() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/roles").
		WithResponse(new(model.KeystoneListPermissionsResponse))

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Name").
		WithJsonTag("name").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainId").
		WithJsonTag("domain_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Page").
		WithJsonTag("page").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("PerPage").
		WithJsonTag("per_page").
		WithLocationType(def.Query))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneListProjectPermissionsForGroup() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/projects/{project_id}/groups/{group_id}/roles").
		WithResponse(new(model.KeystoneListProjectPermissionsForGroupResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ProjectId").
		WithJsonTag("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("GroupId").
		WithJsonTag("group_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneListProjects() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/projects").
		WithResponse(new(model.KeystoneListProjectsResponse))

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainId").
		WithJsonTag("domain_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Name").
		WithJsonTag("name").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ParentId").
		WithJsonTag("parent_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Enabled").
		WithJsonTag("enabled").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("IsDomain").
		WithJsonTag("is_domain").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Page").
		WithJsonTag("page").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("PerPage").
		WithJsonTag("per_page").
		WithLocationType(def.Query))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneListProjectsForUser() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/users/{user_id}/projects").
		WithResponse(new(model.KeystoneListProjectsForUserResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("UserId").
		WithJsonTag("user_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneListProtocols() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/OS-FEDERATION/identity_providers/{idp_id}/protocols").
		WithResponse(new(model.KeystoneListProtocolsResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("IdpId").
		WithJsonTag("idp_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneListRegions() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/regions").
		WithResponse(new(model.KeystoneListRegionsResponse))

	// request

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneListServices() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/services").
		WithResponse(new(model.KeystoneListServicesResponse))

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Type").
		WithJsonTag("type").
		WithLocationType(def.Query))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneListUsersForGroupByAdmin() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/groups/{group_id}/users").
		WithResponse(new(model.KeystoneListUsersForGroupByAdminResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("GroupId").
		WithJsonTag("group_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneListVersions() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/").
		WithResponse(new(model.KeystoneListVersionsResponse))

	// request

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneRemoveDomainPermissionFromGroup() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v3/domains/{domain_id}/groups/{group_id}/roles/{role_id}").
		WithResponse(new(model.KeystoneRemoveDomainPermissionFromGroupResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainId").
		WithJsonTag("domain_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("GroupId").
		WithJsonTag("group_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("RoleId").
		WithJsonTag("role_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneRemoveProjectPermissionFromGroup() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v3/projects/{project_id}/groups/{group_id}/roles/{role_id}").
		WithResponse(new(model.KeystoneRemoveProjectPermissionFromGroupResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ProjectId").
		WithJsonTag("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("GroupId").
		WithJsonTag("group_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("RoleId").
		WithJsonTag("role_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneRemoveUserFromGroup() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v3/groups/{group_id}/users/{user_id}").
		WithResponse(new(model.KeystoneRemoveUserFromGroupResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("GroupId").
		WithJsonTag("group_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("UserId").
		WithJsonTag("user_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneShowCatalog() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/auth/catalog").
		WithResponse(new(model.KeystoneShowCatalogResponse))

	// request

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneShowEndpoint() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/endpoints/{endpoint_id}").
		WithResponse(new(model.KeystoneShowEndpointResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("EndpointId").
		WithJsonTag("endpoint_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneShowGroup() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/groups/{group_id}").
		WithResponse(new(model.KeystoneShowGroupResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("GroupId").
		WithJsonTag("group_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneShowIdentityProvider() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/OS-FEDERATION/identity_providers/{id}").
		WithResponse(new(model.KeystoneShowIdentityProviderResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Id").
		WithJsonTag("id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneShowMapping() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/OS-FEDERATION/mappings/{id}").
		WithResponse(new(model.KeystoneShowMappingResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Id").
		WithJsonTag("id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneShowPermission() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/roles/{role_id}").
		WithResponse(new(model.KeystoneShowPermissionResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("RoleId").
		WithJsonTag("role_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneShowProject() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/projects/{project_id}").
		WithResponse(new(model.KeystoneShowProjectResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ProjectId").
		WithJsonTag("project_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneShowProtocol() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/OS-FEDERATION/identity_providers/{idp_id}/protocols/{protocol_id}").
		WithResponse(new(model.KeystoneShowProtocolResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("IdpId").
		WithJsonTag("idp_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ProtocolId").
		WithJsonTag("protocol_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneShowRegion() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/regions/{region_id}").
		WithResponse(new(model.KeystoneShowRegionResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("RegionId").
		WithJsonTag("region_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneShowSecurityCompliance() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/domains/{domain_id}/config/security_compliance").
		WithResponse(new(model.KeystoneShowSecurityComplianceResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainId").
		WithJsonTag("domain_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneShowSecurityComplianceByOption() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/domains/{domain_id}/config/security_compliance/{option}").
		WithResponse(new(model.KeystoneShowSecurityComplianceByOptionResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainId").
		WithJsonTag("domain_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Option").
		WithJsonTag("option").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneShowService() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/services/{service_id}").
		WithResponse(new(model.KeystoneShowServiceResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ServiceId").
		WithJsonTag("service_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneShowVersion() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3").
		WithResponse(new(model.KeystoneShowVersionResponse))

	// request

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneUpdateGroup() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPatch).
		WithPath("/v3/groups/{group_id}").
		WithResponse(new(model.KeystoneUpdateGroupResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("GroupId").
		WithJsonTag("group_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneUpdateIdentityProvider() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPatch).
		WithPath("/v3/OS-FEDERATION/identity_providers/{id}").
		WithResponse(new(model.KeystoneUpdateIdentityProviderResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Id").
		WithJsonTag("id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneUpdateMapping() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPatch).
		WithPath("/v3/OS-FEDERATION/mappings/{id}").
		WithResponse(new(model.KeystoneUpdateMappingResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Id").
		WithJsonTag("id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneUpdateProject() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPatch).
		WithPath("/v3/projects/{project_id}").
		WithResponse(new(model.KeystoneUpdateProjectResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ProjectId").
		WithJsonTag("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneUpdateProtocol() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPatch).
		WithPath("/v3/OS-FEDERATION/identity_providers/{idp_id}/protocols/{protocol_id}").
		WithResponse(new(model.KeystoneUpdateProtocolResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("IdpId").
		WithJsonTag("idp_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ProtocolId").
		WithJsonTag("protocol_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListAgencies() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3.0/OS-AGENCY/agencies").
		WithResponse(new(model.ListAgenciesResponse))

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainId").
		WithJsonTag("domain_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("TrustDomainId").
		WithJsonTag("trust_domain_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Name").
		WithJsonTag("name").
		WithLocationType(def.Query))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListAllProjectsPermissionsForAgency() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3.0/OS-INHERIT/domains/{domain_id}/agencies/{agency_id}/roles/inherited_to_projects").
		WithResponse(new(model.ListAllProjectsPermissionsForAgencyResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("AgencyId").
		WithJsonTag("agency_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainId").
		WithJsonTag("domain_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListCustomPolicies() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3.0/OS-ROLE/roles").
		WithResponse(new(model.ListCustomPoliciesResponse))

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Page").
		WithJsonTag("page").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("PerPage").
		WithJsonTag("per_page").
		WithLocationType(def.Query))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListDomainPermissionsForAgency() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3.0/OS-AGENCY/domains/{domain_id}/agencies/{agency_id}/roles").
		WithResponse(new(model.ListDomainPermissionsForAgencyResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainId").
		WithJsonTag("domain_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("AgencyId").
		WithJsonTag("agency_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListProjectPermissionsForAgency() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3.0/OS-AGENCY/projects/{project_id}/agencies/{agency_id}/roles").
		WithResponse(new(model.ListProjectPermissionsForAgencyResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ProjectId").
		WithJsonTag("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("AgencyId").
		WithJsonTag("agency_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForRemoveAllProjectsPermissionFromAgency() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v3.0/OS-INHERIT/domains/{domain_id}/agencies/{agency_id}/roles/{role_id}/inherited_to_projects").
		WithResponse(new(model.RemoveAllProjectsPermissionFromAgencyResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("AgencyId").
		WithJsonTag("agency_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainId").
		WithJsonTag("domain_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("RoleId").
		WithJsonTag("role_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForRemoveDomainPermissionFromAgency() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v3.0/OS-AGENCY/domains/{domain_id}/agencies/{agency_id}/roles/{role_id}").
		WithResponse(new(model.RemoveDomainPermissionFromAgencyResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainId").
		WithJsonTag("domain_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("AgencyId").
		WithJsonTag("agency_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("RoleId").
		WithJsonTag("role_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForRemoveProjectPermissionFromAgency() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v3.0/OS-AGENCY/projects/{project_id}/agencies/{agency_id}/roles/{role_id}").
		WithResponse(new(model.RemoveProjectPermissionFromAgencyResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ProjectId").
		WithJsonTag("project_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("AgencyId").
		WithJsonTag("agency_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("RoleId").
		WithJsonTag("role_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowAgency() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3.0/OS-AGENCY/agencies/{agency_id}").
		WithResponse(new(model.ShowAgencyResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("AgencyId").
		WithJsonTag("agency_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowCustomPolicy() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3.0/OS-ROLE/roles/{role_id}").
		WithResponse(new(model.ShowCustomPolicyResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("RoleId").
		WithJsonTag("role_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowDomainApiAclPolicy() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3.0/OS-SECURITYPOLICY/domains/{domain_id}/api-acl-policy").
		WithResponse(new(model.ShowDomainApiAclPolicyResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainId").
		WithJsonTag("domain_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowDomainConsoleAclPolicy() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3.0/OS-SECURITYPOLICY/domains/{domain_id}/console-acl-policy").
		WithResponse(new(model.ShowDomainConsoleAclPolicyResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainId").
		WithJsonTag("domain_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowDomainLoginPolicy() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3.0/OS-SECURITYPOLICY/domains/{domain_id}/login-policy").
		WithResponse(new(model.ShowDomainLoginPolicyResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainId").
		WithJsonTag("domain_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowDomainPasswordPolicy() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3.0/OS-SECURITYPOLICY/domains/{domain_id}/password-policy").
		WithResponse(new(model.ShowDomainPasswordPolicyResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainId").
		WithJsonTag("domain_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowDomainProtectPolicy() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3.0/OS-SECURITYPOLICY/domains/{domain_id}/protect-policy").
		WithResponse(new(model.ShowDomainProtectPolicyResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainId").
		WithJsonTag("domain_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowDomainQuota() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3.0/OS-QUOTA/domains/{domain_id}").
		WithResponse(new(model.ShowDomainQuotaResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainId").
		WithJsonTag("domain_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Type").
		WithJsonTag("type").
		WithLocationType(def.Query))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowMetadata() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3-ext/OS-FEDERATION/identity_providers/{idp_id}/protocols/{protocol_id}/metadata").
		WithResponse(new(model.ShowMetadataResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("IdpId").
		WithJsonTag("idp_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ProtocolId").
		WithJsonTag("protocol_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowProjectDetailsAndStatus() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3-ext/projects/{project_id}").
		WithResponse(new(model.ShowProjectDetailsAndStatusResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ProjectId").
		WithJsonTag("project_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowProjectQuota() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3.0/OS-QUOTA/projects/{project_id}").
		WithResponse(new(model.ShowProjectQuotaResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ProjectId").
		WithJsonTag("project_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForUpdateAgency() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v3.0/OS-AGENCY/agencies/{agency_id}").
		WithResponse(new(model.UpdateAgencyResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("AgencyId").
		WithJsonTag("agency_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForUpdateAgencyCustomPolicy() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPatch).
		WithPath("/v3.0/OS-ROLE/roles/{role_id}").
		WithResponse(new(model.UpdateAgencyCustomPolicyResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("RoleId").
		WithJsonTag("role_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForUpdateCloudServiceCustomPolicy() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPatch).
		WithPath("/v3.0/OS-ROLE/roles/{role_id}").
		WithResponse(new(model.UpdateCloudServiceCustomPolicyResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("RoleId").
		WithJsonTag("role_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForUpdateDomainApiAclPolicy() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v3.0/OS-SECURITYPOLICY/domains/{domain_id}/api-acl-policy").
		WithResponse(new(model.UpdateDomainApiAclPolicyResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainId").
		WithJsonTag("domain_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForUpdateDomainConsoleAclPolicy() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v3.0/OS-SECURITYPOLICY/domains/{domain_id}/console-acl-policy").
		WithResponse(new(model.UpdateDomainConsoleAclPolicyResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainId").
		WithJsonTag("domain_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForUpdateDomainGroupInheritRole() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v3/OS-INHERIT/domains/{domain_id}/groups/{group_id}/roles/{role_id}/inherited_to_projects").
		WithResponse(new(model.UpdateDomainGroupInheritRoleResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainId").
		WithJsonTag("domain_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("GroupId").
		WithJsonTag("group_id").
		WithLocationType(def.Path))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("RoleId").
		WithJsonTag("role_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForUpdateDomainLoginPolicy() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v3.0/OS-SECURITYPOLICY/domains/{domain_id}/login-policy").
		WithResponse(new(model.UpdateDomainLoginPolicyResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainId").
		WithJsonTag("domain_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForUpdateDomainPasswordPolicy() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v3.0/OS-SECURITYPOLICY/domains/{domain_id}/password-policy").
		WithResponse(new(model.UpdateDomainPasswordPolicyResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainId").
		WithJsonTag("domain_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForUpdateDomainProtectPolicy() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v3.0/OS-SECURITYPOLICY/domains/{domain_id}/protect-policy").
		WithResponse(new(model.UpdateDomainProtectPolicyResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainId").
		WithJsonTag("domain_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForUpdateProjectStatus() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v3-ext/projects/{project_id}").
		WithResponse(new(model.UpdateProjectStatusResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("ProjectId").
		WithJsonTag("project_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreatePermanentAccessKey() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v3.0/OS-CREDENTIAL/credentials").
		WithResponse(new(model.CreatePermanentAccessKeyResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreateTemporaryAccessKeyByAgency() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v3.0/OS-CREDENTIAL/securitytokens").
		WithResponse(new(model.CreateTemporaryAccessKeyByAgencyResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreateTemporaryAccessKeyByToken() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v3.0/OS-CREDENTIAL/securitytokens").
		WithResponse(new(model.CreateTemporaryAccessKeyByTokenResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDeletePermanentAccessKey() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v3.0/OS-CREDENTIAL/credentials/{access_key}").
		WithResponse(new(model.DeletePermanentAccessKeyResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("AccessKey").
		WithJsonTag("access_key").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListPermanentAccessKeys() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3.0/OS-CREDENTIAL/credentials").
		WithResponse(new(model.ListPermanentAccessKeysResponse))

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("UserId").
		WithJsonTag("user_id").
		WithLocationType(def.Query))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowPermanentAccessKey() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3.0/OS-CREDENTIAL/credentials/{access_key}").
		WithResponse(new(model.ShowPermanentAccessKeyResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("AccessKey").
		WithJsonTag("access_key").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForUpdatePermanentAccessKey() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v3.0/OS-CREDENTIAL/credentials/{access_key}").
		WithResponse(new(model.UpdatePermanentAccessKeyResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("AccessKey").
		WithJsonTag("access_key").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreateBindingDevice() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v3.0/OS-MFA/mfa-devices/bind").
		WithResponse(new(model.CreateBindingDeviceResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreateMfaDevice() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v3.0/OS-MFA/virtual-mfa-devices").
		WithResponse(new(model.CreateMfaDeviceResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForCreateUser() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v3.0/OS-USER/users").
		WithResponse(new(model.CreateUserResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDeleteBindingDevice() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v3.0/OS-MFA/mfa-devices/unbind").
		WithResponse(new(model.DeleteBindingDeviceResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForDeleteMfaDevice() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v3.0/OS-MFA/virtual-mfa-devices").
		WithResponse(new(model.DeleteMfaDeviceResponse))

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("UserId").
		WithJsonTag("user_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("SerialNumber").
		WithJsonTag("serial_number").
		WithLocationType(def.Query))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneCreateUser() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v3/users").
		WithResponse(new(model.KeystoneCreateUserResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneDeleteUser() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodDelete).
		WithPath("/v3/users/{user_id}").
		WithResponse(new(model.KeystoneDeleteUserResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("UserId").
		WithJsonTag("user_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneListGroupsForUser() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/users/{user_id}/groups").
		WithResponse(new(model.KeystoneListGroupsForUserResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("UserId").
		WithJsonTag("user_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneListUsers() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/users").
		WithResponse(new(model.KeystoneListUsersResponse))

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("DomainId").
		WithJsonTag("domain_id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Enabled").
		WithJsonTag("enabled").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Name").
		WithJsonTag("name").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("PasswordExpiresAt").
		WithJsonTag("password_expires_at").
		WithLocationType(def.Query))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneShowUser() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/users/{user_id}").
		WithResponse(new(model.KeystoneShowUserResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("UserId").
		WithJsonTag("user_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneUpdateUserByAdmin() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPatch).
		WithPath("/v3/users/{user_id}").
		WithResponse(new(model.KeystoneUpdateUserByAdminResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("UserId").
		WithJsonTag("user_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneUpdateUserPassword() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v3/users/{user_id}/password").
		WithResponse(new(model.KeystoneUpdateUserPasswordResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("UserId").
		WithJsonTag("user_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListUserLoginProtects() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3.0/OS-USER/login-protects").
		WithResponse(new(model.ListUserLoginProtectsResponse))

	// request

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForListUserMfaDevices() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3.0/OS-MFA/virtual-mfa-devices").
		WithResponse(new(model.ListUserMfaDevicesResponse))

	// request

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowUser() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3.0/OS-USER/users/{user_id}").
		WithResponse(new(model.ShowUserResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("UserId").
		WithJsonTag("user_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowUserLoginProtect() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3.0/OS-USER/users/{user_id}/login-protect").
		WithResponse(new(model.ShowUserLoginProtectResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("UserId").
		WithJsonTag("user_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForShowUserMfaDevice() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3.0/OS-MFA/users/{user_id}/virtual-mfa-device").
		WithResponse(new(model.ShowUserMfaDeviceResponse))

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("UserId").
		WithJsonTag("user_id").
		WithLocationType(def.Path))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForUpdateLoginProtect() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v3.0/OS-USER/users/{user_id}/login-protect").
		WithResponse(new(model.UpdateLoginProtectResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("UserId").
		WithJsonTag("user_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForUpdateUser() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v3.0/OS-USER/users/{user_id}").
		WithResponse(new(model.UpdateUserResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("UserId").
		WithJsonTag("user_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForUpdateUserInformation() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPut).
		WithPath("/v3.0/OS-USER/users/{user_id}/info").
		WithResponse(new(model.UpdateUserInformationResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("UserId").
		WithJsonTag("user_id").
		WithLocationType(def.Path))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneCreateAgencyToken() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v3/auth/tokens").
		WithResponse(new(model.KeystoneCreateAgencyTokenResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Nocatalog").
		WithJsonTag("nocatalog").
		WithLocationType(def.Query))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("XSubjectToken").
		WithJsonTag("X-Subject-Token").
		WithKindName("string").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneCreateUserTokenByPassword() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v3/auth/tokens").
		WithResponse(new(model.KeystoneCreateUserTokenByPasswordResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Nocatalog").
		WithJsonTag("nocatalog").
		WithLocationType(def.Query))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("XSubjectToken").
		WithJsonTag("X-Subject-Token").
		WithKindName("string").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneCreateUserTokenByPasswordAndMfa() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodPost).
		WithPath("/v3/auth/tokens").
		WithResponse(new(model.KeystoneCreateUserTokenByPasswordAndMfaResponse)).
		WithContentType("application/json;charset=UTF-8")

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Nocatalog").
		WithJsonTag("nocatalog").
		WithLocationType(def.Query))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Body").
		WithLocationType(def.Body))

	// response

	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("XSubjectToken").
		WithJsonTag("X-Subject-Token").
		WithKindName("string").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func GenReqDefForKeystoneValidateToken() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/v3/auth/tokens").
		WithResponse(new(model.KeystoneValidateTokenResponse))

	// request

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Nocatalog").
		WithJsonTag("nocatalog").
		WithLocationType(def.Query))

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("XSubjectToken").
		WithJsonTag("X-Subject-Token").
		WithLocationType(def.Header))

	// response

	reqDefBuilder.WithResponseField(def.NewFieldDef().
		WithName("XSubjectToken").
		WithJsonTag("X-Subject-Token").
		WithKindName("string").
		WithLocationType(def.Header))

	requestDef := reqDefBuilder.Build()
	return requestDef
}
