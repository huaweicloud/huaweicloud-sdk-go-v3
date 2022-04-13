package v5

import (
	http_client "code.byted.org/ti/huaweicloud-sdk-go-v3/core"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/services/cgs/v5/model"
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

//查询容器节点列表（仅新版本容器安全支持，即将上线，敬请期待！）
func (c *CgsClient) ListContainerNodes(request *model.ListContainerNodesRequest) (*model.ListContainerNodesResponse, error) {
	requestDef := GenReqDefForListContainerNodes()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListContainerNodesResponse), nil
	}
}
