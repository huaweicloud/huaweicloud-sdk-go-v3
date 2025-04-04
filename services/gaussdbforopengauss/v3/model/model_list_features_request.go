package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListFeaturesRequest Request Object
type ListFeaturesRequest struct {

	// 语言。
	XLanguage *ListFeaturesRequestXLanguage `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ListFeaturesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFeaturesRequest struct{}"
	}

	return strings.Join([]string{"ListFeaturesRequest", string(data)}, " ")
}

type ListFeaturesRequestXLanguage struct {
	value string
}

type ListFeaturesRequestXLanguageEnum struct {
	ZH_CN ListFeaturesRequestXLanguage
	EN_US ListFeaturesRequestXLanguage
}

func GetListFeaturesRequestXLanguageEnum() ListFeaturesRequestXLanguageEnum {
	return ListFeaturesRequestXLanguageEnum{
		ZH_CN: ListFeaturesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListFeaturesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListFeaturesRequestXLanguage) Value() string {
	return c.value
}

func (c ListFeaturesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFeaturesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
