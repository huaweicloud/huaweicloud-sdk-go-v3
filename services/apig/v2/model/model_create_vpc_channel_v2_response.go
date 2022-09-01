package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// Response Object
type CreateVpcChannelV2Response struct {

	// VPC通道的名称。  长度为3 ~ 64位的字符串，字符串由中文、英文字母、数字、中划线、下划线组成，且只能以英文或中文开头。 > 中文字符必须为UTF-8或者unicode编码。
	Name string `json:"name" xml:"name"`

	// VPC通道中主机的端口号。  取值范围1 ~ 65535，仅VPC通道类型为2时有效。  VPC通道类型为2时必选。
	Port *int32 `json:"port,omitempty" xml:"port"`

	// 分发算法。 - 1：加权轮询（wrr） - 2：加权最少连接（wleastconn） - 3：源地址哈希（source） - 4：URI哈希（uri）  VPC通道类型为2时必选。
	BalanceStrategy *CreateVpcChannelV2ResponseBalanceStrategy `json:"balance_strategy,omitempty" xml:"balance_strategy"`

	// VPC通道的成员类型。 - ip - ecs  VPC通道类型为2时必选。
	MemberType *CreateVpcChannelV2ResponseMemberType `json:"member_type,omitempty" xml:"member_type"`

	// VPC通道的创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty" xml:"create_time"`

	// VPC通道的编号
	Id *string `json:"id,omitempty" xml:"id"`

	// VPC通道的状态。 - 1：正常 - 2：异常
	Status *CreateVpcChannelV2ResponseStatus `json:"status,omitempty" xml:"status"`

	// 后端云服务器组列表。  暂不支持
	MemberGroups   *[]MemberGroupInfo `json:"member_groups,omitempty" xml:"member_groups"`
	HttpStatusCode int                `json:"-"`
}

func (o CreateVpcChannelV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcChannelV2Response struct{}"
	}

	return strings.Join([]string{"CreateVpcChannelV2Response", string(data)}, " ")
}

type CreateVpcChannelV2ResponseBalanceStrategy struct {
	value int32
}

type CreateVpcChannelV2ResponseBalanceStrategyEnum struct {
	E_1 CreateVpcChannelV2ResponseBalanceStrategy
	E_2 CreateVpcChannelV2ResponseBalanceStrategy
	E_3 CreateVpcChannelV2ResponseBalanceStrategy
	E_4 CreateVpcChannelV2ResponseBalanceStrategy
}

func GetCreateVpcChannelV2ResponseBalanceStrategyEnum() CreateVpcChannelV2ResponseBalanceStrategyEnum {
	return CreateVpcChannelV2ResponseBalanceStrategyEnum{
		E_1: CreateVpcChannelV2ResponseBalanceStrategy{
			value: 1,
		}, E_2: CreateVpcChannelV2ResponseBalanceStrategy{
			value: 2,
		}, E_3: CreateVpcChannelV2ResponseBalanceStrategy{
			value: 3,
		}, E_4: CreateVpcChannelV2ResponseBalanceStrategy{
			value: 4,
		},
	}
}

func (c CreateVpcChannelV2ResponseBalanceStrategy) Value() int32 {
	return c.value
}

func (c CreateVpcChannelV2ResponseBalanceStrategy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateVpcChannelV2ResponseBalanceStrategy) UnmarshalJSON(b []byte) error {
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

type CreateVpcChannelV2ResponseMemberType struct {
	value string
}

type CreateVpcChannelV2ResponseMemberTypeEnum struct {
	IP  CreateVpcChannelV2ResponseMemberType
	ECS CreateVpcChannelV2ResponseMemberType
}

func GetCreateVpcChannelV2ResponseMemberTypeEnum() CreateVpcChannelV2ResponseMemberTypeEnum {
	return CreateVpcChannelV2ResponseMemberTypeEnum{
		IP: CreateVpcChannelV2ResponseMemberType{
			value: "ip",
		},
		ECS: CreateVpcChannelV2ResponseMemberType{
			value: "ecs",
		},
	}
}

func (c CreateVpcChannelV2ResponseMemberType) Value() string {
	return c.value
}

func (c CreateVpcChannelV2ResponseMemberType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateVpcChannelV2ResponseMemberType) UnmarshalJSON(b []byte) error {
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

type CreateVpcChannelV2ResponseStatus struct {
	value int32
}

type CreateVpcChannelV2ResponseStatusEnum struct {
	E_1 CreateVpcChannelV2ResponseStatus
	E_2 CreateVpcChannelV2ResponseStatus
}

func GetCreateVpcChannelV2ResponseStatusEnum() CreateVpcChannelV2ResponseStatusEnum {
	return CreateVpcChannelV2ResponseStatusEnum{
		E_1: CreateVpcChannelV2ResponseStatus{
			value: 1,
		}, E_2: CreateVpcChannelV2ResponseStatus{
			value: 2,
		},
	}
}

func (c CreateVpcChannelV2ResponseStatus) Value() int32 {
	return c.value
}

func (c CreateVpcChannelV2ResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateVpcChannelV2ResponseStatus) UnmarshalJSON(b []byte) error {
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
