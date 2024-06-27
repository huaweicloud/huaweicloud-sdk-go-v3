package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ProtectObjectVo struct {

	// 防护对象id，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)，注意type为0的为互联网边界防护对象id，type为1的为VPC边界防护对象id。
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
