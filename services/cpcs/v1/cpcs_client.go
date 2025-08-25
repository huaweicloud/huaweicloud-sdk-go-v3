package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cpcs/v1/model"
)

type CpcsClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewCpcsClient(hcClient *httpclient.HcHttpClient) *CpcsClient {
	return &CpcsClient{HcClient: hcClient}
}

func CpcsClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// AssociateApps 创建密码服务集群与应用绑定关系
//
// 创建密码服务集群与应用绑定关系
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CpcsClient) AssociateApps(request *model.AssociateAppsRequest) (*model.AssociateAppsResponse, error) {
	requestDef := GenReqDefForAssociateApps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateAppsResponse), nil
	}
}

// AssociateAppsInvoker 创建密码服务集群与应用绑定关系
func (c *CpcsClient) AssociateAppsInvoker(request *model.AssociateAppsRequest) *AssociateAppsInvoker {
	requestDef := GenReqDefForAssociateApps()
	return &AssociateAppsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AuthorizeAccessKeys 密码服务集群授予应用访问密钥的访问权限
//
// 密码服务集群授予应用访问密钥的访问权限
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CpcsClient) AuthorizeAccessKeys(request *model.AuthorizeAccessKeysRequest) (*model.AuthorizeAccessKeysResponse, error) {
	requestDef := GenReqDefForAuthorizeAccessKeys()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AuthorizeAccessKeysResponse), nil
	}
}

// AuthorizeAccessKeysInvoker 密码服务集群授予应用访问密钥的访问权限
func (c *CpcsClient) AuthorizeAccessKeysInvoker(request *model.AuthorizeAccessKeysRequest) *AuthorizeAccessKeysInvoker {
	requestDef := GenReqDefForAuthorizeAccessKeys()
	return &AuthorizeAccessKeysInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDisableAccessKeys 停用应用的访问密钥
//
// 停用应用的访问密钥
// &gt; 只有当访问密钥处于“启用”状态时，方可调用此接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CpcsClient) BatchDisableAccessKeys(request *model.BatchDisableAccessKeysRequest) (*model.BatchDisableAccessKeysResponse, error) {
	requestDef := GenReqDefForBatchDisableAccessKeys()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDisableAccessKeysResponse), nil
	}
}

// BatchDisableAccessKeysInvoker 停用应用的访问密钥
func (c *CpcsClient) BatchDisableAccessKeysInvoker(request *model.BatchDisableAccessKeysRequest) *BatchDisableAccessKeysInvoker {
	requestDef := GenReqDefForBatchDisableAccessKeys()
	return &BatchDisableAccessKeysInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchEnableAccessKeys 启用应用的访问密钥
//
// 启用应用的访问密钥
// &gt; 只有当访问密钥处于“停用”状态时，方可调用此接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CpcsClient) BatchEnableAccessKeys(request *model.BatchEnableAccessKeysRequest) (*model.BatchEnableAccessKeysResponse, error) {
	requestDef := GenReqDefForBatchEnableAccessKeys()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchEnableAccessKeysResponse), nil
	}
}

// BatchEnableAccessKeysInvoker 启用应用的访问密钥
func (c *CpcsClient) BatchEnableAccessKeysInvoker(request *model.BatchEnableAccessKeysRequest) *BatchEnableAccessKeysInvoker {
	requestDef := GenReqDefForBatchEnableAccessKeys()
	return &BatchEnableAccessKeysInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CancelAuthorizeAccessKeys 密码服务集群解除对访问密钥的授权
//
// 密码服务集群解除对访问密钥的授权
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CpcsClient) CancelAuthorizeAccessKeys(request *model.CancelAuthorizeAccessKeysRequest) (*model.CancelAuthorizeAccessKeysResponse, error) {
	requestDef := GenReqDefForCancelAuthorizeAccessKeys()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CancelAuthorizeAccessKeysResponse), nil
	}
}

// CancelAuthorizeAccessKeysInvoker 密码服务集群解除对访问密钥的授权
func (c *CpcsClient) CancelAuthorizeAccessKeysInvoker(request *model.CancelAuthorizeAccessKeysRequest) *CancelAuthorizeAccessKeysInvoker {
	requestDef := GenReqDefForCancelAuthorizeAccessKeys()
	return &CancelAuthorizeAccessKeysInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateApp 创建应用
//
// 创建应用
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CpcsClient) CreateApp(request *model.CreateAppRequest) (*model.CreateAppResponse, error) {
	requestDef := GenReqDefForCreateApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppResponse), nil
	}
}

