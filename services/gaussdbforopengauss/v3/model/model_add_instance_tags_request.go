package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AddInstanceTagsRequest Request Object
type AddInstanceTagsRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**:   - zh-cn   - en-us  **默认取值**: en-us
	XLanguage *AddInstanceTagsRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	Body *AddTagsRequestBody `json:"body,omitempty"`
}

func (o AddInstanceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddInstanceTagsRequest struct{}"
	}

	return strings.Join([]string{"AddInstanceTagsRequest", string(data)}, " ")
}

type AddInstanceTagsRequestXLanguage struct {
	value string
}

type AddInstanceTagsRequestXLanguageEnum struct {
	ZH_CN AddInstanceTagsRequestXLanguage
	EN_US AddInstanceTagsRequestXLanguage
}

func GetAddInstanceTagsRequestXLanguageEnum() AddInstanceTagsRequestXLanguageEnum {
	return AddInstanceTagsRequestXLanguageEnum{
		ZH_CN: AddInstanceTagsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: AddInstanceTagsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c AddInstanceTagsRequestXLanguage) Value() string {
	return c.value
}

func (c AddInstanceTagsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddInstanceTagsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
