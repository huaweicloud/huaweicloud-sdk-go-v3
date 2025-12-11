package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dcc/v1/model"
)

type DccClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewDccClient(hcClient *httpclient.HcHttpClient) *DccClient {
	return &DccClient{HcClient: hcClient}
}

func DccClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// ShowResourceClusters 获取专属计算集群资源
//
// 获取专属计算集群资源
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *DccClient) ShowResourceClusters(request *model.ShowResourceClustersRequest) (*model.ShowResourceClustersResponse, error) {
	requestDef := GenReqDefForShowResourceClusters()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceClustersResponse), nil
	}
}

// ShowResourceClustersInvoker 获取专属计算集群资源
func (c *DccClient) ShowResourceClustersInvoker(request *model.ShowResourceClustersRequest) *ShowResourceClustersInvoker {
	requestDef := GenReqDefForShowResourceClusters()
	return &ShowResourceClustersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
