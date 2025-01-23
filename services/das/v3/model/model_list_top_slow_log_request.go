package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTopSlowLogRequest Request Object
type ListTopSlowLogRequest struct {

	// TOP数量
	Num int32 `json:"num"`

	// 语言
	XLanguage *ListTopSlowLogRequestXLanguage `json:"X-Language,omitempty"`

	// 开始时间
	StartAt int64 `json:"start_at"`

	// 结束时间
	EndAt int64 `json:"end_at"`
}

func (o ListTopSlowLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopSlowLogRequest struct{}"
	}

	return strings.Join([]string{"ListTopSlowLogRequest", string(data)}, " ")
}

type ListTopSlowLogRequestXLanguage struct {
	value string
}

type ListTopSlowLogRequestXLanguageEnum struct {
	ZH_CN ListTopSlowLogRequestXLanguage
	EN_US ListTopSlowLogRequestXLanguage
}

func GetListTopSlowLogRequestXLanguageEnum() ListTopSlowLogRequestXLanguageEnum {
	return ListTopSlowLogRequestXLanguageEnum{
		ZH_CN: ListTopSlowLogRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListTopSlowLogRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListTopSlowLogRequestXLanguage) Value() string {
	return c.value
}

func (c ListTopSlowLogRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTopSlowLogRequestXLanguage) UnmarshalJSON(b []byte) error {
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
