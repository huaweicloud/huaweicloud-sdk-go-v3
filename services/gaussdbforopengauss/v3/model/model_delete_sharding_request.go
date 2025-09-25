package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteShardingRequest Request Object
type DeleteShardingRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *DeleteShardingRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为32个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	Body *ReduceDnRequestBody `json:"body,omitempty"`
}

func (o DeleteShardingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteShardingRequest struct{}"
	}

	return strings.Join([]string{"DeleteShardingRequest", string(data)}, " ")
}

type DeleteShardingRequestXLanguage struct {
	value string
}

type DeleteShardingRequestXLanguageEnum struct {
	ZH_CN DeleteShardingRequestXLanguage
	EN_US DeleteShardingRequestXLanguage
}

func GetDeleteShardingRequestXLanguageEnum() DeleteShardingRequestXLanguageEnum {
	return DeleteShardingRequestXLanguageEnum{
		ZH_CN: DeleteShardingRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: DeleteShardingRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c DeleteShardingRequestXLanguage) Value() string {
	return c.value
}

func (c DeleteShardingRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteShardingRequestXLanguage) UnmarshalJSON(b []byte) error {
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
