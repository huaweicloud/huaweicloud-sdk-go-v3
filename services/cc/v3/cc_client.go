package v3

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cc/v3/model"
)

type CcClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewCcClient(hcClient *httpclient.HcHttpClient) *CcClient {
	return &CcClient{HcClient: hcClient}
}

func CcClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
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

// ListBandwidthPackageTags 查询带宽包的标签信息
//
// 查询带宽包的标签信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListBandwidthPackageTags(request *model.ListBandwidthPackageTagsRequest) (*model.ListBandwidthPackageTagsResponse, error) {
	requestDef := GenReqDefForListBandwidthPackageTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBandwidthPackageTagsResponse), nil
	}
}

// ListBandwidthPackageTagsInvoker 查询带宽包的标签信息
func (c *CcClient) ListBandwidthPackageTagsInvoker(request *model.ListBandwidthPackageTagsRequest) *ListBandwidthPackageTagsInvoker {
	requestDef := GenReqDefForListBandwidthPackageTags()
	return &ListBandwidthPackageTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListBandwidthPackagesByTags 通过标签过滤带宽包实例
//
// 通过标签过滤带宽包实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListBandwidthPackagesByTags(request *model.ListBandwidthPackagesByTagsRequest) (*model.ListBandwidthPackagesByTagsResponse, error) {
	requestDef := GenReqDefForListBandwidthPackagesByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBandwidthPackagesByTagsResponse), nil
	}
}

// ListBandwidthPackagesByTagsInvoker 通过标签过滤带宽包实例
func (c *CcClient) ListBandwidthPackagesByTagsInvoker(request *model.ListBandwidthPackagesByTagsRequest) *ListBandwidthPackagesByTagsInvoker {
	requestDef := GenReqDefForListBandwidthPackagesByTags()
	return &ListBandwidthPackagesByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// TagBandwidthPackage 创建带宽包标签
//
// 创建带宽包标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) TagBandwidthPackage(request *model.TagBandwidthPackageRequest) (*model.TagBandwidthPackageResponse, error) {
	requestDef := GenReqDefForTagBandwidthPackage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.TagBandwidthPackageResponse), nil
	}
}

// TagBandwidthPackageInvoker 创建带宽包标签
func (c *CcClient) TagBandwidthPackageInvoker(request *model.TagBandwidthPackageRequest) *TagBandwidthPackageInvoker {
	requestDef := GenReqDefForTagBandwidthPackage()
	return &TagBandwidthPackageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UntagBandwidthPackage 删除带宽包标签
//
// 删除带宽包标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) UntagBandwidthPackage(request *model.UntagBandwidthPackageRequest) (*model.UntagBandwidthPackageResponse, error) {
	requestDef := GenReqDefForUntagBandwidthPackage()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UntagBandwidthPackageResponse), nil
	}
}

// UntagBandwidthPackageInvoker 删除带宽包标签
func (c *CcClient) UntagBandwidthPackageInvoker(request *model.UntagBandwidthPackageRequest) *UntagBandwidthPackageInvoker {
	requestDef := GenReqDefForUntagBandwidthPackage()
	return &UntagBandwidthPackageInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ApplyCentralNetworkPolicy 应用中心网络策略
//
// 应用中心网络策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ApplyCentralNetworkPolicy(request *model.ApplyCentralNetworkPolicyRequest) (*model.ApplyCentralNetworkPolicyResponse, error) {
	requestDef := GenReqDefForApplyCentralNetworkPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ApplyCentralNetworkPolicyResponse), nil
	}
}

// ApplyCentralNetworkPolicyInvoker 应用中心网络策略
func (c *CcClient) ApplyCentralNetworkPolicyInvoker(request *model.ApplyCentralNetworkPolicyRequest) *ApplyCentralNetworkPolicyInvoker {
	requestDef := GenReqDefForApplyCentralNetworkPolicy()
	return &ApplyCentralNetworkPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCentralNetwork 创建中心网络
//
// 创建中心网络。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) CreateCentralNetwork(request *model.CreateCentralNetworkRequest) (*model.CreateCentralNetworkResponse, error) {
	requestDef := GenReqDefForCreateCentralNetwork()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCentralNetworkResponse), nil
	}
}

// CreateCentralNetworkInvoker 创建中心网络
func (c *CcClient) CreateCentralNetworkInvoker(request *model.CreateCentralNetworkRequest) *CreateCentralNetworkInvoker {
	requestDef := GenReqDefForCreateCentralNetwork()
	return &CreateCentralNetworkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCentralNetworkPolicy 创建一个新版本的中心网络策略
//
// 创建一份只读的中心网络的策略。如果您有策略文档内容改动，需要基于此版本重新创建一个新版本的策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) CreateCentralNetworkPolicy(request *model.CreateCentralNetworkPolicyRequest) (*model.CreateCentralNetworkPolicyResponse, error) {
	requestDef := GenReqDefForCreateCentralNetworkPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCentralNetworkPolicyResponse), nil
	}
}

// CreateCentralNetworkPolicyInvoker 创建一个新版本的中心网络策略
func (c *CcClient) CreateCentralNetworkPolicyInvoker(request *model.CreateCentralNetworkPolicyRequest) *CreateCentralNetworkPolicyInvoker {
	requestDef := GenReqDefForCreateCentralNetworkPolicy()
	return &CreateCentralNetworkPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCentralNetwork 删除中心网络
//
// 删除中心网络，请先清除附件后再删除中心网络。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) DeleteCentralNetwork(request *model.DeleteCentralNetworkRequest) (*model.DeleteCentralNetworkResponse, error) {
	requestDef := GenReqDefForDeleteCentralNetwork()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCentralNetworkResponse), nil
	}
}

