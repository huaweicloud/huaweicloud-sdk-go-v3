package v1

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/mas/v1/model"
)

type MasClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewMasClient(hcClient *httpclient.HcHttpClient) *MasClient {
	return &MasClient{HcClient: hcClient}
}

func MasClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// ShowNameSpaceList 查询命名空间列表
//
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *MasClient) ShowNameSpaceList(request *model.ShowNameSpaceListRequest) (*model.ShowNameSpaceListResponse, error) {
	requestDef := GenReqDefForShowNameSpaceList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNameSpaceListResponse), nil
	}
}

// ShowNameSpaceListInvoker 查询命名空间列表
func (c *MasClient) ShowNameSpaceListInvoker(request *model.ShowNameSpaceListRequest) *ShowNameSpaceListInvoker {
	requestDef := GenReqDefForShowNameSpaceList()
	return &ShowNameSpaceListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
