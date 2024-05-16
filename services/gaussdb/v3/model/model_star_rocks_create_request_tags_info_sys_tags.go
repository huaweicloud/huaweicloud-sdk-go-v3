package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type StarRocksCreateRequestTagsInfoSysTags struct {

	// 标签键。
	Key StarRocksCreateRequestTagsInfoSysTagsKey `json:"key"`

	// 标签值。
	Value StarRocksCreateRequestTagsInfoSysTagsValue `json:"value"`
}

func (o StarRocksCreateRequestTagsInfoSysTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StarRocksCreateRequestTagsInfoSysTags struct{}"
	}

	return strings.Join([]string{"StarRocksCreateRequestTagsInfoSysTags", string(data)}, " ")
}

type StarRocksCreateRequestTagsInfoSysTagsKey struct {
	value string
}

type StarRocksCreateRequestTagsInfoSysTagsKeyEnum struct {
	SYS_ENTERPRISE_PROJECT_ID StarRocksCreateRequestTagsInfoSysTagsKey
}

func GetStarRocksCreateRequestTagsInfoSysTagsKeyEnum() StarRocksCreateRequestTagsInfoSysTagsKeyEnum {
	return StarRocksCreateRequestTagsInfoSysTagsKeyEnum{
		SYS_ENTERPRISE_PROJECT_ID: StarRocksCreateRequestTagsInfoSysTagsKey{
			value: "_sys_enterprise_project_id",
		},
	}
}

func (c StarRocksCreateRequestTagsInfoSysTagsKey) Value() string {
	return c.value
}

func (c StarRocksCreateRequestTagsInfoSysTagsKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StarRocksCreateRequestTagsInfoSysTagsKey) UnmarshalJSON(b []byte) error {
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

type StarRocksCreateRequestTagsInfoSysTagsValue struct {
	value string
}

type StarRocksCreateRequestTagsInfoSysTagsValueEnum struct {
	E_0 StarRocksCreateRequestTagsInfoSysTagsValue
}

func GetStarRocksCreateRequestTagsInfoSysTagsValueEnum() StarRocksCreateRequestTagsInfoSysTagsValueEnum {
	return StarRocksCreateRequestTagsInfoSysTagsValueEnum{
		E_0: StarRocksCreateRequestTagsInfoSysTagsValue{
			value: "0",
		},
	}
}

func (c StarRocksCreateRequestTagsInfoSysTagsValue) Value() string {
	return c.value
}

func (c StarRocksCreateRequestTagsInfoSysTagsValue) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StarRocksCreateRequestTagsInfoSysTagsValue) UnmarshalJSON(b []byte) error {
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