// DeleteCentralNetworkInvoker 删除中心网络
func (c *CcClient) DeleteCentralNetworkInvoker(request *model.DeleteCentralNetworkRequest) *DeleteCentralNetworkInvoker {
	requestDef := GenReqDefForDeleteCentralNetwork()
	return &DeleteCentralNetworkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCentralNetworkPolicy 删除中心网络策略版本
//
// 删除中心网络策略版本。您无法删除正在被应用的中心策略。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) DeleteCentralNetworkPolicy(request *model.DeleteCentralNetworkPolicyRequest) (*model.DeleteCentralNetworkPolicyResponse, error) {
	requestDef := GenReqDefForDeleteCentralNetworkPolicy()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCentralNetworkPolicyResponse), nil
	}
}

// DeleteCentralNetworkPolicyInvoker 删除中心网络策略版本
func (c *CcClient) DeleteCentralNetworkPolicyInvoker(request *model.DeleteCentralNetworkPolicyRequest) *DeleteCentralNetworkPolicyInvoker {
	requestDef := GenReqDefForDeleteCentralNetworkPolicy()
	return &DeleteCentralNetworkPolicyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCentralNetworkPolicies 查询所有版本的中心网络策略列表
//
// 查询所有版本的中心网络策略列表。
// 分页查询使用的参数为marker、limit。limit默认值为0，没有指定marker时返回第一条数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListCentralNetworkPolicies(request *model.ListCentralNetworkPoliciesRequest) (*model.ListCentralNetworkPoliciesResponse, error) {
	requestDef := GenReqDefForListCentralNetworkPolicies()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCentralNetworkPoliciesResponse), nil
	}
}

// ListCentralNetworkPoliciesInvoker 查询所有版本的中心网络策略列表
func (c *CcClient) ListCentralNetworkPoliciesInvoker(request *model.ListCentralNetworkPoliciesRequest) *ListCentralNetworkPoliciesInvoker {
	requestDef := GenReqDefForListCentralNetworkPolicies()
	return &ListCentralNetworkPoliciesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCentralNetworkPolicyChangeSet 查询中心网络策略变化集
//
// 查询与当前应用中心网络策略的变化集。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListCentralNetworkPolicyChangeSet(request *model.ListCentralNetworkPolicyChangeSetRequest) (*model.ListCentralNetworkPolicyChangeSetResponse, error) {
	requestDef := GenReqDefForListCentralNetworkPolicyChangeSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCentralNetworkPolicyChangeSetResponse), nil
	}
}

// ListCentralNetworkPolicyChangeSetInvoker 查询中心网络策略变化集
func (c *CcClient) ListCentralNetworkPolicyChangeSetInvoker(request *model.ListCentralNetworkPolicyChangeSetRequest) *ListCentralNetworkPolicyChangeSetInvoker {
	requestDef := GenReqDefForListCentralNetworkPolicyChangeSet()
	return &ListCentralNetworkPolicyChangeSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCentralNetworkTags 查询中心网络的标签信息
//
// 查询中心网络的标签信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListCentralNetworkTags(request *model.ListCentralNetworkTagsRequest) (*model.ListCentralNetworkTagsResponse, error) {
	requestDef := GenReqDefForListCentralNetworkTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCentralNetworkTagsResponse), nil
	}
}

// ListCentralNetworkTagsInvoker 查询中心网络的标签信息
func (c *CcClient) ListCentralNetworkTagsInvoker(request *model.ListCentralNetworkTagsRequest) *ListCentralNetworkTagsInvoker {
	requestDef := GenReqDefForListCentralNetworkTags()
	return &ListCentralNetworkTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCentralNetworks 查询中心网络列表
//
// 查询中心网络列表。
// 分页查询使用的参数为marker、limit。limit默认值为0，没有指定marker时返回第一条数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListCentralNetworks(request *model.ListCentralNetworksRequest) (*model.ListCentralNetworksResponse, error) {
	requestDef := GenReqDefForListCentralNetworks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCentralNetworksResponse), nil
	}
}

// ListCentralNetworksInvoker 查询中心网络列表
func (c *CcClient) ListCentralNetworksInvoker(request *model.ListCentralNetworksRequest) *ListCentralNetworksInvoker {
	requestDef := GenReqDefForListCentralNetworks()
	return &ListCentralNetworksInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCentralNetworksByTags 通过标签过滤中心网络实例
//
// 通过标签过滤中心网络实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListCentralNetworksByTags(request *model.ListCentralNetworksByTagsRequest) (*model.ListCentralNetworksByTagsResponse, error) {
	requestDef := GenReqDefForListCentralNetworksByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCentralNetworksByTagsResponse), nil
	}
}

// ListCentralNetworksByTagsInvoker 通过标签过滤中心网络实例
func (c *CcClient) ListCentralNetworksByTagsInvoker(request *model.ListCentralNetworksByTagsRequest) *ListCentralNetworksByTagsInvoker {
	requestDef := GenReqDefForListCentralNetworksByTags()
	return &ListCentralNetworksByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCentralNetwork 查询中心网络详情
//
// 查询中心网络详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ShowCentralNetwork(request *model.ShowCentralNetworkRequest) (*model.ShowCentralNetworkResponse, error) {
	requestDef := GenReqDefForShowCentralNetwork()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCentralNetworkResponse), nil
	}
}

// ShowCentralNetworkInvoker 查询中心网络详情
func (c *CcClient) ShowCentralNetworkInvoker(request *model.ShowCentralNetworkRequest) *ShowCentralNetworkInvoker {
	requestDef := GenReqDefForShowCentralNetwork()
	return &ShowCentralNetworkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// TagCentralNetwork 创建中心网络标签
//
// 创建中心网络标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) TagCentralNetwork(request *model.TagCentralNetworkRequest) (*model.TagCentralNetworkResponse, error) {
	requestDef := GenReqDefForTagCentralNetwork()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.TagCentralNetworkResponse), nil
	}
}

