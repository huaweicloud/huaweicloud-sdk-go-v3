package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowInstanceMetricDataRequest Request Object
type ShowInstanceMetricDataRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ShowInstanceMetricDataRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**: 开始时间，时间戳格式，例如：1756971683303。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	StartTime string `json:"start_time"`

	// **参数解释**: 结束时间，时间戳格式，例如：1756975283303。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	EndTime string `json:"end_time"`

	// **参数解释**: 指标ID，可通过“查询指标分组的指标名称”接口获取，例如查询CPU使用率，传值 rds001_cpu_util。 **约束限制**: 不涉及。
	Metric string `json:"metric"`

	// **参数解释**: 节点ID。 **约束限制**: 不涉及。
	NodeId string `json:"node_id"`

	// **参数解释**: 组件ID，例如dn_6001。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	ComponentId *string `json:"component_id,omitempty"`
}

func (o ShowInstanceMetricDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceMetricDataRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceMetricDataRequest", string(data)}, " ")
}

type ShowInstanceMetricDataRequestXLanguage struct {
	value string
}

type ShowInstanceMetricDataRequestXLanguageEnum struct {
	ZH_CN ShowInstanceMetricDataRequestXLanguage
	EN_US ShowInstanceMetricDataRequestXLanguage
}

func GetShowInstanceMetricDataRequestXLanguageEnum() ShowInstanceMetricDataRequestXLanguageEnum {
	return ShowInstanceMetricDataRequestXLanguageEnum{
		ZH_CN: ShowInstanceMetricDataRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowInstanceMetricDataRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowInstanceMetricDataRequestXLanguage) Value() string {
	return c.value
}

func (c ShowInstanceMetricDataRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceMetricDataRequestXLanguage) UnmarshalJSON(b []byte) error {
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
