package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type DeleteJobRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *DeleteJobRequestXLanguage `json:"X-Language,omitempty"`
}

func (o DeleteJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteJobRequest struct{}"
	}

	return strings.Join([]string{"DeleteJobRequest", string(data)}, " ")
}

type DeleteJobRequestXLanguage struct {
	value string
}

type DeleteJobRequestXLanguageEnum struct {
	EN_US DeleteJobRequestXLanguage
	ZH_CN DeleteJobRequestXLanguage
}

func GetDeleteJobRequestXLanguageEnum() DeleteJobRequestXLanguageEnum {
	return DeleteJobRequestXLanguageEnum{
		EN_US: DeleteJobRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: DeleteJobRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c DeleteJobRequestXLanguage) Value() string {
	return c.value
}

func (c DeleteJobRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteJobRequestXLanguage) UnmarshalJSON(b []byte) error {
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