// TagCentralNetworkInvoker 创建中心网络标签
func (c *CcClient) TagCentralNetworkInvoker(request *model.TagCentralNetworkRequest) *TagCentralNetworkInvoker {
	requestDef := GenReqDefForTagCentralNetwork()
	return &TagCentralNetworkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UntagCentralNetwork 删除中心网络标签
//
// 删除中心网络标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) UntagCentralNetwork(request *model.UntagCentralNetworkRequest) (*model.UntagCentralNetworkResponse, error) {
	requestDef := GenReqDefForUntagCentralNetwork()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UntagCentralNetworkResponse), nil
	}
}

// UntagCentralNetworkInvoker 删除中心网络标签
func (c *CcClient) UntagCentralNetworkInvoker(request *model.UntagCentralNetworkRequest) *UntagCentralNetworkInvoker {
	requestDef := GenReqDefForUntagCentralNetwork()
	return &UntagCentralNetworkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCentralNetwork 更新中心网络详情
//
// 更新中心网络详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) UpdateCentralNetwork(request *model.UpdateCentralNetworkRequest) (*model.UpdateCentralNetworkResponse, error) {
	requestDef := GenReqDefForUpdateCentralNetwork()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCentralNetworkResponse), nil
	}
}

// UpdateCentralNetworkInvoker 更新中心网络详情
func (c *CcClient) UpdateCentralNetworkInvoker(request *model.UpdateCentralNetworkRequest) *UpdateCentralNetworkInvoker {
	requestDef := GenReqDefForUpdateCentralNetwork()
	return &UpdateCentralNetworkInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCentralNetworkErRouteTableAttachment 创建中心网络ER路由表附件
//
// 创建中心网络的路由表附件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) CreateCentralNetworkErRouteTableAttachment(request *model.CreateCentralNetworkErRouteTableAttachmentRequest) (*model.CreateCentralNetworkErRouteTableAttachmentResponse, error) {
	requestDef := GenReqDefForCreateCentralNetworkErRouteTableAttachment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCentralNetworkErRouteTableAttachmentResponse), nil
	}
}

// CreateCentralNetworkErRouteTableAttachmentInvoker 创建中心网络ER路由表附件
func (c *CcClient) CreateCentralNetworkErRouteTableAttachmentInvoker(request *model.CreateCentralNetworkErRouteTableAttachmentRequest) *CreateCentralNetworkErRouteTableAttachmentInvoker {
	requestDef := GenReqDefForCreateCentralNetworkErRouteTableAttachment()
	return &CreateCentralNetworkErRouteTableAttachmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCentralNetworkGdgwAttachment 创建中心网络GDGW附件
//
// 创建中心网络的GDGW附件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) CreateCentralNetworkGdgwAttachment(request *model.CreateCentralNetworkGdgwAttachmentRequest) (*model.CreateCentralNetworkGdgwAttachmentResponse, error) {
	requestDef := GenReqDefForCreateCentralNetworkGdgwAttachment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCentralNetworkGdgwAttachmentResponse), nil
	}
}

// CreateCentralNetworkGdgwAttachmentInvoker 创建中心网络GDGW附件
func (c *CcClient) CreateCentralNetworkGdgwAttachmentInvoker(request *model.CreateCentralNetworkGdgwAttachmentRequest) *CreateCentralNetworkGdgwAttachmentInvoker {
	requestDef := GenReqDefForCreateCentralNetworkGdgwAttachment()
	return &CreateCentralNetworkGdgwAttachmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCentralNetworkAttachment 删除中心网络附件
//
// 删除中心网络附件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) DeleteCentralNetworkAttachment(request *model.DeleteCentralNetworkAttachmentRequest) (*model.DeleteCentralNetworkAttachmentResponse, error) {
	requestDef := GenReqDefForDeleteCentralNetworkAttachment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCentralNetworkAttachmentResponse), nil
	}
}

// DeleteCentralNetworkAttachmentInvoker 删除中心网络附件
func (c *CcClient) DeleteCentralNetworkAttachmentInvoker(request *model.DeleteCentralNetworkAttachmentRequest) *DeleteCentralNetworkAttachmentInvoker {
	requestDef := GenReqDefForDeleteCentralNetworkAttachment()
	return &DeleteCentralNetworkAttachmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCentralNetworkAttachments 查询中心网络附件列表
//
// 查询中心网络附件列表，分页查询使用的参数为marker、limit。limit默认值为0，没有指定marker时返回第一条数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListCentralNetworkAttachments(request *model.ListCentralNetworkAttachmentsRequest) (*model.ListCentralNetworkAttachmentsResponse, error) {
	requestDef := GenReqDefForListCentralNetworkAttachments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCentralNetworkAttachmentsResponse), nil
	}
}

// ListCentralNetworkAttachmentsInvoker 查询中心网络附件列表
func (c *CcClient) ListCentralNetworkAttachmentsInvoker(request *model.ListCentralNetworkAttachmentsRequest) *ListCentralNetworkAttachmentsInvoker {
	requestDef := GenReqDefForListCentralNetworkAttachments()
	return &ListCentralNetworkAttachmentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCentralNetworkErRouteTableAttachments 查询中心网络ER路由表附件列表
//
// 查询中心网络ER路由表附件列表。
// 分页查询使用的参数为marker、limit。limit默认值为0，没有指定marker时返回第一条数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListCentralNetworkErRouteTableAttachments(request *model.ListCentralNetworkErRouteTableAttachmentsRequest) (*model.ListCentralNetworkErRouteTableAttachmentsResponse, error) {
	requestDef := GenReqDefForListCentralNetworkErRouteTableAttachments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCentralNetworkErRouteTableAttachmentsResponse), nil
	}
}

