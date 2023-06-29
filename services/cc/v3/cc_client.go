package v3

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cc/v3/model"
)

type CcClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCcClient(hcClient *http_client.HcHttpClient) *CcClient {
	return &CcClient{HcClient: hcClient}
}

func CcClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

// CreateAuthorisation 创建授权
//
// 网络实例所属租户授予云连接实例所属租户加载其网络实例的权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) CreateAuthorisation(request *model.CreateAuthorisationRequest) (*model.CreateAuthorisationResponse, error) {
	requestDef := GenReqDefForCreateAuthorisation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAuthorisationResponse), nil
	}
}

// CreateAuthorisationInvoker 创建授权
func (c *CcClient) CreateAuthorisationInvoker(request *model.CreateAuthorisationRequest) *CreateAuthorisationInvoker {
	requestDef := GenReqDefForCreateAuthorisation()
	return &CreateAuthorisationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAuthorisation 删除授权
//
// 网络实例所属租户取消授予云连接实例所属租户加载其网络实例的权限。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) DeleteAuthorisation(request *model.DeleteAuthorisationRequest) (*model.DeleteAuthorisationResponse, error) {
	requestDef := GenReqDefForDeleteAuthorisation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAuthorisationResponse), nil
	}
}

// DeleteAuthorisationInvoker 删除授权
func (c *CcClient) DeleteAuthorisationInvoker(request *model.DeleteAuthorisationRequest) *DeleteAuthorisationInvoker {
	requestDef := GenReqDefForDeleteAuthorisation()
	return &DeleteAuthorisationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAuthorisations 查询授权列表
//
// 网络实例所属租户查看其已经授予其他租户的权限。
// 分页查询使用的参数为marker、limit。marker和limit一起使用时才会生效，单独使用无效。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListAuthorisations(request *model.ListAuthorisationsRequest) (*model.ListAuthorisationsResponse, error) {
	requestDef := GenReqDefForListAuthorisations()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAuthorisationsResponse), nil
	}
}

// ListAuthorisationsInvoker 查询授权列表
func (c *CcClient) ListAuthorisationsInvoker(request *model.ListAuthorisationsRequest) *ListAuthorisationsInvoker {
	requestDef := GenReqDefForListAuthorisations()
	return &ListAuthorisationsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListPermissions 查询权限列表
//
// 云连接实例所属租户查询其可加载其他租户的网络实例权限。
// 分页查询使用的参数为marker、limit。marker和limit一起使用时才会生效，单独使用无效。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListPermissions(request *model.ListPermissionsRequest) (*model.ListPermissionsResponse, error) {
	requestDef := GenReqDefForListPermissions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPermissionsResponse), nil
	}
}

// ListPermissionsInvoker 查询权限列表
func (c *CcClient) ListPermissionsInvoker(request *model.ListPermissionsRequest) *ListPermissionsInvoker {
	requestDef := GenReqDefForListPermissions()
	return &ListPermissionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAuthorisation 更新授权
//
// 更新授权实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) UpdateAuthorisation(request *model.UpdateAuthorisationRequest) (*model.UpdateAuthorisationResponse, error) {
	requestDef := GenReqDefForUpdateAuthorisation()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAuthorisationResponse), nil
	}
}

// UpdateAuthorisationInvoker 更新授权
func (c *CcClient) UpdateAuthorisationInvoker(request *model.UpdateAuthorisationRequest) *UpdateAuthorisationInvoker {
	requestDef := GenReqDefForUpdateAuthorisation()
	return &UpdateAuthorisationInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateBandwidthPackage 将带宽包实例绑定到云连接实例
//
// 将带宽包实例绑定到云连接实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) AssociateBandwidthPackage(request *model.AssociateBandwidthPackageRequest) (*model.AssociateBandwidthPackageResponse, error) {
	requestDef := GenReqDefForAssociateBandwidthPackage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateBandwidthPackageResponse), nil
	}
}

