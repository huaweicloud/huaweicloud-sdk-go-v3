package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListProcessesRequest Request Object
type ListProcessesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 数据库用户ID。用户使用数据库账号与数据库建立的连接ID（数据库用户ID由注册数据库用户接口创建）。
	DbUserId string `json:"db_user_id"`

	// 用户
	User *string `json:"user,omitempty"`

	// 数据库
	Database *string `json:"database,omitempty"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 每页记录数，默认为20，最大取值100。
	Limit *int32 `json:"limit,omitempty"`

	// 语言
	XLanguage *ListProcessesRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListProcessesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProcessesRequest struct{}"
	}

	return strings.Join([]string{"ListProcessesRequest", string(data)}, " ")
}

type ListProcessesRequestXLanguage struct {
	value string
}

type ListProcessesRequestXLanguageEnum struct {
	ZH_CN ListProcessesRequestXLanguage
	EN_US ListProcessesRequestXLanguage
}

func GetListProcessesRequestXLanguageEnum() ListProcessesRequestXLanguageEnum {
	return ListProcessesRequestXLanguageEnum{
		ZH_CN: ListProcessesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListProcessesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListProcessesRequestXLanguage) Value() string {
	return c.value
}

func (c ListProcessesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListProcessesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