// ListCentralNetworkErRouteTableAttachmentsInvoker 查询中心网络ER路由表附件列表
func (c *CcClient) ListCentralNetworkErRouteTableAttachmentsInvoker(request *model.ListCentralNetworkErRouteTableAttachmentsRequest) *ListCentralNetworkErRouteTableAttachmentsInvoker {
	requestDef := GenReqDefForListCentralNetworkErRouteTableAttachments()
	return &ListCentralNetworkErRouteTableAttachmentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCentralNetworkGdgwAttachments 查询中心网络GDGW附件列表
//
// 查询中心网络GDGW附件列表。
// 分页查询使用的参数为marker、limit。limit默认值为0，没有指定marker时返回第一条数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListCentralNetworkGdgwAttachments(request *model.ListCentralNetworkGdgwAttachmentsRequest) (*model.ListCentralNetworkGdgwAttachmentsResponse, error) {
	requestDef := GenReqDefForListCentralNetworkGdgwAttachments()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCentralNetworkGdgwAttachmentsResponse), nil
	}
}

// ListCentralNetworkGdgwAttachmentsInvoker 查询中心网络GDGW附件列表
func (c *CcClient) ListCentralNetworkGdgwAttachmentsInvoker(request *model.ListCentralNetworkGdgwAttachmentsRequest) *ListCentralNetworkGdgwAttachmentsInvoker {
	requestDef := GenReqDefForListCentralNetworkGdgwAttachments()
	return &ListCentralNetworkGdgwAttachmentsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCentralNetworkErRouteTableAttachment 查询中心网络ER路由表附件详情
//
// 查询中心网络ER路由表附件详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ShowCentralNetworkErRouteTableAttachment(request *model.ShowCentralNetworkErRouteTableAttachmentRequest) (*model.ShowCentralNetworkErRouteTableAttachmentResponse, error) {
	requestDef := GenReqDefForShowCentralNetworkErRouteTableAttachment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCentralNetworkErRouteTableAttachmentResponse), nil
	}
}

// ShowCentralNetworkErRouteTableAttachmentInvoker 查询中心网络ER路由表附件详情
func (c *CcClient) ShowCentralNetworkErRouteTableAttachmentInvoker(request *model.ShowCentralNetworkErRouteTableAttachmentRequest) *ShowCentralNetworkErRouteTableAttachmentInvoker {
	requestDef := GenReqDefForShowCentralNetworkErRouteTableAttachment()
	return &ShowCentralNetworkErRouteTableAttachmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCentralNetworkGdgwAttachment 查询中心网络GDGW附件详情
//
// 查询中心网络GDGW附件详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ShowCentralNetworkGdgwAttachment(request *model.ShowCentralNetworkGdgwAttachmentRequest) (*model.ShowCentralNetworkGdgwAttachmentResponse, error) {
	requestDef := GenReqDefForShowCentralNetworkGdgwAttachment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCentralNetworkGdgwAttachmentResponse), nil
	}
}

// ShowCentralNetworkGdgwAttachmentInvoker 查询中心网络GDGW附件详情
func (c *CcClient) ShowCentralNetworkGdgwAttachmentInvoker(request *model.ShowCentralNetworkGdgwAttachmentRequest) *ShowCentralNetworkGdgwAttachmentInvoker {
	requestDef := GenReqDefForShowCentralNetworkGdgwAttachment()
	return &ShowCentralNetworkGdgwAttachmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCentralNetworkErRouteTableAttachment 更新中心网络ER路由表附件
//
// 更新中心网络ER路由表附件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) UpdateCentralNetworkErRouteTableAttachment(request *model.UpdateCentralNetworkErRouteTableAttachmentRequest) (*model.UpdateCentralNetworkErRouteTableAttachmentResponse, error) {
	requestDef := GenReqDefForUpdateCentralNetworkErRouteTableAttachment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCentralNetworkErRouteTableAttachmentResponse), nil
	}
}

// UpdateCentralNetworkErRouteTableAttachmentInvoker 更新中心网络ER路由表附件
func (c *CcClient) UpdateCentralNetworkErRouteTableAttachmentInvoker(request *model.UpdateCentralNetworkErRouteTableAttachmentRequest) *UpdateCentralNetworkErRouteTableAttachmentInvoker {
	requestDef := GenReqDefForUpdateCentralNetworkErRouteTableAttachment()
	return &UpdateCentralNetworkErRouteTableAttachmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCentralNetworkGdgwAttachment 更新中心网络GDGW附件
//
// 更新中心网络GDGW附件。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) UpdateCentralNetworkGdgwAttachment(request *model.UpdateCentralNetworkGdgwAttachmentRequest) (*model.UpdateCentralNetworkGdgwAttachmentResponse, error) {
	requestDef := GenReqDefForUpdateCentralNetworkGdgwAttachment()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCentralNetworkGdgwAttachmentResponse), nil
	}
}

// UpdateCentralNetworkGdgwAttachmentInvoker 更新中心网络GDGW附件
func (c *CcClient) UpdateCentralNetworkGdgwAttachmentInvoker(request *model.UpdateCentralNetworkGdgwAttachmentRequest) *UpdateCentralNetworkGdgwAttachmentInvoker {
	requestDef := GenReqDefForUpdateCentralNetworkGdgwAttachment()
	return &UpdateCentralNetworkGdgwAttachmentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCentralNetworkCapabilities 查询中心网络能力列表
//
// 查询中心网络能力列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListCentralNetworkCapabilities(request *model.ListCentralNetworkCapabilitiesRequest) (*model.ListCentralNetworkCapabilitiesResponse, error) {
	requestDef := GenReqDefForListCentralNetworkCapabilities()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCentralNetworkCapabilitiesResponse), nil
	}
}