// AssociateBandwidthPackageInvoker 将带宽包实例绑定到云连接实例
func (c *CcClient) AssociateBandwidthPackageInvoker(request *model.AssociateBandwidthPackageRequest) *AssociateBandwidthPackageInvoker {
	requestDef := GenReqDefForAssociateBandwidthPackage()
	return &AssociateBandwidthPackageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateBandwidthPackage 创建带宽包实例
//
// 创建带宽包实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) CreateBandwidthPackage(request *model.CreateBandwidthPackageRequest) (*model.CreateBandwidthPackageResponse, error) {
	requestDef := GenReqDefForCreateBandwidthPackage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateBandwidthPackageResponse), nil
	}
}

// CreateBandwidthPackageInvoker 创建带宽包实例
func (c *CcClient) CreateBandwidthPackageInvoker(request *model.CreateBandwidthPackageRequest) *CreateBandwidthPackageInvoker {
	requestDef := GenReqDefForCreateBandwidthPackage()
	return &CreateBandwidthPackageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBandwidthPackage 删除带宽包实例
//
// 删除带宽包实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) DeleteBandwidthPackage(request *model.DeleteBandwidthPackageRequest) (*model.DeleteBandwidthPackageResponse, error) {
	requestDef := GenReqDefForDeleteBandwidthPackage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBandwidthPackageResponse), nil
	}
}

// DeleteBandwidthPackageInvoker 删除带宽包实例
func (c *CcClient) DeleteBandwidthPackageInvoker(request *model.DeleteBandwidthPackageRequest) *DeleteBandwidthPackageInvoker {
	requestDef := GenReqDefForDeleteBandwidthPackage()
	return &DeleteBandwidthPackageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateBandwidthPackage 解除带宽包实例与云连接实例的绑定
//
// 解除带宽包实例与云连接实例的绑定。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) DisassociateBandwidthPackage(request *model.DisassociateBandwidthPackageRequest) (*model.DisassociateBandwidthPackageResponse, error) {
	requestDef := GenReqDefForDisassociateBandwidthPackage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateBandwidthPackageResponse), nil
	}
}

// DisassociateBandwidthPackageInvoker 解除带宽包实例与云连接实例的绑定
func (c *CcClient) DisassociateBandwidthPackageInvoker(request *model.DisassociateBandwidthPackageRequest) *DisassociateBandwidthPackageInvoker {
	requestDef := GenReqDefForDisassociateBandwidthPackage()
	return &DisassociateBandwidthPackageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBandwidthPackages 查询带宽包列表
//
// 查询带宽包列表。
// 分页查询使用的参数为marker、limit。marker和limit一起使用时才会生效，单独使用无效。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListBandwidthPackages(request *model.ListBandwidthPackagesRequest) (*model.ListBandwidthPackagesResponse, error) {
	requestDef := GenReqDefForListBandwidthPackages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBandwidthPackagesResponse), nil
	}
}

// ListBandwidthPackagesInvoker 查询带宽包列表
func (c *CcClient) ListBandwidthPackagesInvoker(request *model.ListBandwidthPackagesRequest) *ListBandwidthPackagesInvoker {
	requestDef := GenReqDefForListBandwidthPackages()
	return &ListBandwidthPackagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowBandwidthPackage 查询带宽包实例
//
// 查询带宽包实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ShowBandwidthPackage(request *model.ShowBandwidthPackageRequest) (*model.ShowBandwidthPackageResponse, error) {
	requestDef := GenReqDefForShowBandwidthPackage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowBandwidthPackageResponse), nil
	}
}

// ShowBandwidthPackageInvoker 查询带宽包实例
func (c *CcClient) ShowBandwidthPackageInvoker(request *model.ShowBandwidthPackageRequest) *ShowBandwidthPackageInvoker {
	requestDef := GenReqDefForShowBandwidthPackage()
	return &ShowBandwidthPackageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateBandwidthPackage 更新带宽包实例
//
// 更新带宽包实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) UpdateBandwidthPackage(request *model.UpdateBandwidthPackageRequest) (*model.UpdateBandwidthPackageResponse, error) {
	requestDef := GenReqDefForUpdateBandwidthPackage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateBandwidthPackageResponse), nil
	}
}

