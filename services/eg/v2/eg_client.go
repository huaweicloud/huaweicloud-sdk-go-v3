package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eg/v2/model"
)

type EgClient struct {
	HcClient *http_client.HcHttpClient
}

func NewEgClient(hcClient *http_client.HcHttpClient) *EgClient {
	return &EgClient{HcClient: hcClient}
}

func EgClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// PutEvents 用户自定义事件集成入口
//
// 用户自定义事件集成入口
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) PutEvents(request *model.PutEventsRequest) (*model.PutEventsResponse, error) {
	requestDef := GenReqDefForPutEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PutEventsResponse), nil
	}
}

// PutEventsInvoker 用户自定义事件集成入口
func (c *EgClient) PutEventsInvoker(request *model.PutEventsRequest) *PutEventsInvoker {
	requestDef := GenReqDefForPutEvents()
	return &PutEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// PutOfficialEvents 发布官方事件到事件通道
//
// 发布官方事件到事件通道。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EgClient) PutOfficialEvents(request *model.PutOfficialEventsRequest) (*model.PutOfficialEventsResponse, error) {
	requestDef := GenReqDefForPutOfficialEvents()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.PutOfficialEventsResponse), nil
	}
}

// PutOfficialEventsInvoker 发布官方事件到事件通道
func (c *EgClient) PutOfficialEventsInvoker(request *model.PutOfficialEventsRequest) *PutOfficialEventsInvoker {
	requestDef := GenReqDefForPutOfficialEvents()
	return &PutOfficialEventsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
