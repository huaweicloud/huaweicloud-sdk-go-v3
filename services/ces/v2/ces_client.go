package v2

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ces/v2/model"
)

type CesClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCesClient(hcClient *http_client.HcHttpClient) *CesClient {
	return &CesClient{HcClient: hcClient}
}

func CesClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//查询告警历史列表
func (c *CesClient) ListAlarmHistories(request *model.ListAlarmHistoriesRequest) (*model.ListAlarmHistoriesResponse, error) {
	requestDef := GenReqDefForListAlarmHistories()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAlarmHistoriesResponse), nil
	}
}
