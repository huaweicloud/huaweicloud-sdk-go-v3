package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cloudtable/v2/model"
)

type CloudTableClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCloudTableClient(hcClient *http_client.HcHttpClient) *CloudTableClient {
	return &CloudTableClient{HcClient: hcClient}
}

func CloudTableClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// CreateCluster 创建CloudTable集群
//
// 创建一个CloudTable集群。
// 使用接口前，您需要先获取如下资源信息。
// - 通过VPC创建或查询VPC、子网
// - 通过安全组创建或查询可用的security_group_id
//
// 本接口是一个同步接口，当创建CloudTable集群成功后会返回集群id。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudTableClient) CreateCluster(request *model.CreateClusterRequest) (*model.CreateClusterResponse, error) {
	requestDef := GenReqDefForCreateCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClusterResponse), nil
	}
}

// CreateClusterInvoker 创建CloudTable集群
func (c *CloudTableClient) CreateClusterInvoker(request *model.CreateClusterRequest) *CreateClusterInvoker {
	requestDef := GenReqDefForCreateCluster()
	return &CreateClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteCluster 删除CloudTable指定集群
//
// 集群ID为集群唯一标识，根据集群ID删除指定集群。
// 如下状态的集群不允许删除：
// - 创建中
// - 扩容中
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudTableClient) DeleteCluster(request *model.DeleteClusterRequest) (*model.DeleteClusterResponse, error) {
	requestDef := GenReqDefForDeleteCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteClusterResponse), nil
	}
}

// DeleteClusterInvoker 删除CloudTable指定集群
func (c *CloudTableClient) DeleteClusterInvoker(request *model.DeleteClusterRequest) *DeleteClusterInvoker {
	requestDef := GenReqDefForDeleteCluster()
	return &DeleteClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// EnableComponent 开启opentsdb组件
//
// 开启opentsdb组件
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudTableClient) EnableComponent(request *model.EnableComponentRequest) (*model.EnableComponentResponse, error) {
	requestDef := GenReqDefForEnableComponent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.EnableComponentResponse), nil
	}
}

// EnableComponentInvoker 开启opentsdb组件
func (c *CloudTableClient) EnableComponentInvoker(request *model.EnableComponentRequest) *EnableComponentInvoker {
	requestDef := GenReqDefForEnableComponent()
	return &EnableComponentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ExpandClusterComponent 扩容组件
//
// 扩容指定类型的集群节点
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudTableClient) ExpandClusterComponent(request *model.ExpandClusterComponentRequest) (*model.ExpandClusterComponentResponse, error) {
	requestDef := GenReqDefForExpandClusterComponent()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ExpandClusterComponentResponse), nil
	}
}

// ExpandClusterComponentInvoker 扩容组件
func (c *CloudTableClient) ExpandClusterComponentInvoker(request *model.ExpandClusterComponentRequest) *ExpandClusterComponentInvoker {
	requestDef := GenReqDefForExpandClusterComponent()
	return &ExpandClusterComponentInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListClusters 查询CloudTable集群列表
//
// 查看用户创建的集群列表信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudTableClient) ListClusters(request *model.ListClustersRequest) (*model.ListClustersResponse, error) {
	requestDef := GenReqDefForListClusters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClustersResponse), nil
	}
}

// ListClustersInvoker 查询CloudTable集群列表
func (c *CloudTableClient) ListClustersInvoker(request *model.ListClustersRequest) *ListClustersInvoker {
	requestDef := GenReqDefForListClusters()
	return &ListClustersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// RebootCloudTableCluster 重启集群的api入口
//
// 重启集群的api入口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudTableClient) RebootCloudTableCluster(request *model.RebootCloudTableClusterRequest) (*model.RebootCloudTableClusterResponse, error) {
	requestDef := GenReqDefForRebootCloudTableCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RebootCloudTableClusterResponse), nil
	}
}

// RebootCloudTableClusterInvoker 重启集群的api入口
func (c *CloudTableClient) RebootCloudTableClusterInvoker(request *model.RebootCloudTableClusterRequest) *RebootCloudTableClusterInvoker {
	requestDef := GenReqDefForRebootCloudTableCluster()
	return &RebootCloudTableClusterInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowClusterDetail 查询CloudTable集群详情
//
// 通过集群ID唯一标识一个集群，根据集群ID查询特定CloudTable集群的详情信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudTableClient) ShowClusterDetail(request *model.ShowClusterDetailRequest) (*model.ShowClusterDetailResponse, error) {
	requestDef := GenReqDefForShowClusterDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClusterDetailResponse), nil
	}
}

// ShowClusterDetailInvoker 查询CloudTable集群详情
func (c *CloudTableClient) ShowClusterDetailInvoker(request *model.ShowClusterDetailRequest) *ShowClusterDetailInvoker {
	requestDef := GenReqDefForShowClusterDetail()
	return &ShowClusterDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowClusterSetting 查询集群配置
//
// 查询集群配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudTableClient) ShowClusterSetting(request *model.ShowClusterSettingRequest) (*model.ShowClusterSettingResponse, error) {
	requestDef := GenReqDefForShowClusterSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClusterSettingResponse), nil
	}
}

// ShowClusterSettingInvoker 查询集群配置
func (c *CloudTableClient) ShowClusterSettingInvoker(request *model.ShowClusterSettingRequest) *ShowClusterSettingInvoker {
	requestDef := GenReqDefForShowClusterSetting()
	return &ShowClusterSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateClusterSetting 修改集群配置
//
// 修改集群配置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CloudTableClient) UpdateClusterSetting(request *model.UpdateClusterSettingRequest) (*model.UpdateClusterSettingResponse, error) {
	requestDef := GenReqDefForUpdateClusterSetting()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateClusterSettingResponse), nil
	}
}

// UpdateClusterSettingInvoker 修改集群配置
func (c *CloudTableClient) UpdateClusterSettingInvoker(request *model.UpdateClusterSettingRequest) *UpdateClusterSettingInvoker {
	requestDef := GenReqDefForUpdateClusterSetting()
	return &UpdateClusterSettingInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
