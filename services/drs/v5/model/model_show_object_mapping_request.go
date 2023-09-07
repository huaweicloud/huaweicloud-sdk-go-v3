package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowObjectMappingRequest Request Object
type ShowObjectMappingRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *ShowObjectMappingRequestXLanguage `json:"X-Language,omitempty"`

	Body *QueryUserSelectedObjectInfoReq `json:"body,omitempty"`
}

func (o ShowObjectMappingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowObjectMappingRequest struct{}"
	}

	return strings.Join([]string{"ShowObjectMappingRequest", string(data)}, " ")
}

type ShowObjectMappingRequestXLanguage struct {
	value string
}

type ShowObjectMappingRequestXLanguageEnum struct {
	EN_US ShowObjectMappingRequestXLanguage
	ZH_CN ShowObjectMappingRequestXLanguage
}

func GetShowObjectMappingRequestXLanguageEnum() ShowObjectMappingRequestXLanguageEnum {
	return ShowObjectMappingRequestXLanguageEnum{
		EN_US: ShowObjectMappingRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowObjectMappingRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowObjectMappingRequestXLanguage) Value() string {
	return c.value
}

func (c ShowObjectMappingRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowObjectMappingRequestXLanguage) UnmarshalJSON(b []byte) error {
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
