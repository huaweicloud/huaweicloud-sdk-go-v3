package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/identitycenter/v1/model"
)

type IdentityCenterClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewIdentityCenterClient(hcClient *httpclient.HcHttpClient) *IdentityCenterClient {
	return &IdentityCenterClient{HcClient: hcClient}
}

func IdentityCenterClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder().WithCredentialsType("global.Credentials")
	return builder
}

// ListInstances 列出实例
//
// 查询IAM身份中心的实例列表。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *IdentityCenterClient) ListInstances(request *model.ListInstancesRequest) (*model.ListInstancesResponse, error) {
	requestDef := GenReqDefForListInstances()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListInstancesResponse), nil
	}
}

// ListInstancesInvoker 列出实例
func (c *IdentityCenterClient) ListInstancesInvoker(request *model.ListInstancesRequest) *ListInstancesInvoker {
	requestDef := GenReqDefForListInstances()
	return &ListInstancesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
