package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListInstanceTopSlowLogRequest Request Object
type ListInstanceTopSlowLogRequest struct {

	// TOP数量
	Num int32 `json:"num"`

	// 语言
	XLanguage *ListInstanceTopSlowLogRequestXLanguage `json:"X-Language,omitempty"`

	// 开始时间
	StartAt int64 `json:"start_at"`

	// 结束时间
	EndAt int64 `json:"end_at"`

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o ListInstanceTopSlowLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceTopSlowLogRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceTopSlowLogRequest", string(data)}, " ")
}

type ListInstanceTopSlowLogRequestXLanguage struct {
	value string
}

type ListInstanceTopSlowLogRequestXLanguageEnum struct {
	ZH_CN ListInstanceTopSlowLogRequestXLanguage
	EN_US ListInstanceTopSlowLogRequestXLanguage
}

func GetListInstanceTopSlowLogRequestXLanguageEnum() ListInstanceTopSlowLogRequestXLanguageEnum {
	return ListInstanceTopSlowLogRequestXLanguageEnum{
		ZH_CN: ListInstanceTopSlowLogRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListInstanceTopSlowLogRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListInstanceTopSlowLogRequestXLanguage) Value() string {
	return c.value
}

func (c ListInstanceTopSlowLogRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstanceTopSlowLogRequestXLanguage) UnmarshalJSON(b []byte) error {
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