// ListCentralNetworkCapabilitiesInvoker 查询中心网络能力列表
func (c *CcClient) ListCentralNetworkCapabilitiesInvoker(request *model.ListCentralNetworkCapabilitiesRequest) *ListCentralNetworkCapabilitiesInvoker {
	requestDef := GenReqDefForListCentralNetworkCapabilities()
	return &ListCentralNetworkCapabilitiesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCentralNetworkConnections 查询中心网络连接列表
//
// 查询中心网络连接列表接口。
// 分页查询使用的参数为marker、limit。limit默认值为0，没有指定marker时返回第一条数据。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListCentralNetworkConnections(request *model.ListCentralNetworkConnectionsRequest) (*model.ListCentralNetworkConnectionsResponse, error) {
	requestDef := GenReqDefForListCentralNetworkConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCentralNetworkConnectionsResponse), nil
	}
}

// ListCentralNetworkConnectionsInvoker 查询中心网络连接列表
func (c *CcClient) ListCentralNetworkConnectionsInvoker(request *model.ListCentralNetworkConnectionsRequest) *ListCentralNetworkConnectionsInvoker {
	requestDef := GenReqDefForListCentralNetworkConnections()
	return &ListCentralNetworkConnectionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateCentralNetworkConnection 更新中心网络连接接口
//
// 更新中心网络连接接口（仅支持更新带宽）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) UpdateCentralNetworkConnection(request *model.UpdateCentralNetworkConnectionRequest) (*model.UpdateCentralNetworkConnectionResponse, error) {
	requestDef := GenReqDefForUpdateCentralNetworkConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCentralNetworkConnectionResponse), nil
	}
}

// UpdateCentralNetworkConnectionInvoker 更新中心网络连接接口
func (c *CcClient) UpdateCentralNetworkConnectionInvoker(request *model.UpdateCentralNetworkConnectionRequest) *UpdateCentralNetworkConnectionInvoker {
	requestDef := GenReqDefForUpdateCentralNetworkConnection()
	return &UpdateCentralNetworkConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCentralNetworkQuotas 查询中心网络配额
//
// 查询中心网络配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListCentralNetworkQuotas(request *model.ListCentralNetworkQuotasRequest) (*model.ListCentralNetworkQuotasResponse, error) {
	requestDef := GenReqDefForListCentralNetworkQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCentralNetworkQuotasResponse), nil
	}
}

// ListCentralNetworkQuotasInvoker 查询中心网络配额
func (c *CcClient) ListCentralNetworkQuotasInvoker(request *model.ListCentralNetworkQuotasRequest) *ListCentralNetworkQuotasInvoker {
	requestDef := GenReqDefForListCentralNetworkQuotas()
	return &ListCentralNetworkQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListCloudConnectionTags 查询云连接实例的标签信息
//
// 查询云连接实例的标签信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListCloudConnectionTags(request *model.ListCloudConnectionTagsRequest) (*model.ListCloudConnectionTagsResponse, error) {
	requestDef := GenReqDefForListCloudConnectionTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCloudConnectionTagsResponse), nil
	}
}

// ListCloudConnectionTagsInvoker 查询云连接实例的标签信息
func (c *CcClient) ListCloudConnectionTagsInvoker(request *model.ListCloudConnectionTagsRequest) *ListCloudConnectionTagsInvoker {
	requestDef := GenReqDefForListCloudConnectionTags()
	return &ListCloudConnectionTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListCloudConnectionsByTags 通过标签过滤云连接实例
//
// 通过标签过滤云连接实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListCloudConnectionsByTags(request *model.ListCloudConnectionsByTagsRequest) (*model.ListCloudConnectionsByTagsResponse, error) {
	requestDef := GenReqDefForListCloudConnectionsByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCloudConnectionsByTagsResponse), nil
	}
}

// ListCloudConnectionsByTagsInvoker 通过标签过滤云连接实例
func (c *CcClient) ListCloudConnectionsByTagsInvoker(request *model.ListCloudConnectionsByTagsRequest) *ListCloudConnectionsByTagsInvoker {
	requestDef := GenReqDefForListCloudConnectionsByTags()
	return &ListCloudConnectionsByTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// TagCloudConnection 创建云连接实例标签
//
// 创建云连接实例标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) TagCloudConnection(request *model.TagCloudConnectionRequest) (*model.TagCloudConnectionResponse, error) {
	requestDef := GenReqDefForTagCloudConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.TagCloudConnectionResponse), nil
	}
}

// TagCloudConnectionInvoker 创建云连接实例标签
func (c *CcClient) TagCloudConnectionInvoker(request *model.TagCloudConnectionRequest) *TagCloudConnectionInvoker {
	requestDef := GenReqDefForTagCloudConnection()
	return &TagCloudConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UntagCloudConnection 删除云连接实例标签
//
// 删除云连接实例标签。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) UntagCloudConnection(request *model.UntagCloudConnectionRequest) (*model.UntagCloudConnectionResponse, error) {
	requestDef := GenReqDefForUntagCloudConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UntagCloudConnectionResponse), nil
	}
}

// UntagCloudConnectionInvoker 删除云连接实例标签
func (c *CcClient) UntagCloudConnectionInvoker(request *model.UntagCloudConnectionRequest) *UntagCloudConnectionInvoker {
	requestDef := GenReqDefForUntagCloudConnection()
	return &UntagCloudConnectionInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// ListCloudConnectionQuotas 查询云连接配额
//
// 查询云连接配额。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListCloudConnectionQuotas(request *model.ListCloudConnectionQuotasRequest) (*model.ListCloudConnectionQuotasResponse, error) {
	requestDef := GenReqDefForListCloudConnectionQuotas()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCloudConnectionQuotasResponse), nil
	}
}

