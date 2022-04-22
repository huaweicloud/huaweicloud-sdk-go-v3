package v3

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

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

// 创建云连接实例
//
// 创建云连接实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcClient) CreateCloudConnection(request *model.CreateCloudConnectionRequest) (*model.CreateCloudConnectionResponse, error) {
	requestDef := GenReqDefForCreateCloudConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCloudConnectionResponse), nil
	}
}

// 创建网络实例
//
// 创建网络实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcClient) CreateNetworkInstance(request *model.CreateNetworkInstanceRequest) (*model.CreateNetworkInstanceResponse, error) {
	requestDef := GenReqDefForCreateNetworkInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateNetworkInstanceResponse), nil
	}
}

// 删除云连接实例
//
// 删除云连接实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcClient) DeleteCloudConnection(request *model.DeleteCloudConnectionRequest) (*model.DeleteCloudConnectionResponse, error) {
	requestDef := GenReqDefForDeleteCloudConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCloudConnectionResponse), nil
	}
}

// 删除网络实例
//
// 删除网络实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcClient) DeleteNetworkInstance(request *model.DeleteNetworkInstanceRequest) (*model.DeleteNetworkInstanceResponse, error) {
	requestDef := GenReqDefForDeleteNetworkInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteNetworkInstanceResponse), nil
	}
}

// 查询云连接路由条目列表
//
// 查询云连接路由条目列表。
// 分页查询使用的参数为marker、limit。marker和limit一起使用时才会生效，单独使用无效。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcClient) ListCloudConnectionRoutes(request *model.ListCloudConnectionRoutesRequest) (*model.ListCloudConnectionRoutesResponse, error) {
	requestDef := GenReqDefForListCloudConnectionRoutes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCloudConnectionRoutesResponse), nil
	}
}

// 查询云连接列表
//
// 查询云连接列表。
// 分页查询使用的参数为marker、limit。marker和limit一起使用时才会生效，单独使用无效。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcClient) ListCloudConnections(request *model.ListCloudConnectionsRequest) (*model.ListCloudConnectionsResponse, error) {
	requestDef := GenReqDefForListCloudConnections()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCloudConnectionsResponse), nil
	}
}

// 查询网络实例列表
//
// 查询云连接列表。
// 分页查询使用的参数为marker、limit。marker和limit一起使用时才会生效，单独使用无效。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcClient) ListNetworkInstances(request *model.ListNetworkInstancesRequest) (*model.ListNetworkInstancesResponse, error) {
	requestDef := GenReqDefForListNetworkInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNetworkInstancesResponse), nil
	}
}

// 查询云连接实例
//
// 查询云连接实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcClient) ShowCloudConnection(request *model.ShowCloudConnectionRequest) (*model.ShowCloudConnectionResponse, error) {
	requestDef := GenReqDefForShowCloudConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCloudConnectionResponse), nil
	}
}

// 查询云连接路由条目详情
//
// 查询云连接路由条目列表。
// 分页查询使用的参数为marker、limit。marker和limit一起使用时才会生效，单独使用无效。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcClient) ShowCloudConnectionRoutes(request *model.ShowCloudConnectionRoutesRequest) (*model.ShowCloudConnectionRoutesResponse, error) {
	requestDef := GenReqDefForShowCloudConnectionRoutes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowCloudConnectionRoutesResponse), nil
	}
}

// 查询网络实例
//
// 查询网络实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcClient) ShowNetworkInstance(request *model.ShowNetworkInstanceRequest) (*model.ShowNetworkInstanceResponse, error) {
	requestDef := GenReqDefForShowNetworkInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNetworkInstanceResponse), nil
	}
}

// 更新云连接实例
//
// 更新云连接实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcClient) UpdateCloudConnection(request *model.UpdateCloudConnectionRequest) (*model.UpdateCloudConnectionResponse, error) {
	requestDef := GenReqDefForUpdateCloudConnection()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCloudConnectionResponse), nil
	}
}

// 更新网络实例
//
// 更新网络实例。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *CcClient) UpdateNetworkInstance(request *model.UpdateNetworkInstanceRequest) (*model.UpdateNetworkInstanceResponse, error) {
	requestDef := GenReqDefForUpdateNetworkInstance()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateNetworkInstanceResponse), nil
	}
}