// UpdateBandwidthPackageInvoker 更新带宽包实例
func (c *CcClient) UpdateBandwidthPackageInvoker(request *model.UpdateBandwidthPackageRequest) *UpdateBandwidthPackageInvoker {
	requestDef := GenReqDefForUpdateBandwidthPackage()
	return &UpdateBandwidthPackageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCloudConnection 创建云连接实例
//
// 创建云连接实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) CreateCloudConnection(request *model.CreateCloudConnectionRequest) (*model.CreateCloudConnectionResponse, error) {
	requestDef := GenReqDefForCreateCloudConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCloudConnectionResponse), nil
	}
}

// CreateCloudConnectionInvoker 创建云连接实例
func (c *CcClient) CreateCloudConnectionInvoker(request *model.CreateCloudConnectionRequest) *CreateCloudConnectionInvoker {
	requestDef := GenReqDefForCreateCloudConnection()
	return &CreateCloudConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCloudConnection 删除云连接实例
//
// 删除云连接实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) DeleteCloudConnection(request *model.DeleteCloudConnectionRequest) (*model.DeleteCloudConnectionResponse, error) {
	requestDef := GenReqDefForDeleteCloudConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCloudConnectionResponse), nil
	}
}

// DeleteCloudConnectionInvoker 删除云连接实例
func (c *CcClient) DeleteCloudConnectionInvoker(request *model.DeleteCloudConnectionRequest) *DeleteCloudConnectionInvoker {
	requestDef := GenReqDefForDeleteCloudConnection()
	return &DeleteCloudConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCloudConnections 查询云连接列表
//
// 查询云连接列表。
// 分页查询使用的参数为marker、limit。marker和limit一起使用时才会生效，单独使用无效。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListCloudConnections(request *model.ListCloudConnectionsRequest) (*model.ListCloudConnectionsResponse, error) {
	requestDef := GenReqDefForListCloudConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCloudConnectionsResponse), nil
	}
}

// ListCloudConnectionsInvoker 查询云连接列表
func (c *CcClient) ListCloudConnectionsInvoker(request *model.ListCloudConnectionsRequest) *ListCloudConnectionsInvoker {
	requestDef := GenReqDefForListCloudConnections()
	return &ListCloudConnectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCloudConnection 查询云连接实例
//
// 查询云连接实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ShowCloudConnection(request *model.ShowCloudConnectionRequest) (*model.ShowCloudConnectionResponse, error) {
	requestDef := GenReqDefForShowCloudConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCloudConnectionResponse), nil
	}
}

// ShowCloudConnectionInvoker 查询云连接实例
func (c *CcClient) ShowCloudConnectionInvoker(request *model.ShowCloudConnectionRequest) *ShowCloudConnectionInvoker {
	requestDef := GenReqDefForShowCloudConnection()
	return &ShowCloudConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCloudConnection 更新云连接实例
//
// 更新云连接实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) UpdateCloudConnection(request *model.UpdateCloudConnectionRequest) (*model.UpdateCloudConnectionResponse, error) {
	requestDef := GenReqDefForUpdateCloudConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCloudConnectionResponse), nil
	}
}

// UpdateCloudConnectionInvoker 更新云连接实例
func (c *CcClient) UpdateCloudConnectionInvoker(request *model.UpdateCloudConnectionRequest) *UpdateCloudConnectionInvoker {
	requestDef := GenReqDefForUpdateCloudConnection()
	return &UpdateCloudConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCloudConnectionRoutes 查询云连接路由条目列表
//
// 查询云连接路由条目列表。
// 分页查询使用的参数为marker、limit。marker和limit一起使用时才会生效，单独使用无效。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListCloudConnectionRoutes(request *model.ListCloudConnectionRoutesRequest) (*model.ListCloudConnectionRoutesResponse, error) {
	requestDef := GenReqDefForListCloudConnectionRoutes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCloudConnectionRoutesResponse), nil
	}
}

