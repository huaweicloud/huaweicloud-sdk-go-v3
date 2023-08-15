package v2

import (
	http_client "github.com/dysodeng/huaweicloud-sdk-go-v3/core"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/services/eihealth/v2/model"
)

type EiHealthClient struct {
	HcClient *http_client.HcHttpClient
}

func NewEiHealthClient(hcClient *http_client.HcHttpClient) *EiHealthClient {
	return &EiHealthClient{HcClient: hcClient}
}

func EiHealthClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// ShowAdmetWithCustomProps ADMET属性预测接口(默认+自定义属性)
//
// 计算小分子的物化性质，包括默认的吸收(adsorption)、分布(distribution)、代谢(metabolism)、清除(excretion)与毒性(toxicity)，以及用户自定义的属性。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *EiHealthClient) ShowAdmetWithCustomProps(request *model.ShowAdmetWithCustomPropsRequest) (*model.ShowAdmetWithCustomPropsResponse, error) {
	requestDef := GenReqDefForShowAdmetWithCustomProps()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAdmetWithCustomPropsResponse), nil
	}
}

// ShowAdmetWithCustomPropsInvoker ADMET属性预测接口(默认+自定义属性)
func (c *EiHealthClient) ShowAdmetWithCustomPropsInvoker(request *model.ShowAdmetWithCustomPropsRequest) *ShowAdmetWithCustomPropsInvoker {
	requestDef := GenReqDefForShowAdmetWithCustomProps()
	return &ShowAdmetWithCustomPropsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