// CreateAppInvoker 创建应用
func (c *CpcsClient) CreateAppInvoker(request *model.CreateAppRequest) *CreateAppInvoker {
	requestDef := GenReqDefForCreateApp()
	return &CreateAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateAppAccessKey 创建访问密钥
//
// 创建访问密钥
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CpcsClient) CreateAppAccessKey(request *model.CreateAppAccessKeyRequest) (*model.CreateAppAccessKeyResponse, error) {
	requestDef := GenReqDefForCreateAppAccessKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateAppAccessKeyResponse), nil
	}
}

// CreateAppAccessKeyInvoker 创建访问密钥
func (c *CpcsClient) CreateAppAccessKeyInvoker(request *model.CreateAppAccessKeyRequest) *CreateAppAccessKeyInvoker {
	requestDef := GenReqDefForCreateAppAccessKey()
	return &CreateAppAccessKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// CreateCluster 创建密码服务集群
//
// 创建密码服务集群
// &gt; 调用接口之后返回订单ID，需要到“待支付订单里面”支付成功之后才能创建密码服务集群。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CpcsClient) CreateCluster(request *model.CreateClusterRequest) (*model.CreateClusterResponse, error) {
	requestDef := GenReqDefForCreateCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClusterResponse), nil
	}
}

// CreateClusterInvoker 创建密码服务集群
func (c *CpcsClient) CreateClusterInvoker(request *model.CreateClusterRequest) *CreateClusterInvoker {
	requestDef := GenReqDefForCreateCluster()
	return &CreateClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAccessKey 删除应用的访问密钥
//
// 删除应用的访问密钥
// &gt; 只有当访问密钥处于“停用”状态时，方可调用此接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CpcsClient) DeleteAccessKey(request *model.DeleteAccessKeyRequest) (*model.DeleteAccessKeyResponse, error) {
	requestDef := GenReqDefForDeleteAccessKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAccessKeyResponse), nil
	}
}

// DeleteAccessKeyInvoker 删除应用的访问密钥
func (c *CpcsClient) DeleteAccessKeyInvoker(request *model.DeleteAccessKeyRequest) *DeleteAccessKeyInvoker {
	requestDef := GenReqDefForDeleteAccessKey()
	return &DeleteAccessKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteApp 删除应用
//
// 删除应用
// &gt; 只有当应用与任何其它服务没有绑定关系的情况下，方可进行集群删除操作该操作不可恢复，请谨慎执行
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CpcsClient) DeleteApp(request *model.DeleteAppRequest) (*model.DeleteAppResponse, error) {
	requestDef := GenReqDefForDeleteApp()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAppResponse), nil
	}
}

// DeleteAppInvoker 删除应用
func (c *CpcsClient) DeleteAppInvoker(request *model.DeleteAppRequest) *DeleteAppInvoker {
	requestDef := GenReqDefForDeleteApp()
	return &DeleteAppInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCcspCluster 删除密码服务集群
//
// 删除密码服务集群，即释放密码服务集群所有服务实例以及对应的VSM集群实例，并删除集群相关信息
// &gt; 只有当集群与任何应用都没有绑定关系的情况下，且处于“运行中”状态，且退订了集群中所有实例，方可进行集群删除操作，该操作不可恢复，请谨慎执行
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CpcsClient) DeleteCcspCluster(request *model.DeleteCcspClusterRequest) (*model.DeleteCcspClusterResponse, error) {
	requestDef := GenReqDefForDeleteCcspCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCcspClusterResponse), nil
	}
}

// DeleteCcspClusterInvoker 删除密码服务集群
func (c *CpcsClient) DeleteCcspClusterInvoker(request *model.DeleteCcspClusterRequest) *DeleteCcspClusterInvoker {
	requestDef := GenReqDefForDeleteCcspCluster()
	return &DeleteCcspClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisableCcspInstance 停用密码服务实例的业务功能
//
// 停用密码服务实例的业务功能
// &gt; 只有当密码服务实例处于“启用”状态时，方可调用此接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CpcsClient) DisableCcspInstance(request *model.DisableCcspInstanceRequest) (*model.DisableCcspInstanceResponse, error) {
	requestDef := GenReqDefForDisableCcspInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisableCcspInstanceResponse), nil
	}
}

// DisableCcspInstanceInvoker 停用密码服务实例的业务功能
func (c *CpcsClient) DisableCcspInstanceInvoker(request *model.DisableCcspInstanceRequest) *DisableCcspInstanceInvoker {
	requestDef := GenReqDefForDisableCcspInstance()
	return &DisableCcspInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DisassociateApps 解除密码服务集群与应用绑定关系
//
// 解除密码服务集群与应用绑定关系
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CpcsClient) DisassociateApps(request *model.DisassociateAppsRequest) (*model.DisassociateAppsResponse, error) {
	requestDef := GenReqDefForDisassociateApps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateAppsResponse), nil
	}
}

