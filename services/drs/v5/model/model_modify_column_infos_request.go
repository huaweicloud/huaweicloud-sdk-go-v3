package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ModifyColumnInfosRequest Request Object
type ModifyColumnInfosRequest struct {

	// 请求语言类型。
	XLanguage *ModifyColumnInfosRequestXLanguage `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`

	Body *ModifyColumnInfoReq `json:"body,omitempty"`
}

func (o ModifyColumnInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyColumnInfosRequest struct{}"
	}

	return strings.Join([]string{"ModifyColumnInfosRequest", string(data)}, " ")
}

type ModifyColumnInfosRequestXLanguage struct {
	value string
}

type ModifyColumnInfosRequestXLanguageEnum struct {
	EN_US ModifyColumnInfosRequestXLanguage
	ZH_CN ModifyColumnInfosRequestXLanguage
}

func GetModifyColumnInfosRequestXLanguageEnum() ModifyColumnInfosRequestXLanguageEnum {
	return ModifyColumnInfosRequestXLanguageEnum{
		EN_US: ModifyColumnInfosRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ModifyColumnInfosRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ModifyColumnInfosRequestXLanguage) Value() string {
	return c.value
}

func (c ModifyColumnInfosRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyColumnInfosRequestXLanguage) UnmarshalJSON(b []byte) error {
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
