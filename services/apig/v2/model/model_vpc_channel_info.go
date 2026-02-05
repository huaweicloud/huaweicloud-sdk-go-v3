package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type VpcChannelInfo struct {

	// VPC通道的名称。  长度为3 ~ 64位的字符串，字符串由中文、英文字母、数字、中划线、下划线、点组成，且只能以英文或中文开头。 > 中文字符必须为UTF-8或者unicode编码。
	Name string `json:"name"`

	// VPC通道中主机的端口号。  取值范围1 ~ 65535。
	Port int32 `json:"port"`

	// 分发算法。 - 1：加权轮询（wrr） - 2：加权最少连接（wleastconn） - 3：源地址哈希（source） - 4：URI哈希（uri）
	BalanceStrategy VpcChannelInfoBalanceStrategy `json:"balance_strategy"`

	// VPC通道的成员类型。 - ip - ecs
	MemberType VpcChannelInfoMemberType `json:"member_type"`

	// vpc通道类型，默认为服务器类型。 - 2：服务器类型 - 3：微服务类型  当vpc_channel_type字段为空时，负载通道类型由type字段控制： 当type不为3或microservice_info为空，VCP通道类型默认为服务器类型。 当type=3，microservice_info不为空，VPC通道类型为微服务类型。  修改负载通道时vpc通道类型不会修改，直接使用原有的vpc通道类型。  请优先使用vpc_channel_type字段指定负载通道类型。
	Type *int32 `json:"type,omitempty"`

	// vpc通道类型。 - builtin：服务器类型 - microservice： 微服务类型 - reference：引用负载通道类型  当vpc_channel_type为空时，负载通道类型取决于type字段的取值。 当vpc_channel_type不为空，但type字段非空或不为0时，当vpc_channel_type的指定类型与type字段指定的类型冲突时会校验报错。 当vpc_channel_type不为空，且type字段为空或等于0时，直接使用vpc_channel_type字段的值指定负载通道类型。  修改负载通道时vpc通道类型不会修改，直接使用原有的vpc通道类型。
	VpcChannelType *VpcChannelInfoVpcChannelType `json:"vpc_channel_type,omitempty"`

	// VPC通道的字典编码  支持英文，数字，特殊字符（-_.）  暂不支持
	DictCode *string `json:"dict_code,omitempty"`

	// VPC通道的创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// VPC通道的编号
	Id *string `json:"id,omitempty"`

	// VPC通道的状态。 - 1：正常 - 2：异常
	Status *VpcChannelInfoStatus `json:"status,omitempty"`

	// 后端云服务器组列表。
	MemberGroups *[]MemberGroupInfo `json:"member_groups,omitempty"`

	MicroserviceInfo *MicroServiceInfo `json:"microservice_info,omitempty"`
}

func (o VpcChannelInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcChannelInfo struct{}"
	}

	return strings.Join([]string{"VpcChannelInfo", string(data)}, " ")
}

type VpcChannelInfoBalanceStrategy struct {
	value int32
}

type VpcChannelInfoBalanceStrategyEnum struct {
	E_1 VpcChannelInfoBalanceStrategy
	E_2 VpcChannelInfoBalanceStrategy
	E_3 VpcChannelInfoBalanceStrategy
	E_4 VpcChannelInfoBalanceStrategy
}

func GetVpcChannelInfoBalanceStrategyEnum() VpcChannelInfoBalanceStrategyEnum {
	return VpcChannelInfoBalanceStrategyEnum{
		E_1: VpcChannelInfoBalanceStrategy{
			value: 1,
		}, E_2: VpcChannelInfoBalanceStrategy{
			value: 2,
		}, E_3: VpcChannelInfoBalanceStrategy{
			value: 3,
		}, E_4: VpcChannelInfoBalanceStrategy{
			value: 4,
		},
	}
}

func (c VpcChannelInfoBalanceStrategy) Value() int32 {
	return c.value
}

func (c VpcChannelInfoBalanceStrategy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VpcChannelInfoBalanceStrategy) UnmarshalJSON(b []byte) error {
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

type VpcChannelInfoMemberType struct {
	value string
}

type VpcChannelInfoMemberTypeEnum struct {
	IP  VpcChannelInfoMemberType
	ECS VpcChannelInfoMemberType
}

func GetVpcChannelInfoMemberTypeEnum() VpcChannelInfoMemberTypeEnum {
	return VpcChannelInfoMemberTypeEnum{
		IP: VpcChannelInfoMemberType{
			value: "ip",
		},
		ECS: VpcChannelInfoMemberType{
			value: "ecs",
		},
	}
}

func (c VpcChannelInfoMemberType) Value() string {
	return c.value
}

func (c VpcChannelInfoMemberType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VpcChannelInfoMemberType) UnmarshalJSON(b []byte) error {
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

type VpcChannelInfoVpcChannelType struct {
	value string
}

type VpcChannelInfoVpcChannelTypeEnum struct {
	BUILTIN      VpcChannelInfoVpcChannelType
	MICROSERVICE VpcChannelInfoVpcChannelType
	REFERENCE    VpcChannelInfoVpcChannelType
}

func GetVpcChannelInfoVpcChannelTypeEnum() VpcChannelInfoVpcChannelTypeEnum {
	return VpcChannelInfoVpcChannelTypeEnum{
		BUILTIN: VpcChannelInfoVpcChannelType{
			value: "builtin",
		},
		MICROSERVICE: VpcChannelInfoVpcChannelType{
			value: "microservice",
		},
		REFERENCE: VpcChannelInfoVpcChannelType{
			value: "reference",
		},
	}
}

func (c VpcChannelInfoVpcChannelType) Value() string {
	return c.value
}

func (c VpcChannelInfoVpcChannelType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VpcChannelInfoVpcChannelType) UnmarshalJSON(b []byte) error {
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

type VpcChannelInfoStatus struct {
	value int32
}

type VpcChannelInfoStatusEnum struct {
	E_1 VpcChannelInfoStatus
	E_2 VpcChannelInfoStatus
}

func GetVpcChannelInfoStatusEnum() VpcChannelInfoStatusEnum {
	return VpcChannelInfoStatusEnum{
		E_1: VpcChannelInfoStatus{
			value: 1,
		}, E_2: VpcChannelInfoStatus{
			value: 2,
		},
	}
}

func (c VpcChannelInfoStatus) Value() int32 {
	return c.value
}

func (c VpcChannelInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VpcChannelInfoStatus) UnmarshalJSON(b []byte) error {
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
