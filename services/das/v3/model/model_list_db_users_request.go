package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDbUsersRequest Request Object
type ListDbUsersRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为100，不能为负数，最小值为1，最大值为100。
	Limit *int32 `json:"limit,omitempty"`

	// 数据库用户ID。用户使用数据库账号与数据库建立的连接ID（数据库用户ID由注册数据库用户接口创建）。
	DbUserId *string `json:"db_user_id,omitempty"`

	// 数据库用户名称
	DbUsername *string `json:"db_username,omitempty"`

	// 语言
	XLanguage *ListDbUsersRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListDbUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDbUsersRequest struct{}"
	}

	return strings.Join([]string{"ListDbUsersRequest", string(data)}, " ")
}

type ListDbUsersRequestXLanguage struct {
	value string
}

type ListDbUsersRequestXLanguageEnum struct {
	ZH_CN ListDbUsersRequestXLanguage
	EN_US ListDbUsersRequestXLanguage
}

func GetListDbUsersRequestXLanguageEnum() ListDbUsersRequestXLanguageEnum {
	return ListDbUsersRequestXLanguageEnum{
		ZH_CN: ListDbUsersRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListDbUsersRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListDbUsersRequestXLanguage) Value() string {
	return c.value
}

func (c ListDbUsersRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDbUsersRequestXLanguage) UnmarshalJSON(b []byte) error {
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
