package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AddProtectAccessLevel struct {

	// 提交权限 0:任何人不允许提交，30:开发者及管理员可提交，40:管理员可提交
	PushAccessLevel AddProtectAccessLevelPushAccessLevel `json:"push_access_level"`

	// 合并权限 0:任何人不允许合并，30:开发者及管理员可合并，40:管理员可合并,合并权限必须大于等于提交权限
	MergeAccessLevel AddProtectAccessLevelMergeAccessLevel `json:"merge_access_level"`
}

func (o AddProtectAccessLevel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddProtectAccessLevel struct{}"
	}

	return strings.Join([]string{"AddProtectAccessLevel", string(data)}, " ")
}

type AddProtectAccessLevelPushAccessLevel struct {
	value int32
}

type AddProtectAccessLevelPushAccessLevelEnum struct {
	E_0  AddProtectAccessLevelPushAccessLevel
	E_30 AddProtectAccessLevelPushAccessLevel
	E_40 AddProtectAccessLevelPushAccessLevel
}

func GetAddProtectAccessLevelPushAccessLevelEnum() AddProtectAccessLevelPushAccessLevelEnum {
	return AddProtectAccessLevelPushAccessLevelEnum{
		E_0: AddProtectAccessLevelPushAccessLevel{
			value: 0,
		}, E_30: AddProtectAccessLevelPushAccessLevel{
			value: 30,
		}, E_40: AddProtectAccessLevelPushAccessLevel{
			value: 40,
		},
	}
}

func (c AddProtectAccessLevelPushAccessLevel) Value() int32 {
	return c.value
}

func (c AddProtectAccessLevelPushAccessLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddProtectAccessLevelPushAccessLevel) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type AddProtectAccessLevelMergeAccessLevel struct {
	value int32
}

type AddProtectAccessLevelMergeAccessLevelEnum struct {
	E_0  AddProtectAccessLevelMergeAccessLevel
	E_30 AddProtectAccessLevelMergeAccessLevel
	E_40 AddProtectAccessLevelMergeAccessLevel
}

func GetAddProtectAccessLevelMergeAccessLevelEnum() AddProtectAccessLevelMergeAccessLevelEnum {
	return AddProtectAccessLevelMergeAccessLevelEnum{
		E_0: AddProtectAccessLevelMergeAccessLevel{
			value: 0,
		}, E_30: AddProtectAccessLevelMergeAccessLevel{
			value: 30,
		}, E_40: AddProtectAccessLevelMergeAccessLevel{
			value: 40,
		},
	}
}

func (c AddProtectAccessLevelMergeAccessLevel) Value() int32 {
	return c.value
}

func (c AddProtectAccessLevelMergeAccessLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddProtectAccessLevelMergeAccessLevel) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
