package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListMetadataLocksRequest Request Object
type ListMetadataLocksRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 数据库用户ID
	DbUserId string `json:"db_user_id"`

	// 会话ID
	ThreadId *string `json:"thread_id,omitempty"`

	// 数据库名称
	Database *string `json:"database,omitempty"`

	// 表名
	Table *string `json:"table,omitempty"`

	// 语言
	XLanguage *ListMetadataLocksRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListMetadataLocksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetadataLocksRequest struct{}"
	}

	return strings.Join([]string{"ListMetadataLocksRequest", string(data)}, " ")
}

type ListMetadataLocksRequestXLanguage struct {
	value string
}

type ListMetadataLocksRequestXLanguageEnum struct {
	ZH_CN ListMetadataLocksRequestXLanguage
	EN_US ListMetadataLocksRequestXLanguage
}

func GetListMetadataLocksRequestXLanguageEnum() ListMetadataLocksRequestXLanguageEnum {
	return ListMetadataLocksRequestXLanguageEnum{
		ZH_CN: ListMetadataLocksRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListMetadataLocksRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListMetadataLocksRequestXLanguage) Value() string {
	return c.value
}

func (c ListMetadataLocksRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMetadataLocksRequestXLanguage) UnmarshalJSON(b []byte) error {
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
