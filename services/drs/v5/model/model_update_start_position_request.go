package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateStartPositionRequest Request Object
type UpdateStartPositionRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *UpdateStartPositionRequestXLanguage `json:"X-Language,omitempty"`

	Body *ModifyStartPositionReq `json:"body,omitempty"`
}

func (o UpdateStartPositionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStartPositionRequest struct{}"
	}

	return strings.Join([]string{"UpdateStartPositionRequest", string(data)}, " ")
}

type UpdateStartPositionRequestXLanguage struct {
	value string
}

type UpdateStartPositionRequestXLanguageEnum struct {
	EN_US UpdateStartPositionRequestXLanguage
	ZH_CN UpdateStartPositionRequestXLanguage
}

func GetUpdateStartPositionRequestXLanguageEnum() UpdateStartPositionRequestXLanguageEnum {
	return UpdateStartPositionRequestXLanguageEnum{
		EN_US: UpdateStartPositionRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: UpdateStartPositionRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c UpdateStartPositionRequestXLanguage) Value() string {
	return c.value
}

func (c UpdateStartPositionRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateStartPositionRequestXLanguage) UnmarshalJSON(b []byte) error {
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
