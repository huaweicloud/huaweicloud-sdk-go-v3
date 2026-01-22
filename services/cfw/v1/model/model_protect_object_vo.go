package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ProtectObjectVo struct {

	// **参数解释**： 防护对象id，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id。 **取值范围**： 不涉及
	ObjectId *string `json:"object_id,omitempty"`

	// **参数解释**： 防护对象名称 **取值范围**： 不涉及
	ObjectName *string `json:"object_name,omitempty"`

	// **参数解释**： 防护对象类型 **取值范围**： - 0：南北向 - 1：东西向
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
