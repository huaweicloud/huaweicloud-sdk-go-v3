package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowChildNumRequest Request Object
type ShowChildNumRequest struct {

	// ddm或者gaussdbv5数据库的实例id
	InstanceId string `json:"instance_id"`

	// 局点，默认是当前region
	Region *string `json:"region,omitempty"`

	// 请求语言类型。
	XLanguage *ShowChildNumRequestXLanguage `json:"X-Language,omitempty"`

	// 数据库实例的类型
	DbType ShowChildNumRequestDbType `json:"db_type"`
}

func (o ShowChildNumRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowChildNumRequest struct{}"
	}

	return strings.Join([]string{"ShowChildNumRequest", string(data)}, " ")
}

type ShowChildNumRequestXLanguage struct {
	value string
}

type ShowChildNumRequestXLanguageEnum struct {
	EN_US ShowChildNumRequestXLanguage
	ZH_CN ShowChildNumRequestXLanguage
}

func GetShowChildNumRequestXLanguageEnum() ShowChildNumRequestXLanguageEnum {
	return ShowChildNumRequestXLanguageEnum{
		EN_US: ShowChildNumRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowChildNumRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowChildNumRequestXLanguage) Value() string {
	return c.value
}

func (c ShowChildNumRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowChildNumRequestXLanguage) UnmarshalJSON(b []byte) error {
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

type ShowChildNumRequestDbType struct {
	value string
}

type ShowChildNumRequestDbTypeEnum struct {
	GAUSSDBV5 ShowChildNumRequestDbType
	DDM       ShowChildNumRequestDbType
}

func GetShowChildNumRequestDbTypeEnum() ShowChildNumRequestDbTypeEnum {
	return ShowChildNumRequestDbTypeEnum{
		GAUSSDBV5: ShowChildNumRequestDbType{
			value: "gaussdbv5",
		},
		DDM: ShowChildNumRequestDbType{
			value: "ddm",
		},
	}
}

func (c ShowChildNumRequestDbType) Value() string {
	return c.value
}

func (c ShowChildNumRequestDbType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowChildNumRequestDbType) UnmarshalJSON(b []byte) error {
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
