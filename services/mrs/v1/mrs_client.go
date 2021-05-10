package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/mrs/v1/model"
)

type MrsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewMrsClient(hcClient *http_client.HcHttpClient) *MrsClient {
	return &MrsClient{HcClient: hcClient}
}

func MrsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//使用标签过滤集群。  集群默认按照创建时间倒序，集群tag也按照创建时间倒序。
func (c *MrsClient) ListClustersByTags(request *model.ListClustersByTagsRequest) (*model.ListClustersByTagsResponse, error) {
	requestDef := GenReqDefForListClustersByTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClustersByTagsResponse), nil
	}
}

//为指定集群批量添加标签。  一个集群上最多有10个标签。  此接口为幂等接口：  - 创建时，同一个集群不允许重复key，如果数据库存在就覆盖。
func (c *MrsClient) BatchCreateClusterTags(request *model.BatchCreateClusterTagsRequest) (*model.BatchCreateClusterTagsResponse, error) {
	requestDef := GenReqDefForBatchCreateClusterTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateClusterTagsResponse), nil
	}
}

//为指定集群批量删除标签。  一个集群上最多有10个标签。  此接口为幂等接口：   - 删除时，如果删除的标签不存在，默认处理成功，删除时不对标签字符集范围做校验。Key长度36个unicode字符，value为43个unicode字符。删除时tags结构体不能缺失，key不能为空，或者空字符串。
func (c *MrsClient) BatchDeleteClusterTags(request *model.BatchDeleteClusterTagsRequest) (*model.BatchDeleteClusterTagsResponse, error) {
	requestDef := GenReqDefForBatchDeleteClusterTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteClusterTagsResponse), nil
	}
}

//创建一个MRS集群，并在集群中提交一个作业。该接口不兼容Sahara。 支持同一时间并发创建10个集群。 使用接口前，您需要先获取如[表1](mrs_02_0028.xml#mrs_02_0028__tbbd2986d18874f82a8ab886ac25a57f8)所示的资源信息。 约束说明： 1.针对MRS 1.8.0以前的版本，仅当“safe_mode”配置为“1”时需要配置cluster_admin_secret。 针对MRS 1.8.0及以后版本，cluster_admin_secret为必选参数，不受参数“safe_mode”配置的影响。 2.集群登录方式有密码和密钥对两种，两者必选其一。 使用密码方式需要配置访问集群节点的root密码，即cluster_master_secret。 使用密钥对方式需要配置密钥对名称，即node_public_cert_name 3.磁盘参数可以使用volume_type和volume_size表示，也可以使用多磁盘相关的参数（master_data_volume_type、master_data_volume_size、master_data_volume_count、core_data_volume_type、core_data_volume_size和core_data_volume_count）表示。
func (c *MrsClient) CreateCluster(request *model.CreateClusterRequest) (*model.CreateClusterResponse, error) {
	requestDef := GenReqDefForCreateCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClusterResponse), nil
	}
}

//为特定的集群添加一个tag。 一个集群上最多有10个标签，此接口为幂等接口。添加标签时，如果创建的标签已经存在（key相同），则覆盖。
func (c *MrsClient) CreateClusterTag(request *model.CreateClusterTagRequest) (*model.CreateClusterTagResponse, error) {
	requestDef := GenReqDefForCreateClusterTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateClusterTagResponse), nil
	}
}

//数据完成处理分析后或者集群运行异常无法提供服务时可删除集群服务。该接口兼容Sahara。  处于如下状态的集群不允许删除： scaling-out：扩容中 scaling-in：缩容中 starting：启动中 terminating：删除中 terminated：已删除 failed：失败
func (c *MrsClient) DeleteCluster(request *model.DeleteClusterRequest) (*model.DeleteClusterResponse, error) {
	requestDef := GenReqDefForDeleteCluster()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteClusterResponse), nil
	}
}

//删除特定集群的标签。
func (c *MrsClient) DeleteClusterTag(request *model.DeleteClusterTagRequest) (*model.DeleteClusterTagResponse, error) {
	requestDef := GenReqDefForDeleteClusterTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteClusterTagResponse), nil
	}
}

//查询租户在指定Region下的所有标签集合 。
func (c *MrsClient) ListAllTags(request *model.ListAllTagsRequest) (*model.ListAllTagsResponse, error) {
	requestDef := GenReqDefForListAllTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAllTagsResponse), nil
	}
}

//查询指定集群的标签信息。
func (c *MrsClient) ListClusterTags(request *model.ListClusterTagsRequest) (*model.ListClusterTagsResponse, error) {
	requestDef := GenReqDefForListClusterTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListClusterTagsResponse), nil
	}
}
