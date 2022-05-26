package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type UpdateApplicationRequest struct {

	// 语言类型 中文:zh-cn 英文:en-us
	XLanguage *UpdateApplicationRequestXLanguage `json:"X-Language,omitempty"`

	// 应用id
	ApplicationId string `json:"application_id"`

	Body *ApplicationModifyInfo `json:"body,omitempty"`
}

func (o UpdateApplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApplicationRequest struct{}"
	}

	return strings.Join([]string{"UpdateApplicationRequest", string(data)}, " ")
}

type UpdateApplicationRequestXLanguage struct {
	value string
}

type UpdateApplicationRequestXLanguageEnum struct {
	ZH_CN UpdateApplicationRequestXLanguage
	EN_US UpdateApplicationRequestXLanguage
}

func GetUpdateApplicationRequestXLanguageEnum() UpdateApplicationRequestXLanguageEnum {
	return UpdateApplicationRequestXLanguageEnum{
		ZH_CN: UpdateApplicationRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: UpdateApplicationRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c UpdateApplicationRequestXLanguage) Value() string {
	return c.value
}

func (c UpdateApplicationRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateApplicationRequestXLanguage) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