// ListCloudConnectionRoutesInvoker 查询云连接路由条目列表
func (c *CcClient) ListCloudConnectionRoutesInvoker(request *model.ListCloudConnectionRoutesRequest) *ListCloudConnectionRoutesInvoker {
	requestDef := GenReqDefForListCloudConnectionRoutes()
	return &ListCloudConnectionRoutesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCloudConnectionRoutes 查询云连接路由条目详情
//
// 查询云连接路由条目列表。
// 分页查询使用的参数为marker、limit。marker和limit一起使用时才会生效，单独使用无效。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ShowCloudConnectionRoutes(request *model.ShowCloudConnectionRoutesRequest) (*model.ShowCloudConnectionRoutesResponse, error) {
	requestDef := GenReqDefForShowCloudConnectionRoutes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCloudConnectionRoutesResponse), nil
	}
}

// ShowCloudConnectionRoutesInvoker 查询云连接路由条目详情
func (c *CcClient) ShowCloudConnectionRoutesInvoker(request *model.ShowCloudConnectionRoutesRequest) *ShowCloudConnectionRoutesInvoker {
	requestDef := GenReqDefForShowCloudConnectionRoutes()
	return &ShowCloudConnectionRoutesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateInterRegionBandwidth 创建域间带宽实例
//
// 创建域间带宽实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) CreateInterRegionBandwidth(request *model.CreateInterRegionBandwidthRequest) (*model.CreateInterRegionBandwidthResponse, error) {
	requestDef := GenReqDefForCreateInterRegionBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateInterRegionBandwidthResponse), nil
	}
}

// CreateInterRegionBandwidthInvoker 创建域间带宽实例
func (c *CcClient) CreateInterRegionBandwidthInvoker(request *model.CreateInterRegionBandwidthRequest) *CreateInterRegionBandwidthInvoker {
	requestDef := GenReqDefForCreateInterRegionBandwidth()
	return &CreateInterRegionBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteInterRegionBandwidth 删除域间带宽实例
//
// 删除域间带宽实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) DeleteInterRegionBandwidth(request *model.DeleteInterRegionBandwidthRequest) (*model.DeleteInterRegionBandwidthResponse, error) {
	requestDef := GenReqDefForDeleteInterRegionBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteInterRegionBandwidthResponse), nil
	}
}

// DeleteInterRegionBandwidthInvoker 删除域间带宽实例
func (c *CcClient) DeleteInterRegionBandwidthInvoker(request *model.DeleteInterRegionBandwidthRequest) *DeleteInterRegionBandwidthInvoker {
	requestDef := GenReqDefForDeleteInterRegionBandwidth()
	return &DeleteInterRegionBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListInterRegionBandwidths 查询域间带宽列表
//
// 查询域间带宽列表。
// 分页查询使用的参数为marker、limit。marker和limit一起使用时才会生效，单独使用无效。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListInterRegionBandwidths(request *model.ListInterRegionBandwidthsRequest) (*model.ListInterRegionBandwidthsResponse, error) {
	requestDef := GenReqDefForListInterRegionBandwidths()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInterRegionBandwidthsResponse), nil
	}
}

// ListInterRegionBandwidthsInvoker 查询域间带宽列表
func (c *CcClient) ListInterRegionBandwidthsInvoker(request *model.ListInterRegionBandwidthsRequest) *ListInterRegionBandwidthsInvoker {
	requestDef := GenReqDefForListInterRegionBandwidths()
	return &ListInterRegionBandwidthsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowInterRegionBandwidth 查询域间带宽实例
//
// 查询域间带宽实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ShowInterRegionBandwidth(request *model.ShowInterRegionBandwidthRequest) (*model.ShowInterRegionBandwidthResponse, error) {
	requestDef := GenReqDefForShowInterRegionBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowInterRegionBandwidthResponse), nil
	}
}

