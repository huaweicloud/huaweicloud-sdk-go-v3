package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AssociatePublicipsOption struct {

	// 功能说明：端口所属实例类型 取值范围：PORT、NATGW、VPN、ELB 约束：associate_instance_type和associate_instance_id都不为空时表示绑定实例 约束：associate_instance_type字段不允许为空 约束：双栈公网IP不允许修改绑定的实例
	AssociateInstanceType AssociatePublicipsOptionAssociateInstanceType `json:"associate_instance_type"`

	// 功能说明：端口所属实例ID，例如RDS实例ID 约束：associate_instance_type和associate_instance_id都不为空时表示绑定实例； 约束：associate_instance_id不允许为空 约束：双栈公网IP不允许修改绑定的实例
	AssociateInstanceId string `json:"associate_instance_id"`
}

func (o AssociatePublicipsOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociatePublicipsOption struct{}"
	}

	return strings.Join([]string{"AssociatePublicipsOption", string(data)}, " ")
}

type AssociatePublicipsOptionAssociateInstanceType struct {
	value string
}

type AssociatePublicipsOptionAssociateInstanceTypeEnum struct {
	PORT  AssociatePublicipsOptionAssociateInstanceType
	NATGW AssociatePublicipsOptionAssociateInstanceType
	VPN   AssociatePublicipsOptionAssociateInstanceType
	ELB   AssociatePublicipsOptionAssociateInstanceType
}

func GetAssociatePublicipsOptionAssociateInstanceTypeEnum() AssociatePublicipsOptionAssociateInstanceTypeEnum {
	return AssociatePublicipsOptionAssociateInstanceTypeEnum{
		PORT: AssociatePublicipsOptionAssociateInstanceType{
			value: "PORT",
		},
		NATGW: AssociatePublicipsOptionAssociateInstanceType{
			value: "NATGW",
		},
		VPN: AssociatePublicipsOptionAssociateInstanceType{
			value: "VPN",
		},
		ELB: AssociatePublicipsOptionAssociateInstanceType{
			value: "ELB",
		},
	}
}

func (c AssociatePublicipsOptionAssociateInstanceType) Value() string {
	return c.value
}

func (c AssociatePublicipsOptionAssociateInstanceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AssociatePublicipsOptionAssociateInstanceType) UnmarshalJSON(b []byte) error {
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
