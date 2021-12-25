package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// Response Object
type UpdateVpcChannelV2Response struct {
	// VPC通道的名称。  长度为3 ~ 64位的字符串，字符串由中文、英文字母、数字、中划线、下划线组成，且只能以英文或中文开头。 > 中文字符必须为UTF-8或者unicode编码。

	Name string `json:"name"`
	// VPC通道中主机的端口号。  取值范围1 ~ 65535。

	Port int32 `json:"port"`
	// 分发算法。 - 1：加权轮询（wrr） - 2：加权最少连接（wleastconn） - 3：源地址哈希（source） - 4：URI哈希（uri）

	BalanceStrategy UpdateVpcChannelV2ResponseBalanceStrategy `json:"balance_strategy"`
	// VPC通道的成员类型。[site场景必须修改成IP类型](tag:Site) - ip - ecs

	MemberType UpdateVpcChannelV2ResponseMemberType `json:"member_type"`
	// VPC通道的字典编码  支持英文，数字，特殊字符（-_.）  暂不支持

	DictCode *string `json:"dict_code,omitempty"`
	// VPC通道的创建时间

	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`
	// VPC通道的编号

	Id *string `json:"id,omitempty"`
	// VPC通道的状态。 - 1：正常 - 2：异常

	Status *UpdateVpcChannelV2ResponseStatus `json:"status,omitempty"`
	// 后端云服务器组列表。  [暂不支持](tag:Site)

	MemberGroups   *[]MemberGroupInfo `json:"member_groups,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o UpdateVpcChannelV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpcChannelV2Response struct{}"
	}

	return strings.Join([]string{"UpdateVpcChannelV2Response", string(data)}, " ")
}

type UpdateVpcChannelV2ResponseBalanceStrategy struct {
	value int32
}

type UpdateVpcChannelV2ResponseBalanceStrategyEnum struct {
	E_1 UpdateVpcChannelV2ResponseBalanceStrategy
	E_2 UpdateVpcChannelV2ResponseBalanceStrategy
	E_3 UpdateVpcChannelV2ResponseBalanceStrategy
	E_4 UpdateVpcChannelV2ResponseBalanceStrategy
}

func GetUpdateVpcChannelV2ResponseBalanceStrategyEnum() UpdateVpcChannelV2ResponseBalanceStrategyEnum {
	return UpdateVpcChannelV2ResponseBalanceStrategyEnum{
		E_1: UpdateVpcChannelV2ResponseBalanceStrategy{
			value: 1,
		}, E_2: UpdateVpcChannelV2ResponseBalanceStrategy{
			value: 2,
		}, E_3: UpdateVpcChannelV2ResponseBalanceStrategy{
			value: 3,
		}, E_4: UpdateVpcChannelV2ResponseBalanceStrategy{
			value: 4,
		},
	}
}

func (c UpdateVpcChannelV2ResponseBalanceStrategy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateVpcChannelV2ResponseBalanceStrategy) UnmarshalJSON(b []byte) error {
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

type UpdateVpcChannelV2ResponseMemberType struct {
	value string
}

type UpdateVpcChannelV2ResponseMemberTypeEnum struct {
	IP  UpdateVpcChannelV2ResponseMemberType
	ECS UpdateVpcChannelV2ResponseMemberType
}

func GetUpdateVpcChannelV2ResponseMemberTypeEnum() UpdateVpcChannelV2ResponseMemberTypeEnum {
	return UpdateVpcChannelV2ResponseMemberTypeEnum{
		IP: UpdateVpcChannelV2ResponseMemberType{
			value: "ip",
		},
		ECS: UpdateVpcChannelV2ResponseMemberType{
			value: "ecs",
		},
	}
}

func (c UpdateVpcChannelV2ResponseMemberType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateVpcChannelV2ResponseMemberType) UnmarshalJSON(b []byte) error {
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

type UpdateVpcChannelV2ResponseStatus struct {
	value int32
}

type UpdateVpcChannelV2ResponseStatusEnum struct {
	E_1 UpdateVpcChannelV2ResponseStatus
	E_2 UpdateVpcChannelV2ResponseStatus
}

func GetUpdateVpcChannelV2ResponseStatusEnum() UpdateVpcChannelV2ResponseStatusEnum {
	return UpdateVpcChannelV2ResponseStatusEnum{
		E_1: UpdateVpcChannelV2ResponseStatus{
			value: 1,
		}, E_2: UpdateVpcChannelV2ResponseStatus{
			value: 2,
		},
	}
}

func (c UpdateVpcChannelV2ResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateVpcChannelV2ResponseStatus) UnmarshalJSON(b []byte) error {
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
