package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ModifyGroupAndStreamRequest Request Object
type ModifyGroupAndStreamRequest struct {

	// 请求语言类型。
	XLanguage *ModifyGroupAndStreamRequestXLanguage `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`

	Body *LtsInfo `json:"body,omitempty"`
}

func (o ModifyGroupAndStreamRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyGroupAndStreamRequest struct{}"
	}

	return strings.Join([]string{"ModifyGroupAndStreamRequest", string(data)}, " ")
}

type ModifyGroupAndStreamRequestXLanguage struct {
	value string
}

type ModifyGroupAndStreamRequestXLanguageEnum struct {
	EN_US ModifyGroupAndStreamRequestXLanguage
	ZH_CN ModifyGroupAndStreamRequestXLanguage
}

func GetModifyGroupAndStreamRequestXLanguageEnum() ModifyGroupAndStreamRequestXLanguageEnum {
	return ModifyGroupAndStreamRequestXLanguageEnum{
		EN_US: ModifyGroupAndStreamRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ModifyGroupAndStreamRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ModifyGroupAndStreamRequestXLanguage) Value() string {
	return c.value
}

func (c ModifyGroupAndStreamRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyGroupAndStreamRequestXLanguage) UnmarshalJSON(b []byte) error {
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
