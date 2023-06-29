package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type VpcCreate struct {

	// VPC通道的名称。  长度为3 ~ 64位的字符串，字符串由中文、英文字母、数字、中划线、下划线组成，且只能以英文或中文开头。 > 中文字符必须为UTF-8或者unicode编码。
	Name string `json:"name"`

	// VPC通道中主机的端口号。  取值范围1 ~ 65535。
	Port int32 `json:"port"`

	// 分发算法。 - 1：加权轮询（wrr） - 2：加权最少连接（wleastconn） - 3：源地址哈希（source） - 4：URI哈希（uri）
	BalanceStrategy VpcCreateBalanceStrategy `json:"balance_strategy"`

	// VPC通道的成员类型。 - ip - ecs
	MemberType VpcCreateMemberType `json:"member_type"`

	// vpc通道类型，默认为服务器类型。 - 2：服务器类型 - 3：微服务类型
	Type *int32 `json:"type,omitempty"`

	// VPC通道的字典编码  支持英文，数字，特殊字符（-_.）  暂不支持
	DictCode *string `json:"dict_code,omitempty"`

	// VPC通道后端服务器组列表
	MemberGroups *[]MemberGroupCreate `json:"member_groups,omitempty"`

	// VPC后端实例列表。
	Members *[]MemberInfo `json:"members,omitempty"`

	VpcHealthConfig *VpcHealthConfig `json:"vpc_health_config,omitempty"`

	MicroserviceInfo *MicroServiceCreate `json:"microservice_info,omitempty"`
}

func (o VpcCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcCreate struct{}"
	}

	return strings.Join([]string{"VpcCreate", string(data)}, " ")
}

type VpcCreateBalanceStrategy struct {
	value int32
}

type VpcCreateBalanceStrategyEnum struct {
	E_1 VpcCreateBalanceStrategy
	E_2 VpcCreateBalanceStrategy
	E_3 VpcCreateBalanceStrategy
	E_4 VpcCreateBalanceStrategy
}

func GetVpcCreateBalanceStrategyEnum() VpcCreateBalanceStrategyEnum {
	return VpcCreateBalanceStrategyEnum{
		E_1: VpcCreateBalanceStrategy{
			value: 1,
		}, E_2: VpcCreateBalanceStrategy{
			value: 2,
		}, E_3: VpcCreateBalanceStrategy{
			value: 3,
		}, E_4: VpcCreateBalanceStrategy{
			value: 4,
		},
	}
}

func (c VpcCreateBalanceStrategy) Value() int32 {
	return c.value
}

func (c VpcCreateBalanceStrategy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VpcCreateBalanceStrategy) UnmarshalJSON(b []byte) error {
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

type VpcCreateMemberType struct {
	value string
}

type VpcCreateMemberTypeEnum struct {
	IP  VpcCreateMemberType
	ECS VpcCreateMemberType
}

func GetVpcCreateMemberTypeEnum() VpcCreateMemberTypeEnum {
	return VpcCreateMemberTypeEnum{
		IP: VpcCreateMemberType{
			value: "ip",
		},
		ECS: VpcCreateMemberType{
			value: "ecs",
		},
	}
}

func (c VpcCreateMemberType) Value() string {
	return c.value
}

func (c VpcCreateMemberType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VpcCreateMemberType) UnmarshalJSON(b []byte) error {
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
