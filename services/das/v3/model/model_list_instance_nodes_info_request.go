package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListInstanceNodesInfoRequest Request Object
type ListInstanceNodesInfoRequest struct {

	// 语言
	XLanguage *ListInstanceNodesInfoRequestXLanguage `json:"X-Language,omitempty"`

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 数据库类型
	DatastoreType string `json:"datastore_type"`
}

func (o ListInstanceNodesInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceNodesInfoRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceNodesInfoRequest", string(data)}, " ")
}

type ListInstanceNodesInfoRequestXLanguage struct {
	value string
}

type ListInstanceNodesInfoRequestXLanguageEnum struct {
	ZH_CN ListInstanceNodesInfoRequestXLanguage
	EN_US ListInstanceNodesInfoRequestXLanguage
}

func GetListInstanceNodesInfoRequestXLanguageEnum() ListInstanceNodesInfoRequestXLanguageEnum {
	return ListInstanceNodesInfoRequestXLanguageEnum{
		ZH_CN: ListInstanceNodesInfoRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListInstanceNodesInfoRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListInstanceNodesInfoRequestXLanguage) Value() string {
	return c.value
}

func (c ListInstanceNodesInfoRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstanceNodesInfoRequestXLanguage) UnmarshalJSON(b []byte) error {
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