// DisassociateAppsInvoker 解除密码服务集群与应用绑定关系
func (c *CpcsClient) DisassociateAppsInvoker(request *model.DisassociateAppsRequest) *DisassociateAppsInvoker {
	requestDef := GenReqDefForDisassociateApps()
	return &DisassociateAppsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableCcspInstance 启用密码服务实例的业务功能
//
// 启用密码服务实例的业务功能
// &gt; 只有当密码服务实例处于“停用”状态时，方可调用此接口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CpcsClient) EnableCcspInstance(request *model.EnableCcspInstanceRequest) (*model.EnableCcspInstanceResponse, error) {
	requestDef := GenReqDefForEnableCcspInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableCcspInstanceResponse), nil
	}
}

// EnableCcspInstanceInvoker 启用密码服务实例的业务功能
func (c *CpcsClient) EnableCcspInstanceInvoker(request *model.EnableCcspInstanceRequest) *EnableCcspInstanceInvoker {
	requestDef := GenReqDefForEnableCcspInstance()
	return &EnableCcspInstanceInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListCcspTenantImages 查询密码服务的镜像
//
// 查询密码服务的镜像
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CpcsClient) ListCcspTenantImages(request *model.ListCcspTenantImagesRequest) (*model.ListCcspTenantImagesResponse, error) {
	requestDef := GenReqDefForListCcspTenantImages()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCcspTenantImagesResponse), nil
	}
}

// ListCcspTenantImagesInvoker 查询密码服务的镜像
func (c *CpcsClient) ListCcspTenantImagesInvoker(request *model.ListCcspTenantImagesRequest) *ListCcspTenantImagesInvoker {
	requestDef := GenReqDefForListCcspTenantImages()
	return &ListCcspTenantImagesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAccessKey 下载访问密钥
//
// 下载访问密钥且只能下载一次。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CpcsClient) ShowAccessKey(request *model.ShowAccessKeyRequest) (*model.ShowAccessKeyResponse, error) {
	requestDef := GenReqDefForShowAccessKey()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAccessKeyResponse), nil
	}
}

// ShowAccessKeyInvoker 下载访问密钥
func (c *CpcsClient) ShowAccessKeyInvoker(request *model.ShowAccessKeyRequest) *ShowAccessKeyInvoker {
	requestDef := GenReqDefForShowAccessKey()
	return &ShowAccessKeyInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAppAccessKeyList 查询应用的访问密钥列表
//
// 查询应用的访问密钥列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CpcsClient) ShowAppAccessKeyList(request *model.ShowAppAccessKeyListRequest) (*model.ShowAppAccessKeyListResponse, error) {
	requestDef := GenReqDefForShowAppAccessKeyList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppAccessKeyListResponse), nil
	}
}

// ShowAppAccessKeyListInvoker 查询应用的访问密钥列表
func (c *CpcsClient) ShowAppAccessKeyListInvoker(request *model.ShowAppAccessKeyListRequest) *ShowAppAccessKeyListInvoker {
	requestDef := GenReqDefForShowAppAccessKeyList()
	return &ShowAppAccessKeyListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAppList 查询应用列表
//
// 查询应用列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CpcsClient) ShowAppList(request *model.ShowAppListRequest) (*model.ShowAppListResponse, error) {
	requestDef := GenReqDefForShowAppList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAppListResponse), nil
	}
}

// ShowAppListInvoker 查询应用列表
func (c *CpcsClient) ShowAppListInvoker(request *model.ShowAppListRequest) *ShowAppListInvoker {
	requestDef := GenReqDefForShowAppList()
	return &ShowAppListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAssociationList 查询密码服务集群与应用的绑定关系列表
//
// 查询密码服务集群与应用的绑定关系列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CpcsClient) ShowAssociationList(request *model.ShowAssociationListRequest) (*model.ShowAssociationListResponse, error) {
	requestDef := GenReqDefForShowAssociationList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAssociationListResponse), nil
	}
}

// ShowAssociationListInvoker 查询密码服务集群与应用的绑定关系列表
func (c *CpcsClient) ShowAssociationListInvoker(request *model.ShowAssociationListRequest) *ShowAssociationListInvoker {
	requestDef := GenReqDefForShowAssociationList()
	return &ShowAssociationListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowAvailableAz 查询可创建密码服务集群的可用区列表
//
// 查询可创建密码服务集群的可用区列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CpcsClient) ShowAvailableAz(request *model.ShowAvailableAzRequest) (*model.ShowAvailableAzResponse, error) {
	requestDef := GenReqDefForShowAvailableAz()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAvailableAzResponse), nil
	}
}

// ShowAvailableAzInvoker 查询可创建密码服务集群的可用区列表
func (c *CpcsClient) ShowAvailableAzInvoker(request *model.ShowAvailableAzRequest) *ShowAvailableAzInvoker {
	requestDef := GenReqDefForShowAvailableAz()
	return &ShowAvailableAzInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCcspCluster 查询密码服务集群详情
//
// 查询密码服务集群信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CpcsClient) ShowCcspCluster(request *model.ShowCcspClusterRequest) (*model.ShowCcspClusterResponse, error) {
	requestDef := GenReqDefForShowCcspCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCcspClusterResponse), nil
	}
}

