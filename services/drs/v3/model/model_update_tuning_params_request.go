package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateTuningParamsRequest Request Object
type UpdateTuningParamsRequest struct {

	// 租户在某一Region下的Job ID，如果是主备任务，使用父任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型
	XLanguage *UpdateTuningParamsRequestXLanguage `json:"X-Language,omitempty"`

	Body *ModifyTuningParamsReq `json:"body,omitempty"`
}

func (o UpdateTuningParamsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTuningParamsRequest struct{}"
	}

	return strings.Join([]string{"UpdateTuningParamsRequest", string(data)}, " ")
}

type UpdateTuningParamsRequestXLanguage struct {
	value string
}

type UpdateTuningParamsRequestXLanguageEnum struct {
	EN_US UpdateTuningParamsRequestXLanguage
	ZH_CN UpdateTuningParamsRequestXLanguage
}

func GetUpdateTuningParamsRequestXLanguageEnum() UpdateTuningParamsRequestXLanguageEnum {
	return UpdateTuningParamsRequestXLanguageEnum{
		EN_US: UpdateTuningParamsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: UpdateTuningParamsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c UpdateTuningParamsRequestXLanguage) Value() string {
	return c.value
}

func (c UpdateTuningParamsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTuningParamsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
