package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dws/v2/model"
)

type DwsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewDwsClient(hcClient *http_client.HcHttpClient) *DwsClient {
	return &DwsClient{HcClient: hcClient}
}

func DwsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// 创建集群
//
// 该接口用于创建集群。
// 集群必须要运行在VPC之内，创建集群前，您需要先创建VPC，并获取VPC和子网的
// id。
// 该接口为异步接口，创建集群需要10～15分钟。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DwsClient) CreateCluster(request *model.CreateClusterRequest) (*model.CreateClusterResponse, error) {
	requestDef := GenReqDefForCreateCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClusterResponse), nil
	}
}

// 创建快照
//
// 该接口用于为指定集群创建快照。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DwsClient) CreateSnapshot(request *model.CreateSnapshotRequest) (*model.CreateSnapshotResponse, error) {
	requestDef := GenReqDefForCreateSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateSnapshotResponse), nil
	}
}

// 删除集群
//
// 此接口用于删除集群。集群删除后将释放此集群的所有资源，包括客户数据。为了安全起见，请在删除集群前为这个集群创建快照。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DwsClient) DeleteCluster(request *model.DeleteClusterRequest) (*model.DeleteClusterResponse, error) {
	requestDef := GenReqDefForDeleteCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteClusterResponse), nil
	}
}

// 删除快照
//
// 该接口用于删除一个指定手动快照。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DwsClient) DeleteSnapshot(request *model.DeleteSnapshotRequest) (*model.DeleteSnapshotResponse, error) {
	requestDef := GenReqDefForDeleteSnapshot()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteSnapshotResponse), nil
	}
}

// 查询集群详情
//
// 该接口用于查询集群详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DwsClient) ListClusterDetails(request *model.ListClusterDetailsRequest) (*model.ListClusterDetailsResponse, error) {
	requestDef := GenReqDefForListClusterDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClusterDetailsResponse), nil
	}
}

// 查询集群列表
//
// 该接口用于查询并显示集群列表
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DwsClient) ListClusters(request *model.ListClustersRequest) (*model.ListClustersResponse, error) {
	requestDef := GenReqDefForListClusters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClustersResponse), nil
	}
}

// 查询节点类型
//
// 该接口用于查询所有GaussDB(DWS)服务支持的节点类型。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DwsClient) ListNodeTypes(request *model.ListNodeTypesRequest) (*model.ListNodeTypesResponse, error) {
	requestDef := GenReqDefForListNodeTypes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNodeTypesResponse), nil
	}
}

// 查询快照详情
//
// 该接口用于使用快照ID查询快照详情。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DwsClient) ListSnapshotDetails(request *model.ListSnapshotDetailsRequest) (*model.ListSnapshotDetailsResponse, error) {
	requestDef := GenReqDefForListSnapshotDetails()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSnapshotDetailsResponse), nil
	}
}

// 查询快照列表
//
// 该接口用于查询快照列表。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DwsClient) ListSnapshots(request *model.ListSnapshotsRequest) (*model.ListSnapshotsResponse, error) {
	requestDef := GenReqDefForListSnapshots()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListSnapshotsResponse), nil
	}
}

// 重置密码
//
// 此接口用于重置集群管理员密码。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DwsClient) ResetPassword(request *model.ResetPasswordRequest) (*model.ResetPasswordResponse, error) {
	requestDef := GenReqDefForResetPassword()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResetPasswordResponse), nil
	}
}

// 扩容集群调整集群大小
//
// 此接口用于扩容集群。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DwsClient) ResizeCluster(request *model.ResizeClusterRequest) (*model.ResizeClusterResponse, error) {
	requestDef := GenReqDefForResizeCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ResizeClusterResponse), nil
	}
}

// 重启集群
//
// 此接口用于重启集群。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DwsClient) RestartCluster(request *model.RestartClusterRequest) (*model.RestartClusterResponse, error) {
	requestDef := GenReqDefForRestartCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestartClusterResponse), nil
	}
}

// 恢复集群
//
// 该接口用于使用快照恢复集群。
//
// 详细说明请参考华为云API Explorer。
// Please refer to Huawei cloud API Explorer for details.
func (c *DwsClient) RestoreCluster(request *model.RestoreClusterRequest) (*model.RestoreClusterResponse, error) {
	requestDef := GenReqDefForRestoreCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestoreClusterResponse), nil
	}
}