// ShowCcspClusterInvoker 查询密码服务集群详情
func (c *CpcsClient) ShowCcspClusterInvoker(request *model.ShowCcspClusterRequest) *ShowCcspClusterInvoker {
	requestDef := GenReqDefForShowCcspCluster()
	return &ShowCcspClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCcspClusterList 查询密码服务集群列表
//
// 查询密码服务集群列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CpcsClient) ShowCcspClusterList(request *model.ShowCcspClusterListRequest) (*model.ShowCcspClusterListResponse, error) {
	requestDef := GenReqDefForShowCcspClusterList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCcspClusterListResponse), nil
	}
}

// ShowCcspClusterListInvoker 查询密码服务集群列表
func (c *CpcsClient) ShowCcspClusterListInvoker(request *model.ShowCcspClusterListRequest) *ShowCcspClusterListInvoker {
	requestDef := GenReqDefForShowCcspClusterList()
	return &ShowCcspClusterListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowCcspInstanceInfo 查询密码服务实例列表
//
// 查询密码服务实例列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CpcsClient) ShowCcspInstanceInfo(request *model.ShowCcspInstanceInfoRequest) (*model.ShowCcspInstanceInfoResponse, error) {
	requestDef := GenReqDefForShowCcspInstanceInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCcspInstanceInfoResponse), nil
	}
}

// ShowCcspInstanceInfoInvoker 查询密码服务实例列表
func (c *CpcsClient) ShowCcspInstanceInfoInvoker(request *model.ShowCcspInstanceInfoRequest) *ShowCcspInstanceInfoInvoker {
	requestDef := GenReqDefForShowCcspInstanceInfo()
	return &ShowCcspInstanceInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowClusterAccessKeyList 查询密码服务集群已授权的访问密钥列表
//
// 查询密码服务集群已授权的访问密钥列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CpcsClient) ShowClusterAccessKeyList(request *model.ShowClusterAccessKeyListRequest) (*model.ShowClusterAccessKeyListResponse, error) {
	requestDef := GenReqDefForShowClusterAccessKeyList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClusterAccessKeyListResponse), nil
	}
}

// ShowClusterAccessKeyListInvoker 查询密码服务集群已授权的访问密钥列表
func (c *CpcsClient) ShowClusterAccessKeyListInvoker(request *model.ShowClusterAccessKeyListRequest) *ShowClusterAccessKeyListInvoker {
	requestDef := GenReqDefForShowClusterAccessKeyList()
	return &ShowClusterAccessKeyListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowClusterUri 获取密码服务管理界面URL
//
// 获取密码服务管理界面URL
// &gt; URL存在有效期，URL失效后无法跳转管理界面，需要重新获取URL
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CpcsClient) ShowClusterUri(request *model.ShowClusterUriRequest) (*model.ShowClusterUriResponse, error) {
	requestDef := GenReqDefForShowClusterUri()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClusterUriResponse), nil
	}
}

// ShowClusterUriInvoker 获取密码服务管理界面URL
func (c *CpcsClient) ShowClusterUriInvoker(request *model.ShowClusterUriRequest) *ShowClusterUriInvoker {
	requestDef := GenReqDefForShowClusterUri()
	return &ShowClusterUriInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowResourceInfo 查询租户的资源分布信息
//
// 查询租户的资源分布信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CpcsClient) ShowResourceInfo(request *model.ShowResourceInfoRequest) (*model.ShowResourceInfoResponse, error) {
	requestDef := GenReqDefForShowResourceInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceInfoResponse), nil
	}
}

// ShowResourceInfoInvoker 查询租户的资源分布信息
func (c *CpcsClient) ShowResourceInfoInvoker(request *model.ShowResourceInfoRequest) *ShowResourceInfoInvoker {
	requestDef := GenReqDefForShowResourceInfo()
	return &ShowResourceInfoInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowVmMonitor 密码资源指标监控
//
// 获取密码服务实例与虚拟密码机实例的指标（cpu使用率，内存使用率等指标）。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CpcsClient) ShowVmMonitor(request *model.ShowVmMonitorRequest) (*model.ShowVmMonitorResponse, error) {
	requestDef := GenReqDefForShowVmMonitor()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowVmMonitorResponse), nil
	}
}

// ShowVmMonitorInvoker 密码资源指标监控
func (c *CpcsClient) ShowVmMonitorInvoker(request *model.ShowVmMonitorRequest) *ShowVmMonitorInvoker {
	requestDef := GenReqDefForShowVmMonitor()
	return &ShowVmMonitorInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