// ListCloudConnectionQuotasInvoker 查询云连接配额
func (c *CcClient) ListCloudConnectionQuotasInvoker(request *model.ListCloudConnectionQuotasRequest) *ListCloudConnectionQuotasInvoker {
	requestDef := GenReqDefForListCloudConnectionQuotas()
	return &ListCloudConnectionQuotasInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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

// BatchCreateGcbResourceTags 批量添加账户全域互联带宽资源标签
//
// # TMS批量添加资源标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) BatchCreateGcbResourceTags(request *model.BatchCreateGcbResourceTagsRequest) (*model.BatchCreateGcbResourceTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateGcbResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateGcbResourceTagsResponse), nil
	}
}

// BatchCreateGcbResourceTagsInvoker 批量添加账户全域互联带宽资源标签
func (c *CcClient) BatchCreateGcbResourceTagsInvoker(request *model.BatchCreateGcbResourceTagsRequest) *BatchCreateGcbResourceTagsInvoker {
	requestDef := GenReqDefForBatchCreateGcbResourceTags()
	return &BatchCreateGcbResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteGcbResourceTags 批量删除账户全域互联带宽资源标签
//
// 批量删除账户全域互联带宽资源标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) BatchDeleteGcbResourceTags(request *model.BatchDeleteGcbResourceTagsRequest) (*model.BatchDeleteGcbResourceTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteGcbResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteGcbResourceTagsResponse), nil
	}
}

// BatchDeleteGcbResourceTagsInvoker 批量删除账户全域互联带宽资源标签
func (c *CcClient) BatchDeleteGcbResourceTagsInvoker(request *model.BatchDeleteGcbResourceTagsRequest) *BatchDeleteGcbResourceTagsInvoker {
	requestDef := GenReqDefForBatchDeleteGcbResourceTags()
	return &BatchDeleteGcbResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CountGcbResourceByTag 查询账户全域互联带宽资源标签数量
//
// 查询账户全域互联带宽资源标签数量
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) CountGcbResourceByTag(request *model.CountGcbResourceByTagRequest) (*model.CountGcbResourceByTagResponse, error) {
	requestDef := GenReqDefForCountGcbResourceByTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CountGcbResourceByTagResponse), nil
	}
}

// CountGcbResourceByTagInvoker 查询账户全域互联带宽资源标签数量
func (c *CcClient) CountGcbResourceByTagInvoker(request *model.CountGcbResourceByTagRequest) *CountGcbResourceByTagInvoker {
	requestDef := GenReqDefForCountGcbResourceByTag()
	return &CountGcbResourceByTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGcbResourceTag 添加账户全域互联带宽资源标签
//
// 添加账户全域互联带宽资源标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) CreateGcbResourceTag(request *model.CreateGcbResourceTagRequest) (*model.CreateGcbResourceTagResponse, error) {
	requestDef := GenReqDefForCreateGcbResourceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGcbResourceTagResponse), nil
	}
}

// CreateGcbResourceTagInvoker 添加账户全域互联带宽资源标签
func (c *CcClient) CreateGcbResourceTagInvoker(request *model.CreateGcbResourceTagRequest) *CreateGcbResourceTagInvoker {
	requestDef := GenReqDefForCreateGcbResourceTag()
	return &CreateGcbResourceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGcbResourceTag 删除账户全域互联带宽资源标签
//
// 删除账户全域互联带宽资源标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) DeleteGcbResourceTag(request *model.DeleteGcbResourceTagRequest) (*model.DeleteGcbResourceTagResponse, error) {
	requestDef := GenReqDefForDeleteGcbResourceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGcbResourceTagResponse), nil
	}
}

// DeleteGcbResourceTagInvoker 删除账户全域互联带宽资源标签
func (c *CcClient) DeleteGcbResourceTagInvoker(request *model.DeleteGcbResourceTagRequest) *DeleteGcbResourceTagInvoker {
	requestDef := GenReqDefForDeleteGcbResourceTag()
	return &DeleteGcbResourceTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGcbResourceByTag 查询账户全域互联带宽资源实例列表
//
// 查询账户全域互联带宽资源实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListGcbResourceByTag(request *model.ListGcbResourceByTagRequest) (*model.ListGcbResourceByTagResponse, error) {
	requestDef := GenReqDefForListGcbResourceByTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGcbResourceByTagResponse), nil
	}
}

// ListGcbResourceByTagInvoker 查询账户全域互联带宽资源实例列表
func (c *CcClient) ListGcbResourceByTagInvoker(request *model.ListGcbResourceByTagRequest) *ListGcbResourceByTagInvoker {
	requestDef := GenReqDefForListGcbResourceByTag()
	return &ListGcbResourceByTagInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGcbResourceTags 查询账户全域互联带宽资源的标签
//
// 查询账户全域互联带宽资源的标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListGcbResourceTags(request *model.ListGcbResourceTagsRequest) (*model.ListGcbResourceTagsResponse, error) {
	requestDef := GenReqDefForListGcbResourceTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGcbResourceTagsResponse), nil
	}
}

// ListGcbResourceTagsInvoker 查询账户全域互联带宽资源的标签
func (c *CcClient) ListGcbResourceTagsInvoker(request *model.ListGcbResourceTagsRequest) *ListGcbResourceTagsInvoker {
	requestDef := GenReqDefForListGcbResourceTags()
	return &ListGcbResourceTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGcbTenantTags 查询账户全域互联带宽所有资源标签
//
// 查询账户全域互联带宽所有资源标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListGcbTenantTags(request *model.ListGcbTenantTagsRequest) (*model.ListGcbTenantTagsResponse, error) {
	requestDef := GenReqDefForListGcbTenantTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGcbTenantTagsResponse), nil
	}
}

