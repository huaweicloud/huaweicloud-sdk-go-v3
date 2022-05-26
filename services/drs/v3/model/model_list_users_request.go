package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListUsersRequest struct {

	// 请求语言类型
	XLanguage *ListUsersRequestXLanguage `json:"X-Language,omitempty"`

	// 任务ID
	JobId string `json:"job_id"`
}

func (o ListUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsersRequest struct{}"
	}

	return strings.Join([]string{"ListUsersRequest", string(data)}, " ")
}

type ListUsersRequestXLanguage struct {
	value string
}

type ListUsersRequestXLanguageEnum struct {
	EN_US ListUsersRequestXLanguage
	ZH_CN ListUsersRequestXLanguage
}

func GetListUsersRequestXLanguageEnum() ListUsersRequestXLanguageEnum {
	return ListUsersRequestXLanguageEnum{
		EN_US: ListUsersRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListUsersRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListUsersRequestXLanguage) Value() string {
	return c.value
}

func (c ListUsersRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListUsersRequestXLanguage) UnmarshalJSON(b []byte) error {
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
