package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 场地信息
type UpdateLocation struct {

	// 场地名称（已废弃），传入该参数不会再生效，新建站点也不会再返回该字段
	Name *string `json:"name,omitempty"`

	// 场地所在国家
	Country *UpdateLocationCountry `json:"country,omitempty"`

	// 场地所在省/自治区/直辖市
	Province *string `json:"province,omitempty"`

	// 场地所在市/区
	City *string `json:"city,omitempty"`

	// 场地所在区/县
	District *string `json:"district,omitempty"`

	Condition *UpdateCondition `json:"condition,omitempty"`

	// 场地描述，最大支持长度为255个字节，不允许包含<>
	Description *string `json:"description,omitempty"`
}

func (o UpdateLocation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLocation struct{}"
	}

	return strings.Join([]string{"UpdateLocation", string(data)}, " ")
}

type UpdateLocationCountry struct {
	value string
}

type UpdateLocationCountryEnum struct {
	CN UpdateLocationCountry
}

func GetUpdateLocationCountryEnum() UpdateLocationCountryEnum {
	return UpdateLocationCountryEnum{
		CN: UpdateLocationCountry{
			value: "CN",
		},
	}
}

func (c UpdateLocationCountry) Value() string {
	return c.value
}

func (c UpdateLocationCountry) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateLocationCountry) UnmarshalJSON(b []byte) error {
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
