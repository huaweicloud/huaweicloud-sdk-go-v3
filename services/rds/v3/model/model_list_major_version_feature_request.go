package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListMajorVersionFeatureRequest Request Object
type ListMajorVersionFeatureRequest struct {

	// 语言。
	XLanguage *ListMajorVersionFeatureRequestXLanguage `json:"X-Language,omitempty"`

	// 版本，例如：2022_SE, 2022_EE, 2022_WEB。
	Version string `json:"version"`

	// 是否是单机版。
	Single bool `json:"single"`
}

func (o ListMajorVersionFeatureRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMajorVersionFeatureRequest struct{}"
	}

	return strings.Join([]string{"ListMajorVersionFeatureRequest", string(data)}, " ")
}

type ListMajorVersionFeatureRequestXLanguage struct {
	value string
}

type ListMajorVersionFeatureRequestXLanguageEnum struct {
	ZH_CN ListMajorVersionFeatureRequestXLanguage
	EN_US ListMajorVersionFeatureRequestXLanguage
}

func GetListMajorVersionFeatureRequestXLanguageEnum() ListMajorVersionFeatureRequestXLanguageEnum {
	return ListMajorVersionFeatureRequestXLanguageEnum{
		ZH_CN: ListMajorVersionFeatureRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListMajorVersionFeatureRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListMajorVersionFeatureRequestXLanguage) Value() string {
	return c.value
}

func (c ListMajorVersionFeatureRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMajorVersionFeatureRequestXLanguage) UnmarshalJSON(b []byte) error {
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
