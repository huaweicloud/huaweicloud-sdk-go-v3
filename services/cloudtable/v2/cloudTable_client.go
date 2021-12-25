package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

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

//创建一个CloudTable集群。 使用接口前，您需要先获取如下资源信息。 - 通过VPC创建或查询VPC、子网 - 通过安全组创建或查询可用的security_group_id  本接口是一个同步接口，当创建CloudTable集群成功后会返回集群id。
func (c *CloudTableClient) CreateCluster(request *model.CreateClusterRequest) (*model.CreateClusterResponse, error) {
	requestDef := GenReqDefForCreateCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClusterResponse), nil
	}
}

//集群ID为集群唯一标识，根据集群ID删除指定集群。 如下状态的集群不允许删除： - 创建中 - 扩容中
func (c *CloudTableClient) DeleteCluster(request *model.DeleteClusterRequest) (*model.DeleteClusterResponse, error) {
	requestDef := GenReqDefForDeleteCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteClusterResponse), nil
	}
}

//查看用户创建的集群列表信息。
func (c *CloudTableClient) ListClusters(request *model.ListClustersRequest) (*model.ListClustersResponse, error) {
	requestDef := GenReqDefForListClusters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClustersResponse), nil
	}
}

//通过集群ID唯一标识一个集群，根据集群ID查询特定CloudTable集群的详情信息。
func (c *CloudTableClient) ShowClusterDetail(request *model.ShowClusterDetailRequest) (*model.ShowClusterDetailResponse, error) {
	requestDef := GenReqDefForShowClusterDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowClusterDetailResponse), nil
	}
}
