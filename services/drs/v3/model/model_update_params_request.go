package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type UpdateParamsRequest struct {

	// 任务ID
	JobId string `json:"job_id" xml:"job_id"`

	// 请求语言类型
	XLanguage *UpdateParamsRequestXLanguage `json:"X-Language,omitempty" xml:"X-Language"`

	Body *ModifyTargetParamsReq `json:"body,omitempty" xml:"body"`
}

func (o UpdateParamsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateParamsRequest struct{}"
	}

	return strings.Join([]string{"UpdateParamsRequest", string(data)}, " ")
}

type UpdateParamsRequestXLanguage struct {
	value string
}

type UpdateParamsRequestXLanguageEnum struct {
	EN_US UpdateParamsRequestXLanguage
	ZH_CN UpdateParamsRequestXLanguage
}

func GetUpdateParamsRequestXLanguageEnum() UpdateParamsRequestXLanguageEnum {
	return UpdateParamsRequestXLanguageEnum{
		EN_US: UpdateParamsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: UpdateParamsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c UpdateParamsRequestXLanguage) Value() string {
	return c.value
}

func (c UpdateParamsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateParamsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
