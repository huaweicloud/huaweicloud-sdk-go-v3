package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type VpcBase struct {
	// VPC通道的名称。  长度为3 ~ 64位的字符串，字符串由中文、英文字母、数字、中划线、下划线组成，且只能以英文或中文开头。 > 中文字符必须为UTF-8或者unicode编码。

	Name string `json:"name"`
	// VPC通道中主机的端口号。  取值范围1 ~ 65535，仅VPC通道类型为2时有效。  VPC通道类型为2时必选。

	Port *int32 `json:"port,omitempty"`
	// 分发算法。 - 1：加权轮询（wrr） - 2：加权最少连接（wleastconn） - 3：源地址哈希（source） - 4：URI哈希（uri）  VPC通道类型为2时必选。

	BalanceStrategy *VpcBaseBalanceStrategy `json:"balance_strategy,omitempty"`
	// VPC通道的成员类型。 - ip - ecs  VPC通道类型为2时必选。

	MemberType *VpcBaseMemberType `json:"member_type,omitempty"`
}

func (o VpcBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcBase struct{}"
	}

	return strings.Join([]string{"VpcBase", string(data)}, " ")
}

type VpcBaseBalanceStrategy struct {
	value int32
}

type VpcBaseBalanceStrategyEnum struct {
	E_1 VpcBaseBalanceStrategy
	E_2 VpcBaseBalanceStrategy
	E_3 VpcBaseBalanceStrategy
	E_4 VpcBaseBalanceStrategy
}

func GetVpcBaseBalanceStrategyEnum() VpcBaseBalanceStrategyEnum {
	return VpcBaseBalanceStrategyEnum{
		E_1: VpcBaseBalanceStrategy{
			value: 1,
		}, E_2: VpcBaseBalanceStrategy{
			value: 2,
		}, E_3: VpcBaseBalanceStrategy{
			value: 3,
		}, E_4: VpcBaseBalanceStrategy{
			value: 4,
		},
	}
}

func (c VpcBaseBalanceStrategy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VpcBaseBalanceStrategy) UnmarshalJSON(b []byte) error {
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

type VpcBaseMemberType struct {
	value string
}

type VpcBaseMemberTypeEnum struct {
	IP  VpcBaseMemberType
	ECS VpcBaseMemberType
}

func GetVpcBaseMemberTypeEnum() VpcBaseMemberTypeEnum {
	return VpcBaseMemberTypeEnum{
		IP: VpcBaseMemberType{
			value: "ip",
		},
		ECS: VpcBaseMemberType{
			value: "ecs",
		},
	}
}

func (c VpcBaseMemberType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VpcBaseMemberType) UnmarshalJSON(b []byte) error {
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
