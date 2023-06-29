package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ProjectVpcCreate struct {

	// VPC通道的名称。  长度为3 ~ 64位的字符串，字符串由中文、英文字母、数字、中划线、下划线组成，且只能以英文或中文开头。 > 中文字符必须为UTF-8或者unicode编码。
	Name string `json:"name"`

	// VPC通道中主机的端口号。  取值范围1 ~ 65535。
	Port int32 `json:"port"`

	// 分发算法。 - 1：加权轮询（wrr） - 2：加权最少连接（wleastconn） - 3：源地址哈希（source） - 4：URI哈希（uri）
	BalanceStrategy ProjectVpcCreateBalanceStrategy `json:"balance_strategy"`

	// VPC通道的成员类型。[site场景必须修改成IP类型](tag:Site) - ip - ecs
	MemberType ProjectVpcCreateMemberType `json:"member_type"`

	// VPC通道的字典编码  支持英文，数字，特殊字符（-_.）  暂不支持
	DictCode *string `json:"dict_code,omitempty"`

	// VPC通道后端服务器组列表
	MemberGroups *[]MemberGroupCreate `json:"member_groups,omitempty"`

	// VPC后端实例列表。
	Members *[]MemberInfo `json:"members,omitempty"`

	VpcHealthConfig *VpcHealthConfig `json:"vpc_health_config,omitempty"`

	// 关联实例列表。至少包含一个实例编号，最多10个，如需扩大配额请联系技术工程师修改PROJECT_VPC_OPERATOR_NUM_LIMIT配置。
	InstanceIds []string `json:"instance_ids"`
}

func (o ProjectVpcCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectVpcCreate struct{}"
	}

	return strings.Join([]string{"ProjectVpcCreate", string(data)}, " ")
}

type ProjectVpcCreateBalanceStrategy struct {
	value int32
}

type ProjectVpcCreateBalanceStrategyEnum struct {
	E_1 ProjectVpcCreateBalanceStrategy
	E_2 ProjectVpcCreateBalanceStrategy
	E_3 ProjectVpcCreateBalanceStrategy
	E_4 ProjectVpcCreateBalanceStrategy
}

func GetProjectVpcCreateBalanceStrategyEnum() ProjectVpcCreateBalanceStrategyEnum {
	return ProjectVpcCreateBalanceStrategyEnum{
		E_1: ProjectVpcCreateBalanceStrategy{
			value: 1,
		}, E_2: ProjectVpcCreateBalanceStrategy{
			value: 2,
		}, E_3: ProjectVpcCreateBalanceStrategy{
			value: 3,
		}, E_4: ProjectVpcCreateBalanceStrategy{
			value: 4,
		},
	}
}

func (c ProjectVpcCreateBalanceStrategy) Value() int32 {
	return c.value
}

func (c ProjectVpcCreateBalanceStrategy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ProjectVpcCreateBalanceStrategy) UnmarshalJSON(b []byte) error {
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

type ProjectVpcCreateMemberType struct {
	value string
}

type ProjectVpcCreateMemberTypeEnum struct {
	IP  ProjectVpcCreateMemberType
	ECS ProjectVpcCreateMemberType
}

func GetProjectVpcCreateMemberTypeEnum() ProjectVpcCreateMemberTypeEnum {
	return ProjectVpcCreateMemberTypeEnum{
		IP: ProjectVpcCreateMemberType{
			value: "ip",
		},
		ECS: ProjectVpcCreateMemberType{
			value: "ecs",
		},
	}
}

func (c ProjectVpcCreateMemberType) Value() string {
	return c.value
}

func (c ProjectVpcCreateMemberType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ProjectVpcCreateMemberType) UnmarshalJSON(b []byte) error {
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
