package v5

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cgs/v5/model"
)

type CgsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCgsClient(hcClient *http_client.HcHttpClient) *CgsClient {
	return &CgsClient{HcClient: hcClient}
}

func CgsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// ListContainerNodes 查询容器节点列表
//
// 查询容器节点列表（仅新版本容器安全支持，即将上线，敬请期待！）
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CgsClient) ListContainerNodes(request *model.ListContainerNodesRequest) (*model.ListContainerNodesResponse, error) {
	requestDef := GenReqDefForListContainerNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListContainerNodesResponse), nil
	}
}

// ListContainerNodesInvoker 查询容器节点列表
func (c *CgsClient) ListContainerNodesInvoker(request *model.ListContainerNodesRequest) *ListContainerNodesInvoker {
	requestDef := GenReqDefForListContainerNodes()
	return &ListContainerNodesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
