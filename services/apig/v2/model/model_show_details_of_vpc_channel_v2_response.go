package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// Response Object
type ShowDetailsOfVpcChannelV2Response struct {

	// VPC通道的名称。  长度为3 ~ 64位的字符串，字符串由中文、英文字母、数字、中划线、下划线组成，且只能以英文或中文开头。 > 中文字符必须为UTF-8或者unicode编码。
	Name string `json:"name"`

	// VPC通道中主机的端口号。  取值范围1 ~ 65535。
	Port int32 `json:"port"`

	// 分发算法。 - 1：加权轮询（wrr） - 2：加权最少连接（wleastconn） - 3：源地址哈希（source） - 4：URI哈希（uri）
	BalanceStrategy ShowDetailsOfVpcChannelV2ResponseBalanceStrategy `json:"balance_strategy"`

	// VPC通道的成员类型。 - ip - ecs
	MemberType ShowDetailsOfVpcChannelV2ResponseMemberType `json:"member_type"`

	// VPC通道的字典编码  支持英文，数字，特殊字符（-_.）  暂不支持
	DictCode *string `json:"dict_code,omitempty"`

	// VPC通道的创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// VPC通道的编号
	Id *string `json:"id,omitempty"`

	// VPC通道的状态。 - 1：正常 - 2：异常
	Status *ShowDetailsOfVpcChannelV2ResponseStatus `json:"status,omitempty"`

	// 后端云服务器组列表。
	MemberGroups *[]MemberGroupInfo `json:"member_groups,omitempty"`

	MicroserviceInfo *MicroServiceInfo `json:"microservice_info,omitempty"`

	// vpc通道类型。 - 2：服务器类型 - 3：微服务类型
	Type *int32 `json:"type,omitempty"`

	// 后端实例列表。
	Members *[]VpcMemberInfo `json:"members,omitempty"`

	VpcHealthConfig *VpcHealthConfigInfo `json:"vpc_health_config,omitempty"`
	HttpStatusCode  int                  `json:"-"`
}

func (o ShowDetailsOfVpcChannelV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfVpcChannelV2Response struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfVpcChannelV2Response", string(data)}, " ")
}

type ShowDetailsOfVpcChannelV2ResponseBalanceStrategy struct {
	value int32
}

type ShowDetailsOfVpcChannelV2ResponseBalanceStrategyEnum struct {
	E_1 ShowDetailsOfVpcChannelV2ResponseBalanceStrategy
	E_2 ShowDetailsOfVpcChannelV2ResponseBalanceStrategy
	E_3 ShowDetailsOfVpcChannelV2ResponseBalanceStrategy
	E_4 ShowDetailsOfVpcChannelV2ResponseBalanceStrategy
}

func GetShowDetailsOfVpcChannelV2ResponseBalanceStrategyEnum() ShowDetailsOfVpcChannelV2ResponseBalanceStrategyEnum {
	return ShowDetailsOfVpcChannelV2ResponseBalanceStrategyEnum{
		E_1: ShowDetailsOfVpcChannelV2ResponseBalanceStrategy{
			value: 1,
		}, E_2: ShowDetailsOfVpcChannelV2ResponseBalanceStrategy{
			value: 2,
		}, E_3: ShowDetailsOfVpcChannelV2ResponseBalanceStrategy{
			value: 3,
		}, E_4: ShowDetailsOfVpcChannelV2ResponseBalanceStrategy{
			value: 4,
		},
	}
}

func (c ShowDetailsOfVpcChannelV2ResponseBalanceStrategy) Value() int32 {
	return c.value
}

func (c ShowDetailsOfVpcChannelV2ResponseBalanceStrategy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailsOfVpcChannelV2ResponseBalanceStrategy) UnmarshalJSON(b []byte) error {
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

type ShowDetailsOfVpcChannelV2ResponseMemberType struct {
	value string
}

type ShowDetailsOfVpcChannelV2ResponseMemberTypeEnum struct {
	IP  ShowDetailsOfVpcChannelV2ResponseMemberType
	ECS ShowDetailsOfVpcChannelV2ResponseMemberType
}

func GetShowDetailsOfVpcChannelV2ResponseMemberTypeEnum() ShowDetailsOfVpcChannelV2ResponseMemberTypeEnum {
	return ShowDetailsOfVpcChannelV2ResponseMemberTypeEnum{
		IP: ShowDetailsOfVpcChannelV2ResponseMemberType{
			value: "ip",
		},
		ECS: ShowDetailsOfVpcChannelV2ResponseMemberType{
			value: "ecs",
		},
	}
}

func (c ShowDetailsOfVpcChannelV2ResponseMemberType) Value() string {
	return c.value
}

func (c ShowDetailsOfVpcChannelV2ResponseMemberType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailsOfVpcChannelV2ResponseMemberType) UnmarshalJSON(b []byte) error {
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

type ShowDetailsOfVpcChannelV2ResponseStatus struct {
	value int32
}

type ShowDetailsOfVpcChannelV2ResponseStatusEnum struct {
	E_1 ShowDetailsOfVpcChannelV2ResponseStatus
	E_2 ShowDetailsOfVpcChannelV2ResponseStatus
}

func GetShowDetailsOfVpcChannelV2ResponseStatusEnum() ShowDetailsOfVpcChannelV2ResponseStatusEnum {
	return ShowDetailsOfVpcChannelV2ResponseStatusEnum{
		E_1: ShowDetailsOfVpcChannelV2ResponseStatus{
			value: 1,
		}, E_2: ShowDetailsOfVpcChannelV2ResponseStatus{
			value: 2,
		},
	}
}

func (c ShowDetailsOfVpcChannelV2ResponseStatus) Value() int32 {
	return c.value
}

func (c ShowDetailsOfVpcChannelV2ResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailsOfVpcChannelV2ResponseStatus) UnmarshalJSON(b []byte) error {
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
