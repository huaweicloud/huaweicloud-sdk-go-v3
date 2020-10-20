package v3

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3/model"
)

type IamClient struct {
	hcClient *http_client.HcHttpClient
}

func NewIamClient(hcClient *http_client.HcHttpClient) *IamClient {
	return &IamClient{hcClient: hcClient}
}

func IamClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials,basic.Credentials")
	return builder
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)为委托授予所有项目服务权限。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) AssociateAgencyWithAllProjectsPermission(request *model.AssociateAgencyWithAllProjectsPermissionRequest) (*model.AssociateAgencyWithAllProjectsPermissionResponse, error) {
	requestDef := GenReqDefForAssociateAgencyWithAllProjectsPermission(request)
	resp, responseDef := GenRespForAssociateAgencyWithAllProjectsPermission()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)为委托授予全局服务权限。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) AssociateAgencyWithDomainPermission(request *model.AssociateAgencyWithDomainPermissionRequest) (*model.AssociateAgencyWithDomainPermissionResponse, error) {
	requestDef := GenReqDefForAssociateAgencyWithDomainPermission(request)
	resp, responseDef := GenRespForAssociateAgencyWithDomainPermission()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)为委托授予项目服务权限。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) AssociateAgencyWithProjectPermission(request *model.AssociateAgencyWithProjectPermissionRequest) (*model.AssociateAgencyWithProjectPermissionResponse, error) {
	requestDef := GenReqDefForAssociateAgencyWithProjectPermission(request)
	resp, responseDef := GenRespForAssociateAgencyWithProjectPermission()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)检查委托是否具有所有项目服务权限。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) CheckAllProjectsPermissionForAgency(request *model.CheckAllProjectsPermissionForAgencyRequest) (*model.CheckAllProjectsPermissionForAgencyResponse, error) {
	requestDef := GenReqDefForCheckAllProjectsPermissionForAgency(request)
	resp, responseDef := GenRespForCheckAllProjectsPermissionForAgency()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)查询委托是否拥有全局服务权限。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) CheckDomainPermissionForAgency(request *model.CheckDomainPermissionForAgencyRequest) (*model.CheckDomainPermissionForAgencyResponse, error) {
	requestDef := GenReqDefForCheckDomainPermissionForAgency(request)
	resp, responseDef := GenRespForCheckDomainPermissionForAgency()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)查询委托是否拥有项目服务权限。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) CheckProjectPermissionForAgency(request *model.CheckProjectPermissionForAgencyRequest) (*model.CheckProjectPermissionForAgencyResponse, error) {
	requestDef := GenReqDefForCheckProjectPermissionForAgency(request)
	resp, responseDef := GenRespForCheckProjectPermissionForAgency()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)创建委托。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) CreateAgency(request *model.CreateAgencyRequest) (*model.CreateAgencyResponse, error) {
	requestDef := GenReqDefForCreateAgency(request)
	resp, responseDef := GenRespForCreateAgency()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)创建委托自定义策略。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) CreateAgencyCustomPolicy(request *model.CreateAgencyCustomPolicyRequest) (*model.CreateAgencyCustomPolicyResponse, error) {
	requestDef := GenReqDefForCreateAgencyCustomPolicy(request)
	resp, responseDef := GenRespForCreateAgencyCustomPolicy()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)创建云服务自定义策略。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) CreateCloudServiceCustomPolicy(request *model.CreateCloudServiceCustomPolicyRequest) (*model.CreateCloudServiceCustomPolicyResponse, error) {
	requestDef := GenReqDefForCreateCloudServiceCustomPolicy(request)
	resp, responseDef := GenRespForCreateCloudServiceCustomPolicy()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)删除委托。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) DeleteAgency(request *model.DeleteAgencyRequest) (*model.DeleteAgencyResponse, error) {
	requestDef := GenReqDefForDeleteAgency(request)
	resp, responseDef := GenRespForDeleteAgency()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)删除自定义策略。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) DeleteCustomPolicy(request *model.DeleteCustomPolicyRequest) (*model.DeleteCustomPolicyResponse, error) {
	requestDef := GenReqDefForDeleteCustomPolicy(request)
	resp, responseDef := GenRespForDeleteCustomPolicy()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)添加IAM用户到用户组。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneAddUserToGroup(request *model.KeystoneAddUserToGroupRequest) (*model.KeystoneAddUserToGroupResponse, error) {
	requestDef := GenReqDefForKeystoneAddUserToGroup(request)
	resp, responseDef := GenRespForKeystoneAddUserToGroup()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)为用户组授予所有项目服务权限。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneAssociateGroupWithAllProjectPermission(request *model.KeystoneAssociateGroupWithAllProjectPermissionRequest) (*model.KeystoneAssociateGroupWithAllProjectPermissionResponse, error) {
	requestDef := GenReqDefForKeystoneAssociateGroupWithAllProjectPermission(request)
	resp, responseDef := GenRespForKeystoneAssociateGroupWithAllProjectPermission()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)为用户组授予全局服务权限。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneAssociateGroupWithDomainPermission(request *model.KeystoneAssociateGroupWithDomainPermissionRequest) (*model.KeystoneAssociateGroupWithDomainPermissionResponse, error) {
	requestDef := GenReqDefForKeystoneAssociateGroupWithDomainPermission(request)
	resp, responseDef := GenRespForKeystoneAssociateGroupWithDomainPermission()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)为用户组授予项目服务权限。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneAssociateGroupWithProjectPermission(request *model.KeystoneAssociateGroupWithProjectPermissionRequest) (*model.KeystoneAssociateGroupWithProjectPermissionResponse, error) {
	requestDef := GenReqDefForKeystoneAssociateGroupWithProjectPermission(request)
	resp, responseDef := GenRespForKeystoneAssociateGroupWithProjectPermission()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)查询用户组是否拥有全局服务权限。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneCheckDomainPermissionForGroup(request *model.KeystoneCheckDomainPermissionForGroupRequest) (*model.KeystoneCheckDomainPermissionForGroupResponse, error) {
	requestDef := GenReqDefForKeystoneCheckDomainPermissionForGroup(request)
	resp, responseDef := GenRespForKeystoneCheckDomainPermissionForGroup()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)查询用户组是否拥有项目服务权限。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneCheckProjectPermissionForGroup(request *model.KeystoneCheckProjectPermissionForGroupRequest) (*model.KeystoneCheckProjectPermissionForGroupResponse, error) {
	requestDef := GenReqDefForKeystoneCheckProjectPermissionForGroup(request)
	resp, responseDef := GenRespForKeystoneCheckProjectPermissionForGroup()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)查询IAM用户是否在用户组中。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneCheckUserInGroup(request *model.KeystoneCheckUserInGroupRequest) (*model.KeystoneCheckUserInGroupResponse, error) {
	requestDef := GenReqDefForKeystoneCheckUserInGroup(request)
	resp, responseDef := GenRespForKeystoneCheckUserInGroup()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)创建用户组。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneCreateGroup(request *model.KeystoneCreateGroupRequest) (*model.KeystoneCreateGroupResponse, error) {
	requestDef := GenReqDefForKeystoneCreateGroup(request)
	resp, responseDef := GenRespForKeystoneCreateGroup()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)创建项目。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneCreateProject(request *model.KeystoneCreateProjectRequest) (*model.KeystoneCreateProjectResponse, error) {
	requestDef := GenReqDefForKeystoneCreateProject(request)
	resp, responseDef := GenRespForKeystoneCreateProject()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)删除用户组。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneDeleteGroup(request *model.KeystoneDeleteGroupRequest) (*model.KeystoneDeleteGroupResponse, error) {
	requestDef := GenReqDefForKeystoneDeleteGroup(request)
	resp, responseDef := GenRespForKeystoneDeleteGroup()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于查询IAM用户可以用访问的账号详情。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneListAuthDomains(request *model.KeystoneListAuthDomainsRequest) (*model.KeystoneListAuthDomainsResponse, error) {
	requestDef := GenReqDefForKeystoneListAuthDomains(request)
	resp, responseDef := GenRespForKeystoneListAuthDomains()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于查询IAM用户可以访问的项目列表。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneListAuthProjects(request *model.KeystoneListAuthProjectsRequest) (*model.KeystoneListAuthProjectsResponse, error) {
	requestDef := GenReqDefForKeystoneListAuthProjects(request)
	resp, responseDef := GenRespForKeystoneListAuthProjects()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)查询全局服务中的用户组权限。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneListDomainPermissionsForGroup(request *model.KeystoneListDomainPermissionsForGroupRequest) (*model.KeystoneListDomainPermissionsForGroupResponse, error) {
	requestDef := GenReqDefForKeystoneListDomainPermissionsForGroup(request)
	resp, responseDef := GenRespForKeystoneListDomainPermissionsForGroup()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于查询终端节点列表。终端节点用来提供服务访问入口。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneListEndpoints(request *model.KeystoneListEndpointsRequest) (*model.KeystoneListEndpointsResponse, error) {
	requestDef := GenReqDefForKeystoneListEndpoints(request)
	resp, responseDef := GenRespForKeystoneListEndpoints()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)查询用户组列表。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneListGroups(request *model.KeystoneListGroupsRequest) (*model.KeystoneListGroupsResponse, error) {
	requestDef := GenReqDefForKeystoneListGroups(request)
	resp, responseDef := GenRespForKeystoneListGroups()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)查询权限列表。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneListPermissions(request *model.KeystoneListPermissionsRequest) (*model.KeystoneListPermissionsResponse, error) {
	requestDef := GenReqDefForKeystoneListPermissions(request)
	resp, responseDef := GenRespForKeystoneListPermissions()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)查询项目服务中的用户组权限。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneListProjectPermissionsForGroup(request *model.KeystoneListProjectPermissionsForGroupRequest) (*model.KeystoneListProjectPermissionsForGroupResponse, error) {
	requestDef := GenReqDefForKeystoneListProjectPermissionsForGroup(request)
	resp, responseDef := GenRespForKeystoneListProjectPermissionsForGroup()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于查询指定条件下的项目列表。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneListProjects(request *model.KeystoneListProjectsRequest) (*model.KeystoneListProjectsResponse, error) {
	requestDef := GenReqDefForKeystoneListProjects(request)
	resp, responseDef := GenRespForKeystoneListProjects()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)查询指定IAM用户的项目列表，或IAM用户查询自己的项目列表。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneListProjectsForUser(request *model.KeystoneListProjectsForUserRequest) (*model.KeystoneListProjectsForUserResponse, error) {
	requestDef := GenReqDefForKeystoneListProjectsForUser(request)
	resp, responseDef := GenRespForKeystoneListProjectsForUser()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于查询区域列表。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneListRegions(request *model.KeystoneListRegionsRequest) (*model.KeystoneListRegionsResponse, error) {
	requestDef := GenReqDefForKeystoneListRegions(request)
	resp, responseDef := GenRespForKeystoneListRegions()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于查询服务列表。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneListServices(request *model.KeystoneListServicesRequest) (*model.KeystoneListServicesResponse, error) {
	requestDef := GenReqDefForKeystoneListServices(request)
	resp, responseDef := GenRespForKeystoneListServices()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)查询用户组中所包含的IAM用户。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneListUsersForGroupByAdmin(request *model.KeystoneListUsersForGroupByAdminRequest) (*model.KeystoneListUsersForGroupByAdminResponse, error) {
	requestDef := GenReqDefForKeystoneListUsersForGroupByAdmin(request)
	resp, responseDef := GenRespForKeystoneListUsersForGroupByAdmin()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口用于查询Keystone API的版本信息。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneListVersions(request *model.KeystoneListVersionsRequest) (*model.KeystoneListVersionsResponse, error) {
	requestDef := GenReqDefForKeystoneListVersions(request)
	resp, responseDef := GenRespForKeystoneListVersions()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)移除用户组的全局服务权限。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneRemoveDomainPermissionFromGroup(request *model.KeystoneRemoveDomainPermissionFromGroupRequest) (*model.KeystoneRemoveDomainPermissionFromGroupResponse, error) {
	requestDef := GenReqDefForKeystoneRemoveDomainPermissionFromGroup(request)
	resp, responseDef := GenRespForKeystoneRemoveDomainPermissionFromGroup()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)移除用户组的项目服务权限。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneRemoveProjectPermissionFromGroup(request *model.KeystoneRemoveProjectPermissionFromGroupRequest) (*model.KeystoneRemoveProjectPermissionFromGroupResponse, error) {
	requestDef := GenReqDefForKeystoneRemoveProjectPermissionFromGroup(request)
	resp, responseDef := GenRespForKeystoneRemoveProjectPermissionFromGroup()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)移除用户组中的IAM用户。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneRemoveUserFromGroup(request *model.KeystoneRemoveUserFromGroupRequest) (*model.KeystoneRemoveUserFromGroupResponse, error) {
	requestDef := GenReqDefForKeystoneRemoveUserFromGroup(request)
	resp, responseDef := GenRespForKeystoneRemoveUserFromGroup()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于查询请求头中X-Auth-Token对应的服务目录。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneShowCatalog(request *model.KeystoneShowCatalogRequest) (*model.KeystoneShowCatalogResponse, error) {
	requestDef := GenReqDefForKeystoneShowCatalog(request)
	resp, responseDef := GenRespForKeystoneShowCatalog()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于查询终端节点详情。终端节点用来提供服务访问入口。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneShowEndpoint(request *model.KeystoneShowEndpointRequest) (*model.KeystoneShowEndpointResponse, error) {
	requestDef := GenReqDefForKeystoneShowEndpoint(request)
	resp, responseDef := GenRespForKeystoneShowEndpoint()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)查询用户组详情。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneShowGroup(request *model.KeystoneShowGroupRequest) (*model.KeystoneShowGroupResponse, error) {
	requestDef := GenReqDefForKeystoneShowGroup(request)
	resp, responseDef := GenRespForKeystoneShowGroup()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)查询权限详情。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneShowPermission(request *model.KeystoneShowPermissionRequest) (*model.KeystoneShowPermissionResponse, error) {
	requestDef := GenReqDefForKeystoneShowPermission(request)
	resp, responseDef := GenRespForKeystoneShowPermission()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于查询项目详情。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneShowProject(request *model.KeystoneShowProjectRequest) (*model.KeystoneShowProjectResponse, error) {
	requestDef := GenReqDefForKeystoneShowProject(request)
	resp, responseDef := GenRespForKeystoneShowProject()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于查询区域详情。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneShowRegion(request *model.KeystoneShowRegionRequest) (*model.KeystoneShowRegionResponse, error) {
	requestDef := GenReqDefForKeystoneShowRegion(request)
	resp, responseDef := GenRespForKeystoneShowRegion()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于查询账号密码强度策略，查询结果包括密码强度策略的正则表达式及其描述。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneShowSecurityCompliance(request *model.KeystoneShowSecurityComplianceRequest) (*model.KeystoneShowSecurityComplianceResponse, error) {
	requestDef := GenReqDefForKeystoneShowSecurityCompliance(request)
	resp, responseDef := GenRespForKeystoneShowSecurityCompliance()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于按条件查询账号密码强度策略，查询结果包括密码强度策略的正则表达式及其描述。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneShowSecurityComplianceByOption(request *model.KeystoneShowSecurityComplianceByOptionRequest) (*model.KeystoneShowSecurityComplianceByOptionResponse, error) {
	requestDef := GenReqDefForKeystoneShowSecurityComplianceByOption(request)
	resp, responseDef := GenRespForKeystoneShowSecurityComplianceByOption()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于查询服务详情。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneShowService(request *model.KeystoneShowServiceRequest) (*model.KeystoneShowServiceResponse, error) {
	requestDef := GenReqDefForKeystoneShowService(request)
	resp, responseDef := GenRespForKeystoneShowService()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口用于查询Keystone API的3.0版本的信息。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneShowVersion(request *model.KeystoneShowVersionRequest) (*model.KeystoneShowVersionResponse, error) {
	requestDef := GenReqDefForKeystoneShowVersion(request)
	resp, responseDef := GenRespForKeystoneShowVersion()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)更新用户组信息。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneUpdateGroup(request *model.KeystoneUpdateGroupRequest) (*model.KeystoneUpdateGroupResponse, error) {
	requestDef := GenReqDefForKeystoneUpdateGroup(request)
	resp, responseDef := GenRespForKeystoneUpdateGroup()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)修改项目信息。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneUpdateProject(request *model.KeystoneUpdateProjectRequest) (*model.KeystoneUpdateProjectResponse, error) {
	requestDef := GenReqDefForKeystoneUpdateProject(request)
	resp, responseDef := GenRespForKeystoneUpdateProject()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)查询指定条件下的委托列表。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) ListAgencies(request *model.ListAgenciesRequest) (*model.ListAgenciesResponse, error) {
	requestDef := GenReqDefForListAgencies(request)
	resp, responseDef := GenRespForListAgencies()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)查询委托所有项目服务权限列表。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) ListAllProjectsPermissionsForAgency(request *model.ListAllProjectsPermissionsForAgencyRequest) (*model.ListAllProjectsPermissionsForAgencyResponse, error) {
	requestDef := GenReqDefForListAllProjectsPermissionsForAgency(request)
	resp, responseDef := GenRespForListAllProjectsPermissionsForAgency()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)查询自定义策略列表。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) ListCustomPolicies(request *model.ListCustomPoliciesRequest) (*model.ListCustomPoliciesResponse, error) {
	requestDef := GenReqDefForListCustomPolicies(request)
	resp, responseDef := GenRespForListCustomPolicies()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)查询全局服务中的委托权限。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) ListDomainPermissionsForAgency(request *model.ListDomainPermissionsForAgencyRequest) (*model.ListDomainPermissionsForAgencyResponse, error) {
	requestDef := GenReqDefForListDomainPermissionsForAgency(request)
	resp, responseDef := GenRespForListDomainPermissionsForAgency()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)查询项目服务中的委托权限。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) ListProjectPermissionsForAgency(request *model.ListProjectPermissionsForAgencyRequest) (*model.ListProjectPermissionsForAgencyResponse, error) {
	requestDef := GenReqDefForListProjectPermissionsForAgency(request)
	resp, responseDef := GenRespForListProjectPermissionsForAgency()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)移除委托的所有项目服务权限。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) RemoveAllProjectsPermissionFromAgency(request *model.RemoveAllProjectsPermissionFromAgencyRequest) (*model.RemoveAllProjectsPermissionFromAgencyResponse, error) {
	requestDef := GenReqDefForRemoveAllProjectsPermissionFromAgency(request)
	resp, responseDef := GenRespForRemoveAllProjectsPermissionFromAgency()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)移除委托的全局服务权限。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) RemoveDomainPermissionFromAgency(request *model.RemoveDomainPermissionFromAgencyRequest) (*model.RemoveDomainPermissionFromAgencyResponse, error) {
	requestDef := GenReqDefForRemoveDomainPermissionFromAgency(request)
	resp, responseDef := GenRespForRemoveDomainPermissionFromAgency()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)移除委托的项目服务权限。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) RemoveProjectPermissionFromAgency(request *model.RemoveProjectPermissionFromAgencyRequest) (*model.RemoveProjectPermissionFromAgencyResponse, error) {
	requestDef := GenReqDefForRemoveProjectPermissionFromAgency(request)
	resp, responseDef := GenRespForRemoveProjectPermissionFromAgency()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)查询委托详情。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) ShowAgency(request *model.ShowAgencyRequest) (*model.ShowAgencyResponse, error) {
	requestDef := GenReqDefForShowAgency(request)
	resp, responseDef := GenRespForShowAgency()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)查询自定义策略详情。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) ShowCustomPolicy(request *model.ShowCustomPolicyRequest) (*model.ShowCustomPolicyResponse, error) {
	requestDef := GenReqDefForShowCustomPolicy(request)
	resp, responseDef := GenRespForShowCustomPolicy()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于查询账号接口访问控制策略。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) ShowDomainApiAclPolicy(request *model.ShowDomainApiAclPolicyRequest) (*model.ShowDomainApiAclPolicyResponse, error) {
	requestDef := GenReqDefForShowDomainApiAclPolicy(request)
	resp, responseDef := GenRespForShowDomainApiAclPolicy()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于查询账号控制台访问控制策略。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) ShowDomainConsoleAclPolicy(request *model.ShowDomainConsoleAclPolicyRequest) (*model.ShowDomainConsoleAclPolicyResponse, error) {
	requestDef := GenReqDefForShowDomainConsoleAclPolicy(request)
	resp, responseDef := GenRespForShowDomainConsoleAclPolicy()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于查询账号登录策略。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) ShowDomainLoginPolicy(request *model.ShowDomainLoginPolicyRequest) (*model.ShowDomainLoginPolicyResponse, error) {
	requestDef := GenReqDefForShowDomainLoginPolicy(request)
	resp, responseDef := GenRespForShowDomainLoginPolicy()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于查询账号密码策略。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) ShowDomainPasswordPolicy(request *model.ShowDomainPasswordPolicyRequest) (*model.ShowDomainPasswordPolicyResponse, error) {
	requestDef := GenReqDefForShowDomainPasswordPolicy(request)
	resp, responseDef := GenRespForShowDomainPasswordPolicy()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于查询账号操作保护策略。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) ShowDomainProtectPolicy(request *model.ShowDomainProtectPolicyRequest) (*model.ShowDomainProtectPolicyResponse, error) {
	requestDef := GenReqDefForShowDomainProtectPolicy(request)
	resp, responseDef := GenRespForShowDomainProtectPolicy()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于查询账号配额。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) ShowDomainQuota(request *model.ShowDomainQuotaRequest) (*model.ShowDomainQuotaResponse, error) {
	requestDef := GenReqDefForShowDomainQuota(request)
	resp, responseDef := GenRespForShowDomainQuota()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)查询项目详情与状态。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) ShowProjectDetailsAndStatus(request *model.ShowProjectDetailsAndStatusRequest) (*model.ShowProjectDetailsAndStatusResponse, error) {
	requestDef := GenReqDefForShowProjectDetailsAndStatus(request)
	resp, responseDef := GenRespForShowProjectDetailsAndStatus()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于查询项目配额。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) ShowProjectQuota(request *model.ShowProjectQuotaRequest) (*model.ShowProjectQuotaResponse, error) {
	requestDef := GenReqDefForShowProjectQuota(request)
	resp, responseDef := GenRespForShowProjectQuota()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)修改委托。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) UpdateAgency(request *model.UpdateAgencyRequest) (*model.UpdateAgencyResponse, error) {
	requestDef := GenReqDefForUpdateAgency(request)
	resp, responseDef := GenRespForUpdateAgency()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)修改委托自定义策略。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) UpdateAgencyCustomPolicy(request *model.UpdateAgencyCustomPolicyRequest) (*model.UpdateAgencyCustomPolicyResponse, error) {
	requestDef := GenReqDefForUpdateAgencyCustomPolicy(request)
	resp, responseDef := GenRespForUpdateAgencyCustomPolicy()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)修改云服务自定义策略。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) UpdateCloudServiceCustomPolicy(request *model.UpdateCloudServiceCustomPolicyRequest) (*model.UpdateCloudServiceCustomPolicyResponse, error) {
	requestDef := GenReqDefForUpdateCloudServiceCustomPolicy(request)
	resp, responseDef := GenRespForUpdateCloudServiceCustomPolicy()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)修改账号接口访问策略。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) UpdateDomainApiAclPolicy(request *model.UpdateDomainApiAclPolicyRequest) (*model.UpdateDomainApiAclPolicyResponse, error) {
	requestDef := GenReqDefForUpdateDomainApiAclPolicy(request)
	resp, responseDef := GenRespForUpdateDomainApiAclPolicy()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)修改账号控制台访问策略。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) UpdateDomainConsoleAclPolicy(request *model.UpdateDomainConsoleAclPolicyRequest) (*model.UpdateDomainConsoleAclPolicyResponse, error) {
	requestDef := GenReqDefForUpdateDomainConsoleAclPolicy(request)
	resp, responseDef := GenRespForUpdateDomainConsoleAclPolicy()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)修改账号登录策略。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) UpdateDomainLoginPolicy(request *model.UpdateDomainLoginPolicyRequest) (*model.UpdateDomainLoginPolicyResponse, error) {
	requestDef := GenReqDefForUpdateDomainLoginPolicy(request)
	resp, responseDef := GenRespForUpdateDomainLoginPolicy()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)修改账号密码策略。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) UpdateDomainPasswordPolicy(request *model.UpdateDomainPasswordPolicyRequest) (*model.UpdateDomainPasswordPolicyResponse, error) {
	requestDef := GenReqDefForUpdateDomainPasswordPolicy(request)
	resp, responseDef := GenRespForUpdateDomainPasswordPolicy()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)修改账号操作保护策略。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) UpdateDomainProtectPolicy(request *model.UpdateDomainProtectPolicyRequest) (*model.UpdateDomainProtectPolicyResponse, error) {
	requestDef := GenReqDefForUpdateDomainProtectPolicy(request)
	resp, responseDef := GenRespForUpdateDomainProtectPolicy()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)设置项目状态。项目状态包括：正常、冻结。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) UpdateProjectStatus(request *model.UpdateProjectStatusRequest) (*model.UpdateProjectStatusResponse, error) {
	requestDef := GenReqDefForUpdateProjectStatus(request)
	resp, responseDef := GenRespForUpdateProjectStatus()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)给IAM用户创建永久访问密钥，或IAM用户给自己创建永久访问密钥。    访问密钥（Access Key ID/Secret Access Key，简称AK/SK），是您通过开发工具（API、CLI、SDK）访问华为云时的身份凭证，不用于登录控制台。系统通过AK识别访问用户的身份，通过SK进行签名验证，通过加密签名验证可以确保请求的机密性、完整性和请求者身份的正确性。在控制台创建访问密钥的方式请参见：[访问密钥](https://support.huaweicloud.com/usermanual-ca/zh-cn_topic_0046606340.html)  。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) CreatePermanentAccessKey(request *model.CreatePermanentAccessKeyRequest) (*model.CreatePermanentAccessKeyResponse, error) {
	requestDef := GenReqDefForCreatePermanentAccessKey(request)
	resp, responseDef := GenRespForCreatePermanentAccessKey()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于通过委托来获取临时访问密钥（临时AK/SK）和securitytoken。    临时AK/SK和securitytoken是系统颁发给IAM用户的临时访问令牌，有效期为15分钟至24小时，过期后需要重新获取。临时AK/SK和securitytoken遵循权限最小化原则。鉴权时，临时AK/SK和securitytoken必须同时使用，请求头中需要添加“x-security-token”字段，使用方法详情请参考：[API签名参考](https://support.huaweicloud.com/devg-apisign/api-sign-provide.html) 。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) CreateTemporaryAccessKeyByAgency(request *model.CreateTemporaryAccessKeyByAgencyRequest) (*model.CreateTemporaryAccessKeyByAgencyResponse, error) {
	requestDef := GenReqDefForCreateTemporaryAccessKeyByAgency(request)
	resp, responseDef := GenRespForCreateTemporaryAccessKeyByAgency()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于通过token来获取临时AK/SK和securitytoken。    临时AK/SK和securitytoken是系统颁发给IAM用户的临时访问令牌，有效期为15分钟至24小时，过期后需要重新获取。临时AK/SK和securitytoken遵循权限最小化原则。鉴权时，临时AK/SK和securitytoken必须同时使用，请求头中需要添加“x-security-token”字段，使用方法详情请参考：[API签名参考](https://support.huaweicloud.com/devg-apisign/api-sign-provide.html)。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) CreateTemporaryAccessKeyByToken(request *model.CreateTemporaryAccessKeyByTokenRequest) (*model.CreateTemporaryAccessKeyByTokenResponse, error) {
	requestDef := GenReqDefForCreateTemporaryAccessKeyByToken(request)
	resp, responseDef := GenRespForCreateTemporaryAccessKeyByToken()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)删除IAM用户的指定永久访问密钥，或IAM用户删除自己的指定永久访问密钥。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) DeletePermanentAccessKey(request *model.DeletePermanentAccessKeyRequest) (*model.DeletePermanentAccessKeyResponse, error) {
	requestDef := GenReqDefForDeletePermanentAccessKey(request)
	resp, responseDef := GenRespForDeletePermanentAccessKey()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)查询IAM用户的所有永久访问密钥，或IAM用户查询自己的所有永久访问密钥。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) ListPermanentAccessKeys(request *model.ListPermanentAccessKeysRequest) (*model.ListPermanentAccessKeysResponse, error) {
	requestDef := GenReqDefForListPermanentAccessKeys(request)
	resp, responseDef := GenRespForListPermanentAccessKeys()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)查询IAM用户的指定永久访问密钥，或IAM用户查询自己的指定永久访问密钥。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) ShowPermanentAccessKey(request *model.ShowPermanentAccessKeyRequest) (*model.ShowPermanentAccessKeyResponse, error) {
	requestDef := GenReqDefForShowPermanentAccessKey(request)
	resp, responseDef := GenRespForShowPermanentAccessKey()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)修改IAM用户的指定永久访问密钥，或IAM用户修改自己的指定永久访问密钥。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) UpdatePermanentAccessKey(request *model.UpdatePermanentAccessKeyRequest) (*model.UpdatePermanentAccessKeyResponse, error) {
	requestDef := GenReqDefForUpdatePermanentAccessKey(request)
	resp, responseDef := GenRespForUpdatePermanentAccessKey()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)创建IAM用户。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) CreateUser(request *model.CreateUserRequest) (*model.CreateUserResponse, error) {
	requestDef := GenReqDefForCreateUser(request)
	resp, responseDef := GenRespForCreateUser()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)创建IAM用户。IAM用户首次登录时需要修改密码。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneCreateUser(request *model.KeystoneCreateUserRequest) (*model.KeystoneCreateUserResponse, error) {
	requestDef := GenReqDefForKeystoneCreateUser(request)
	resp, responseDef := GenRespForKeystoneCreateUser()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)删除指定IAM用户。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneDeleteUser(request *model.KeystoneDeleteUserRequest) (*model.KeystoneDeleteUserResponse, error) {
	requestDef := GenReqDefForKeystoneDeleteUser(request)
	resp, responseDef := GenRespForKeystoneDeleteUser()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)查询IAM用户所属用户组，或IAM用户查询自己所属用户组。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneListGroupsForUser(request *model.KeystoneListGroupsForUserRequest) (*model.KeystoneListGroupsForUserResponse, error) {
	requestDef := GenReqDefForKeystoneListGroupsForUser(request)
	resp, responseDef := GenRespForKeystoneListGroupsForUser()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)查询IAM用户列表。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneListUsers(request *model.KeystoneListUsersRequest) (*model.KeystoneListUsersResponse, error) {
	requestDef := GenReqDefForKeystoneListUsers(request)
	resp, responseDef := GenRespForKeystoneListUsers()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)查询IAM用户详情，或IAM用户查询自己的用户详情。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneShowUser(request *model.KeystoneShowUserRequest) (*model.KeystoneShowUserResponse, error) {
	requestDef := GenReqDefForKeystoneShowUser(request)
	resp, responseDef := GenRespForKeystoneShowUser()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)修改IAM用户信息。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneUpdateUserByAdmin(request *model.KeystoneUpdateUserByAdminRequest) (*model.KeystoneUpdateUserByAdminResponse, error) {
	requestDef := GenReqDefForKeystoneUpdateUserByAdmin(request)
	resp, responseDef := GenRespForKeystoneUpdateUserByAdmin()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于IAM用户修改自己的密码。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) KeystoneUpdateUserPassword(request *model.KeystoneUpdateUserPasswordRequest) (*model.KeystoneUpdateUserPasswordResponse, error) {
	requestDef := GenReqDefForKeystoneUpdateUserPassword(request)
	resp, responseDef := GenRespForKeystoneUpdateUserPassword()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)查询IAM用户的登录保护状态列表。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) ListUserLoginProtects(request *model.ListUserLoginProtectsRequest) (*model.ListUserLoginProtectsResponse, error) {
	requestDef := GenReqDefForListUserLoginProtects(request)
	resp, responseDef := GenRespForListUserLoginProtects()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)查询IAM用户的MFA绑定信息列表。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) ListUserMfaDevices(request *model.ListUserMfaDevicesRequest) (*model.ListUserMfaDevicesResponse, error) {
	requestDef := GenReqDefForListUserMfaDevices(request)
	resp, responseDef := GenRespForListUserMfaDevices()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)查询IAM用户详情，或IAM用户查询自己的详情。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) ShowUser(request *model.ShowUserRequest) (*model.ShowUserResponse, error) {
	requestDef := GenReqDefForShowUser(request)
	resp, responseDef := GenRespForShowUser()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)查询指定IAM用户的登录保护状态信息，或IAM用户查询自己的登录保护状态信息。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) ShowUserLoginProtect(request *model.ShowUserLoginProtectRequest) (*model.ShowUserLoginProtectResponse, error) {
	requestDef := GenReqDefForShowUserLoginProtect(request)
	resp, responseDef := GenRespForShowUserLoginProtect()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)查询指定IAM用户的MFA绑定信息，或IAM用户查询自己的MFA绑定信息。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) ShowUserMfaDevice(request *model.ShowUserMfaDeviceRequest) (*model.ShowUserMfaDeviceResponse, error) {
	requestDef := GenReqDefForShowUserMfaDevice(request)
	resp, responseDef := GenRespForShowUserMfaDevice()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于[管理员](https://support.huaweicloud.com/usermanual-iam/iam_01_0001.html)修改IAM用户信息 。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) UpdateUser(request *model.UpdateUserRequest) (*model.UpdateUserResponse, error) {
	requestDef := GenReqDefForUpdateUser(request)
	resp, responseDef := GenRespForUpdateUser()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

//该接口可以用于IAM用户修改自己的用户信息。    该接口可以使用全局区域的Endpoint和其他区域的Endpoint调用。IAM的Endpoint请参见：[地区和终端节点](https://developer.huaweicloud.com/endpoint?IAM)。
func (c *IamClient) UpdateUserInformation(request *model.UpdateUserInformationRequest) (*model.UpdateUserInformationResponse, error) {
	requestDef := GenReqDefForUpdateUserInformation(request)
	resp, responseDef := GenRespForUpdateUserInformation()

	if _, err := c.hcClient.Sync(request, requestDef, responseDef); err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}
