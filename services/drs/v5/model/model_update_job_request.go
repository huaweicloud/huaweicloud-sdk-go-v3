package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type UpdateJobRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *UpdateJobRequestXLanguage `json:"X-Language,omitempty"`

	Body *SingleUpdateJobReq `json:"body,omitempty"`
}

func (o UpdateJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateJobRequest struct{}"
	}

	return strings.Join([]string{"UpdateJobRequest", string(data)}, " ")
}

type UpdateJobRequestXLanguage struct {
	value string
}

type UpdateJobRequestXLanguageEnum struct {
	EN_US UpdateJobRequestXLanguage
	ZH_CN UpdateJobRequestXLanguage
}

func GetUpdateJobRequestXLanguageEnum() UpdateJobRequestXLanguageEnum {
	return UpdateJobRequestXLanguageEnum{
		EN_US: UpdateJobRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: UpdateJobRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c UpdateJobRequestXLanguage) Value() string {
	return c.value
}

func (c UpdateJobRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateJobRequestXLanguage) UnmarshalJSON(b []byte) error {
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