// ListGcbTenantTagsInvoker 查询账户全域互联带宽所有资源标签
func (c *CcClient) ListGcbTenantTagsInvoker(request *model.ListGcbTenantTagsRequest) *ListGcbTenantTagsInvoker {
	requestDef := GenReqDefForListGcbTenantTags()
	return &ListGcbTenantTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AssociateGlobalConnectionBandwidthInstance 全域互联带宽绑定实例
//
// 全域互联带宽绑定实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) AssociateGlobalConnectionBandwidthInstance(request *model.AssociateGlobalConnectionBandwidthInstanceRequest) (*model.AssociateGlobalConnectionBandwidthInstanceResponse, error) {
	requestDef := GenReqDefForAssociateGlobalConnectionBandwidthInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateGlobalConnectionBandwidthInstanceResponse), nil
	}
}

// AssociateGlobalConnectionBandwidthInstanceInvoker 全域互联带宽绑定实例
func (c *CcClient) AssociateGlobalConnectionBandwidthInstanceInvoker(request *model.AssociateGlobalConnectionBandwidthInstanceRequest) *AssociateGlobalConnectionBandwidthInstanceInvoker {
	requestDef := GenReqDefForAssociateGlobalConnectionBandwidthInstance()
	return &AssociateGlobalConnectionBandwidthInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateGlobalConnectionBandwidth 创建全域互联带宽
//
// 创建全域互联带宽。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) CreateGlobalConnectionBandwidth(request *model.CreateGlobalConnectionBandwidthRequest) (*model.CreateGlobalConnectionBandwidthResponse, error) {
	requestDef := GenReqDefForCreateGlobalConnectionBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateGlobalConnectionBandwidthResponse), nil
	}
}

// CreateGlobalConnectionBandwidthInvoker 创建全域互联带宽
func (c *CcClient) CreateGlobalConnectionBandwidthInvoker(request *model.CreateGlobalConnectionBandwidthRequest) *CreateGlobalConnectionBandwidthInvoker {
	requestDef := GenReqDefForCreateGlobalConnectionBandwidth()
	return &CreateGlobalConnectionBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteGlobalConnectionBandwidth 删除全域互联带宽
//
// 删除全域互联带宽。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) DeleteGlobalConnectionBandwidth(request *model.DeleteGlobalConnectionBandwidthRequest) (*model.DeleteGlobalConnectionBandwidthResponse, error) {
	requestDef := GenReqDefForDeleteGlobalConnectionBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteGlobalConnectionBandwidthResponse), nil
	}
}

// DeleteGlobalConnectionBandwidthInvoker 删除全域互联带宽
func (c *CcClient) DeleteGlobalConnectionBandwidthInvoker(request *model.DeleteGlobalConnectionBandwidthRequest) *DeleteGlobalConnectionBandwidthInvoker {
	requestDef := GenReqDefForDeleteGlobalConnectionBandwidth()
	return &DeleteGlobalConnectionBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateGlobalConnectionBandwidthInstance 全域互联带宽解绑实例
//
// 全域互联带宽解绑实例。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) DisassociateGlobalConnectionBandwidthInstance(request *model.DisassociateGlobalConnectionBandwidthInstanceRequest) (*model.DisassociateGlobalConnectionBandwidthInstanceResponse, error) {
	requestDef := GenReqDefForDisassociateGlobalConnectionBandwidthInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateGlobalConnectionBandwidthInstanceResponse), nil
	}
}

// DisassociateGlobalConnectionBandwidthInstanceInvoker 全域互联带宽解绑实例
func (c *CcClient) DisassociateGlobalConnectionBandwidthInstanceInvoker(request *model.DisassociateGlobalConnectionBandwidthInstanceRequest) *DisassociateGlobalConnectionBandwidthInstanceInvoker {
	requestDef := GenReqDefForDisassociateGlobalConnectionBandwidthInstance()
	return &DisassociateGlobalConnectionBandwidthInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGlobalConnectionBandwidthConfigs 查询全域互联带宽租户配置信息
//
// 查询全域互联带宽租户配置信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListGlobalConnectionBandwidthConfigs(request *model.ListGlobalConnectionBandwidthConfigsRequest) (*model.ListGlobalConnectionBandwidthConfigsResponse, error) {
	requestDef := GenReqDefForListGlobalConnectionBandwidthConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGlobalConnectionBandwidthConfigsResponse), nil
	}
}

// ListGlobalConnectionBandwidthConfigsInvoker 查询全域互联带宽租户配置信息
func (c *CcClient) ListGlobalConnectionBandwidthConfigsInvoker(request *model.ListGlobalConnectionBandwidthConfigsRequest) *ListGlobalConnectionBandwidthConfigsInvoker {
	requestDef := GenReqDefForListGlobalConnectionBandwidthConfigs()
	return &ListGlobalConnectionBandwidthConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGlobalConnectionBandwidthLineLevels 查询线路等级列表
//
// 查询线路等级列表。
// 分页查询使用的参数为marker、limit。marker和limit一起使用时才会生效，单独使用无效。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListGlobalConnectionBandwidthLineLevels(request *model.ListGlobalConnectionBandwidthLineLevelsRequest) (*model.ListGlobalConnectionBandwidthLineLevelsResponse, error) {
	requestDef := GenReqDefForListGlobalConnectionBandwidthLineLevels()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGlobalConnectionBandwidthLineLevelsResponse), nil
	}
}

