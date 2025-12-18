package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTopTableVolumesRequest Request Object
type ListTopTableVolumesRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ListTopTableVolumesRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**: 工作流ID。 **约束限制**: 不涉及 **取值范围**: 不涉及 **默认取值**: 不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**: 工作流执行节点ID。 **约束限制**: 不涉及 **取值范围**: 不涉及 **默认取值**: 不涉及。
	NodeId *string `json:"node_id,omitempty"`
}

func (o ListTopTableVolumesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopTableVolumesRequest struct{}"
	}

	return strings.Join([]string{"ListTopTableVolumesRequest", string(data)}, " ")
}

type ListTopTableVolumesRequestXLanguage struct {
	value string
}

type ListTopTableVolumesRequestXLanguageEnum struct {
	ZH_CN ListTopTableVolumesRequestXLanguage
	EN_US ListTopTableVolumesRequestXLanguage
}

func GetListTopTableVolumesRequestXLanguageEnum() ListTopTableVolumesRequestXLanguageEnum {
	return ListTopTableVolumesRequestXLanguageEnum{
		ZH_CN: ListTopTableVolumesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListTopTableVolumesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListTopTableVolumesRequestXLanguage) Value() string {
	return c.value
}

func (c ListTopTableVolumesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTopTableVolumesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
