package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListInstanceDistributionRequest Request Object
type ListInstanceDistributionRequest struct {

	// 数据库类型
	DatastoreType *string `json:"datastore_type,omitempty"`

	// 语言
	XLanguage *ListInstanceDistributionRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListInstanceDistributionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceDistributionRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceDistributionRequest", string(data)}, " ")
}

type ListInstanceDistributionRequestXLanguage struct {
	value string
}

type ListInstanceDistributionRequestXLanguageEnum struct {
	ZH_CN ListInstanceDistributionRequestXLanguage
	EN_US ListInstanceDistributionRequestXLanguage
}

func GetListInstanceDistributionRequestXLanguageEnum() ListInstanceDistributionRequestXLanguageEnum {
	return ListInstanceDistributionRequestXLanguageEnum{
		ZH_CN: ListInstanceDistributionRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListInstanceDistributionRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListInstanceDistributionRequestXLanguage) Value() string {
	return c.value
}

func (c ListInstanceDistributionRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstanceDistributionRequestXLanguage) UnmarshalJSON(b []byte) error {
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