// ShowInterRegionBandwidthInvoker 查询域间带宽实例
func (c *CcClient) ShowInterRegionBandwidthInvoker(request *model.ShowInterRegionBandwidthRequest) *ShowInterRegionBandwidthInvoker {
	requestDef := GenReqDefForShowInterRegionBandwidth()
	return &ShowInterRegionBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateInterRegionBandwidth 更新域间带宽实例
//
// 更新域间带宽实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) UpdateInterRegionBandwidth(request *model.UpdateInterRegionBandwidthRequest) (*model.UpdateInterRegionBandwidthResponse, error) {
	requestDef := GenReqDefForUpdateInterRegionBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateInterRegionBandwidthResponse), nil
	}
}

// UpdateInterRegionBandwidthInvoker 更新域间带宽实例
func (c *CcClient) UpdateInterRegionBandwidthInvoker(request *model.UpdateInterRegionBandwidthRequest) *UpdateInterRegionBandwidthInvoker {
	requestDef := GenReqDefForUpdateInterRegionBandwidth()
	return &UpdateInterRegionBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateNetworkInstance 创建网络实例
//
// 创建网络实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) CreateNetworkInstance(request *model.CreateNetworkInstanceRequest) (*model.CreateNetworkInstanceResponse, error) {
	requestDef := GenReqDefForCreateNetworkInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNetworkInstanceResponse), nil
	}
}

// CreateNetworkInstanceInvoker 创建网络实例
func (c *CcClient) CreateNetworkInstanceInvoker(request *model.CreateNetworkInstanceRequest) *CreateNetworkInstanceInvoker {
	requestDef := GenReqDefForCreateNetworkInstance()
	return &CreateNetworkInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteNetworkInstance 删除网络实例
//
// 删除网络实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) DeleteNetworkInstance(request *model.DeleteNetworkInstanceRequest) (*model.DeleteNetworkInstanceResponse, error) {
	requestDef := GenReqDefForDeleteNetworkInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNetworkInstanceResponse), nil
	}
}

// DeleteNetworkInstanceInvoker 删除网络实例
func (c *CcClient) DeleteNetworkInstanceInvoker(request *model.DeleteNetworkInstanceRequest) *DeleteNetworkInstanceInvoker {
	requestDef := GenReqDefForDeleteNetworkInstance()
	return &DeleteNetworkInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNetworkInstances 查询网络实例列表
//
// 查询云连接列表。
// 分页查询使用的参数为marker、limit。marker和limit一起使用时才会生效，单独使用无效。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListNetworkInstances(request *model.ListNetworkInstancesRequest) (*model.ListNetworkInstancesResponse, error) {
	requestDef := GenReqDefForListNetworkInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNetworkInstancesResponse), nil
	}
}

// ListNetworkInstancesInvoker 查询网络实例列表
func (c *CcClient) ListNetworkInstancesInvoker(request *model.ListNetworkInstancesRequest) *ListNetworkInstancesInvoker {
	requestDef := GenReqDefForListNetworkInstances()
	return &ListNetworkInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNetworkInstance 查询网络实例
//
// 查询网络实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ShowNetworkInstance(request *model.ShowNetworkInstanceRequest) (*model.ShowNetworkInstanceResponse, error) {
	requestDef := GenReqDefForShowNetworkInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNetworkInstanceResponse), nil
	}
}

// ShowNetworkInstanceInvoker 查询网络实例
func (c *CcClient) ShowNetworkInstanceInvoker(request *model.ShowNetworkInstanceRequest) *ShowNetworkInstanceInvoker {
	requestDef := GenReqDefForShowNetworkInstance()
	return &ShowNetworkInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateNetworkInstance 更新网络实例
//
// 更新网络实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) UpdateNetworkInstance(request *model.UpdateNetworkInstanceRequest) (*model.UpdateNetworkInstanceResponse, error) {
	requestDef := GenReqDefForUpdateNetworkInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNetworkInstanceResponse), nil
	}
}

