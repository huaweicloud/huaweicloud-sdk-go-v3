package model

import (
	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type VpcChannelInfo struct {
	// VPC通道的名称。  长度为3 ~ 64位的字符串，字符串由中文、英文字母、数字、中划线、下划线组成，且只能以英文或中文开头。 > 中文字符必须为UTF-8或者unicode编码。

	Name string `json:"name"`
	// VPC通道中主机的端口号。  取值范围1 ~ 65535，仅VPC通道类型为2时有效。  VPC通道类型为2时必选。

	Port *int32 `json:"port,omitempty"`
	// 分发算法。 - 1：加权轮询（wrr） - 2：加权最少连接（wleastconn） - 3：源地址哈希（source） - 4：URI哈希（uri）  VPC通道类型为2时必选。

	BalanceStrategy *VpcChannelInfoBalanceStrategy `json:"balance_strategy,omitempty"`
	// VPC通道的成员类型。 - ip - ecs  VPC通道类型为2时必选。

	MemberType *VpcChannelInfoMemberType `json:"member_type,omitempty"`
	// VPC通道的创建时间

	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`
	// VPC通道的编号

	Id *string `json:"id,omitempty"`
	// VPC通道的状态。 - 1：正常 - 2：异常

	Status *VpcChannelInfoStatus `json:"status,omitempty"`
	// 后端云服务器组列表。  暂不支持

	MemberGroups *[]MemberGroupInfo `json:"member_groups,omitempty"`
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

func (c VpcChannelInfoBalanceStrategy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VpcChannelInfoBalanceStrategy) UnmarshalJSON(b []byte) error {
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

func (c VpcChannelInfoMemberType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VpcChannelInfoMemberType) UnmarshalJSON(b []byte) error {
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

func (c VpcChannelInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VpcChannelInfoStatus) UnmarshalJSON(b []byte) error {
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
