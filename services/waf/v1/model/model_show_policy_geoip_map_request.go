package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowPolicyGeoipMapRequest Request Object
type ShowPolicyGeoipMapRequest struct {

	// **参数解释：** 语言的类型 - cn代表中文 - en代表英文  **约束限制：** 不涉及 **取值范围：** - cn - en  **默认取值：** - cn
	Lang *ShowPolicyGeoipMapRequestLang `json:"lang,omitempty"`
}

func (o ShowPolicyGeoipMapRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPolicyGeoipMapRequest struct{}"
	}

	return strings.Join([]string{"ShowPolicyGeoipMapRequest", string(data)}, " ")
}

type ShowPolicyGeoipMapRequestLang struct {
	value string
}

type ShowPolicyGeoipMapRequestLangEnum struct {
	CN ShowPolicyGeoipMapRequestLang
	EN ShowPolicyGeoipMapRequestLang
}

func GetShowPolicyGeoipMapRequestLangEnum() ShowPolicyGeoipMapRequestLangEnum {
	return ShowPolicyGeoipMapRequestLangEnum{
		CN: ShowPolicyGeoipMapRequestLang{
			value: "cn",
		},
		EN: ShowPolicyGeoipMapRequestLang{
			value: "en",
		},
	}
}

func (c ShowPolicyGeoipMapRequestLang) Value() string {
	return c.value
}

func (c ShowPolicyGeoipMapRequestLang) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowPolicyGeoipMapRequestLang) UnmarshalJSON(b []byte) error {
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
