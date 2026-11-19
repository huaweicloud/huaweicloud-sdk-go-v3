package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSessionSummaryUseOpsTunnelRequest Request Object
type ShowSessionSummaryUseOpsTunnelRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn - en-us  **默认取值**: en-us
	XLanguage *ShowSessionSummaryUseOpsTunnelRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**: 节点ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	NodeId *string `json:"node_id,omitempty"`

	// **参数解释**: 组件ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	ComponentId *string `json:"component_id,omitempty"`

	// **参数解释**: 是否走紧急通道。 **约束限制**: 不涉及。 **取值范围**: - true：走紧急通道 - false：不走紧急通道  **默认取值**: false。
	UseOpsTunnel bool `json:"use_ops_tunnel"`
}

func (o ShowSessionSummaryUseOpsTunnelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSessionSummaryUseOpsTunnelRequest struct{}"
	}

	return strings.Join([]string{"ShowSessionSummaryUseOpsTunnelRequest", string(data)}, " ")
}

type ShowSessionSummaryUseOpsTunnelRequestXLanguage struct {
	value string
}

type ShowSessionSummaryUseOpsTunnelRequestXLanguageEnum struct {
	ZH_CN ShowSessionSummaryUseOpsTunnelRequestXLanguage
	EN_US ShowSessionSummaryUseOpsTunnelRequestXLanguage
}

func GetShowSessionSummaryUseOpsTunnelRequestXLanguageEnum() ShowSessionSummaryUseOpsTunnelRequestXLanguageEnum {
	return ShowSessionSummaryUseOpsTunnelRequestXLanguageEnum{
		ZH_CN: ShowSessionSummaryUseOpsTunnelRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowSessionSummaryUseOpsTunnelRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowSessionSummaryUseOpsTunnelRequestXLanguage) Value() string {
	return c.value
}

func (c ShowSessionSummaryUseOpsTunnelRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSessionSummaryUseOpsTunnelRequestXLanguage) UnmarshalJSON(b []byte) error {
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
