package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListInnodbLocksRequest Request Object
type ListInnodbLocksRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 数据库用户ID
	DbUserId string `json:"db_user_id"`

	// 语言
	XLanguage *ListInnodbLocksRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListInnodbLocksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInnodbLocksRequest struct{}"
	}

	return strings.Join([]string{"ListInnodbLocksRequest", string(data)}, " ")
}

type ListInnodbLocksRequestXLanguage struct {
	value string
}

type ListInnodbLocksRequestXLanguageEnum struct {
	ZH_CN ListInnodbLocksRequestXLanguage
	EN_US ListInnodbLocksRequestXLanguage
}

func GetListInnodbLocksRequestXLanguageEnum() ListInnodbLocksRequestXLanguageEnum {
	return ListInnodbLocksRequestXLanguageEnum{
		ZH_CN: ListInnodbLocksRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListInnodbLocksRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListInnodbLocksRequestXLanguage) Value() string {
	return c.value
}

func (c ListInnodbLocksRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInnodbLocksRequestXLanguage) UnmarshalJSON(b []byte) error {
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
