package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowLinkFlavorsRequest Request Object
type ShowLinkFlavorsRequest struct {

	// 请求语言类型。
	XLanguage *ShowLinkFlavorsRequestXLanguage `json:"X-Language,omitempty"`

	// 引擎类型。
	EngineType *string `json:"engine_type,omitempty"`

	// 任务场景。取值：     - migration：实时迁移。     - sync：实时同步。     - cloudDataGuard：实时灾备。     - subscription：数据订阅。     - replay：录制回放。     - verify：校验任务。     - cdc：CDC任务。
	JobType *string `json:"job_type,omitempty"`

	// 任务方向。取值：     - up：入云 ，灾备场景时对应本云为备。     - down：出云，灾备场景时对应本云为主。     - non-dbs：自建。
	JobDirection *string `json:"job_direction,omitempty"`

	// 网络类型。取值：     - eip：公网网络。     - vpc：VPC网络，灾备场景不支持选择VPC网络。     - vpn：VPN、专线网络。
	NetType *string `json:"net_type,omitempty"`

	// 任务实例规格类型。取值：     - micro：极小规格。     - small：小规格。     - medium：中规格。     - high：大规格。     - xlarge：超大规格。     - 2xlarge：极大规格。
	NodeType *string `json:"node_type,omitempty"`

	// 灾备类型是否双主灾备。说明：     - job_type 是cloudDataGuard时，必填，灾备类型是双主灾备时，is_multi_write取值true, 否则为false。     - job_type 是其他类型时，is_multi_write是非必选参数。
	IsMultiWrite *bool `json:"is_multi_write,omitempty"`
}

func (o ShowLinkFlavorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLinkFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ShowLinkFlavorsRequest", string(data)}, " ")
}

type ShowLinkFlavorsRequestXLanguage struct {
	value string
}

type ShowLinkFlavorsRequestXLanguageEnum struct {
	EN_US ShowLinkFlavorsRequestXLanguage
	ZH_CN ShowLinkFlavorsRequestXLanguage
}

func GetShowLinkFlavorsRequestXLanguageEnum() ShowLinkFlavorsRequestXLanguageEnum {
	return ShowLinkFlavorsRequestXLanguageEnum{
		EN_US: ShowLinkFlavorsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowLinkFlavorsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowLinkFlavorsRequestXLanguage) Value() string {
	return c.value
}

func (c ShowLinkFlavorsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowLinkFlavorsRequestXLanguage) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
