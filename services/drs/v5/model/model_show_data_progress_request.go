package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDataProgressRequest Request Object
type ShowDataProgressRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *ShowDataProgressRequestXLanguage `json:"X-Language,omitempty"`

	// 偏移量，表示查询该偏移量后面的记录。
	Offset *int32 `json:"offset,omitempty"`

	// 查询返回记录的数量限制。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowDataProgressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDataProgressRequest struct{}"
	}

	return strings.Join([]string{"ShowDataProgressRequest", string(data)}, " ")
}

type ShowDataProgressRequestXLanguage struct {
	value string
}

type ShowDataProgressRequestXLanguageEnum struct {
	EN_US ShowDataProgressRequestXLanguage
	ZH_CN ShowDataProgressRequestXLanguage
}

func GetShowDataProgressRequestXLanguageEnum() ShowDataProgressRequestXLanguageEnum {
	return ShowDataProgressRequestXLanguageEnum{
		EN_US: ShowDataProgressRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowDataProgressRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowDataProgressRequestXLanguage) Value() string {
	return c.value
}

func (c ShowDataProgressRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDataProgressRequestXLanguage) UnmarshalJSON(b []byte) error {
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