// ListGlobalConnectionBandwidthLineLevelsInvoker 查询线路等级列表
func (c *CcClient) ListGlobalConnectionBandwidthLineLevelsInvoker(request *model.ListGlobalConnectionBandwidthLineLevelsRequest) *ListGlobalConnectionBandwidthLineLevelsInvoker {
	requestDef := GenReqDefForListGlobalConnectionBandwidthLineLevels()
	return &ListGlobalConnectionBandwidthLineLevelsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGlobalConnectionBandwidthSites 查询站点列表
//
// 查询站点列表。
// 分页查询使用的参数为marker、limit。marker和limit一起使用时才会生效，单独使用无效。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListGlobalConnectionBandwidthSites(request *model.ListGlobalConnectionBandwidthSitesRequest) (*model.ListGlobalConnectionBandwidthSitesResponse, error) {
	requestDef := GenReqDefForListGlobalConnectionBandwidthSites()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGlobalConnectionBandwidthSitesResponse), nil
	}
}

// ListGlobalConnectionBandwidthSitesInvoker 查询站点列表
func (c *CcClient) ListGlobalConnectionBandwidthSitesInvoker(request *model.ListGlobalConnectionBandwidthSitesRequest) *ListGlobalConnectionBandwidthSitesInvoker {
	requestDef := GenReqDefForListGlobalConnectionBandwidthSites()
	return &ListGlobalConnectionBandwidthSitesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGlobalConnectionBandwidthSpecCodes 查询线路规格列表
//
// 查询线路规格列表。租户白名单控制，默认为空。
// 分页查询使用的参数为marker、limit。marker和limit一起使用时才会生效，单独使用无效。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListGlobalConnectionBandwidthSpecCodes(request *model.ListGlobalConnectionBandwidthSpecCodesRequest) (*model.ListGlobalConnectionBandwidthSpecCodesResponse, error) {
	requestDef := GenReqDefForListGlobalConnectionBandwidthSpecCodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGlobalConnectionBandwidthSpecCodesResponse), nil
	}
}

// ListGlobalConnectionBandwidthSpecCodesInvoker 查询线路规格列表
func (c *CcClient) ListGlobalConnectionBandwidthSpecCodesInvoker(request *model.ListGlobalConnectionBandwidthSpecCodesRequest) *ListGlobalConnectionBandwidthSpecCodesInvoker {
	requestDef := GenReqDefForListGlobalConnectionBandwidthSpecCodes()
	return &ListGlobalConnectionBandwidthSpecCodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListGlobalConnectionBandwidths 查询全域互联带宽列表
//
// 查询全域互联带宽列表。
// 分页查询使用的参数为marker、limit。marker和limit一起使用时才会生效，单独使用无效。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListGlobalConnectionBandwidths(request *model.ListGlobalConnectionBandwidthsRequest) (*model.ListGlobalConnectionBandwidthsResponse, error) {
	requestDef := GenReqDefForListGlobalConnectionBandwidths()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListGlobalConnectionBandwidthsResponse), nil
	}
}

// ListGlobalConnectionBandwidthsInvoker 查询全域互联带宽列表
func (c *CcClient) ListGlobalConnectionBandwidthsInvoker(request *model.ListGlobalConnectionBandwidthsRequest) *ListGlobalConnectionBandwidthsInvoker {
	requestDef := GenReqDefForListGlobalConnectionBandwidths()
	return &ListGlobalConnectionBandwidthsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListSupportBindingConnectionBandwidths 查询符合绑定条件的全域互联带宽列表
//
// 查询符合绑定条件的全域互联带宽列表。
// 分页查询使用的参数为marker、limit。marker和limit一起使用时才会生效，单独使用无效。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ListSupportBindingConnectionBandwidths(request *model.ListSupportBindingConnectionBandwidthsRequest) (*model.ListSupportBindingConnectionBandwidthsResponse, error) {
	requestDef := GenReqDefForListSupportBindingConnectionBandwidths()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSupportBindingConnectionBandwidthsResponse), nil
	}
}

// ListSupportBindingConnectionBandwidthsInvoker 查询符合绑定条件的全域互联带宽列表
func (c *CcClient) ListSupportBindingConnectionBandwidthsInvoker(request *model.ListSupportBindingConnectionBandwidthsRequest) *ListSupportBindingConnectionBandwidthsInvoker {
	requestDef := GenReqDefForListSupportBindingConnectionBandwidths()
	return &ListSupportBindingConnectionBandwidthsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowGlobalConnectionBandwidth 查询全域互联带宽详情
//
// 查询全域互联带宽详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) ShowGlobalConnectionBandwidth(request *model.ShowGlobalConnectionBandwidthRequest) (*model.ShowGlobalConnectionBandwidthResponse, error) {
	requestDef := GenReqDefForShowGlobalConnectionBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowGlobalConnectionBandwidthResponse), nil
	}
}

// ShowGlobalConnectionBandwidthInvoker 查询全域互联带宽详情
func (c *CcClient) ShowGlobalConnectionBandwidthInvoker(request *model.ShowGlobalConnectionBandwidthRequest) *ShowGlobalConnectionBandwidthInvoker {
	requestDef := GenReqDefForShowGlobalConnectionBandwidth()
	return &ShowGlobalConnectionBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateGlobalConnectionBandwidth 更新全域互联带宽详情
//
// 更新全域互联带宽详情。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CcClient) UpdateGlobalConnectionBandwidth(request *model.UpdateGlobalConnectionBandwidthRequest) (*model.UpdateGlobalConnectionBandwidthResponse, error) {
	requestDef := GenReqDefForUpdateGlobalConnectionBandwidth()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateGlobalConnectionBandwidthResponse), nil
	}
}

// UpdateGlobalConnectionBandwidthInvoker 更新全域互联带宽详情
func (c *CcClient) UpdateGlobalConnectionBandwidthInvoker(request *model.UpdateGlobalConnectionBandwidthRequest) *UpdateGlobalConnectionBandwidthInvoker {
	requestDef := GenReqDefForUpdateGlobalConnectionBandwidth()
	return &UpdateGlobalConnectionBandwidthInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
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