// UpdateNetworkInstanceInvoker 更新网络实例
func (c *CcClient) UpdateNetworkInstanceInvoker(request *model.UpdateNetworkInstanceRequest) *UpdateNetworkInstanceInvoker {
	requestDef := GenReqDefForUpdateNetworkInstance()
	return &UpdateNetworkInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListQuotas 查询配额
//
// 查询配额
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListQuotas(request *model.ListQuotasRequest) (*model.ListQuotasResponse, error) {
	requestDef := GenReqDefForListQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListQuotasResponse), nil
	}
}

// ListQuotasInvoker 查询配额
func (c *CcClient) ListQuotasInvoker(request *model.ListQuotasRequest) *ListQuotasInvoker {
	requestDef := GenReqDefForListQuotas()
	return &ListQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchCreateDeleteTags 批量创建和删除资源标签
//
// 批量创建和删除标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) BatchCreateDeleteTags(request *model.BatchCreateDeleteTagsRequest) (*model.BatchCreateDeleteTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateDeleteTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateDeleteTagsResponse), nil
	}
}

// BatchCreateDeleteTagsInvoker 批量创建和删除资源标签
func (c *CcClient) BatchCreateDeleteTagsInvoker(request *model.BatchCreateDeleteTagsRequest) *BatchCreateDeleteTagsInvoker {
	requestDef := GenReqDefForBatchCreateDeleteTags()
	return &BatchCreateDeleteTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateTag 添加资源标签
//
// 添加资源标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) CreateTag(request *model.CreateTagRequest) (*model.CreateTagResponse, error) {
	requestDef := GenReqDefForCreateTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTagResponse), nil
	}
}

// CreateTagInvoker 添加资源标签
func (c *CcClient) CreateTagInvoker(request *model.CreateTagRequest) *CreateTagInvoker {
	requestDef := GenReqDefForCreateTag()
	return &CreateTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteTag 删除资源标签
//
// 删除资源标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) DeleteTag(request *model.DeleteTagRequest) (*model.DeleteTagResponse, error) {
	requestDef := GenReqDefForDeleteTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTagResponse), nil
	}
}

// DeleteTagInvoker 删除资源标签
func (c *CcClient) DeleteTagInvoker(request *model.DeleteTagRequest) *DeleteTagInvoker {
	requestDef := GenReqDefForDeleteTag()
	return &DeleteTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDomainTags 查询账户资源标签
//
// 查询账户资源标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListDomainTags(request *model.ListDomainTagsRequest) (*model.ListDomainTagsResponse, error) {
	requestDef := GenReqDefForListDomainTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDomainTagsResponse), nil
	}
}

// ListDomainTagsInvoker 查询账户资源标签
func (c *CcClient) ListDomainTagsInvoker(request *model.ListDomainTagsRequest) *ListDomainTagsInvoker {
	requestDef := GenReqDefForListDomainTags()
	return &ListDomainTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListResourceByFilterTag 查询资源实例
//
// 查询资源实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListResourceByFilterTag(request *model.ListResourceByFilterTagRequest) (*model.ListResourceByFilterTagResponse, error) {
	requestDef := GenReqDefForListResourceByFilterTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListResourceByFilterTagResponse), nil
	}
}

// ListResourceByFilterTagInvoker 查询资源实例
func (c *CcClient) ListResourceByFilterTagInvoker(request *model.ListResourceByFilterTagRequest) *ListResourceByFilterTagInvoker {
	requestDef := GenReqDefForListResourceByFilterTag()
	return &ListResourceByFilterTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListTags 查询资源标签
//
// 查询资源标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListTags(request *model.ListTagsRequest) (*model.ListTagsResponse, error) {
	requestDef := GenReqDefForListTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagsResponse), nil
	}
}

// ListTagsInvoker 查询资源标签
func (c *CcClient) ListTagsInvoker(request *model.ListTagsRequest) *ListTagsInvoker {
	requestDef := GenReqDefForListTags()
	return &ListTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
