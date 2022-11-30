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
