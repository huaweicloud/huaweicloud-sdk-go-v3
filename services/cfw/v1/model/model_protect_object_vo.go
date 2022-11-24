package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ProtectObjectVo struct {

	// 防护对象id
	ObjectId *string `json:"object_id,omitempty"`

	// 防护对象名称
	ObjectName *string `json:"object_name,omitempty"`

	// 防护对象类型：0 南北向，1 东西向护对象类型
	Type *ProtectObjectVoType `json:"type,omitempty"`
}

func (o ProtectObjectVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectObjectVo struct{}"
	}

	return strings.Join([]string{"ProtectObjectVo", string(data)}, " ")
}

type ProtectObjectVoType struct {
	value int32
}

type ProtectObjectVoTypeEnum struct {
	E_0 ProtectObjectVoType
	E_1 ProtectObjectVoType
}

func GetProtectObjectVoTypeEnum() ProtectObjectVoTypeEnum {
	return ProtectObjectVoTypeEnum{
		E_0: ProtectObjectVoType{
			value: 0,
		}, E_1: ProtectObjectVoType{
			value: 1,
		},
	}
}

func (c ProtectObjectVoType) Value() int32 {
	return c.value
}

func (c ProtectObjectVoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ProtectObjectVoType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
