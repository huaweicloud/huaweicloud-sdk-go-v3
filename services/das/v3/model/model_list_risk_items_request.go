package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListRiskItemsRequest Request Object
type ListRiskItemsRequest struct {

	// 语言
	XLanguage *ListRiskItemsRequestXLanguage `json:"X-Language,omitempty"`

	// 数据库类型
	DatastoreType string `json:"datastore_type"`
}

func (o ListRiskItemsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRiskItemsRequest struct{}"
	}

	return strings.Join([]string{"ListRiskItemsRequest", string(data)}, " ")
}

type ListRiskItemsRequestXLanguage struct {
	value string
}

type ListRiskItemsRequestXLanguageEnum struct {
	ZH_CN ListRiskItemsRequestXLanguage
	EN_US ListRiskItemsRequestXLanguage
}

func GetListRiskItemsRequestXLanguageEnum() ListRiskItemsRequestXLanguageEnum {
	return ListRiskItemsRequestXLanguageEnum{
		ZH_CN: ListRiskItemsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListRiskItemsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListRiskItemsRequestXLanguage) Value() string {
	return c.value
}

func (c ListRiskItemsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRiskItemsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
