package v1

import (
	http_client "code.byted.org/ti/huaweicloud-sdk-go-v3/core"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/services/hilens/v1/model"
)

type HiLensClient struct {
	HcClient *http_client.HcHttpClient
}

func NewHiLensClient(hcClient *http_client.HcHttpClient) *HiLensClient {
	return &HiLensClient{HcClient: hcClient}
}

func HiLensClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//获取设备告警列表
func (c *HiLensClient) ListDeviceAlarms(request *model.ListDeviceAlarmsRequest) (*model.ListDeviceAlarmsResponse, error) {
	requestDef := GenReqDefForListDeviceAlarms()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDeviceAlarmsResponse), nil
	}
}

//获取设备列表
func (c *HiLensClient) ListDevices(request *model.ListDevicesRequest) (*model.ListDevicesResponse, error) {
	requestDef := GenReqDefForListDevices()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDevicesResponse), nil
	}
}
