package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateProtectedIpInPolicyBody 更新全量防护对象的请求体
type UpdateProtectedIpInPolicyBody struct {

	// 防护ip的id
	Id string `json:"id"`

	// 防护ip
	Ip string `json:"ip"`

	// 类型。VPN：VPN；NAT：NAT；VIP：VIP；CCI：CCI；EIP：弹性公网IP；ELB：弹性负载均衡；REROUTING_IP：REROUTING_IP；SubEni：SubEni；NetInterFace：NetInterFace；
	Type UpdateProtectedIpInPolicyBodyType `json:"type"`

	// 名字
	Name *string `json:"name,omitempty"`
}

func (o UpdateProtectedIpInPolicyBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProtectedIpInPolicyBody struct{}"
	}

	return strings.Join([]string{"UpdateProtectedIpInPolicyBody", string(data)}, " ")
}

type UpdateProtectedIpInPolicyBodyType struct {
	value string
}

type UpdateProtectedIpInPolicyBodyTypeEnum struct {
	VPN            UpdateProtectedIpInPolicyBodyType
	NAT            UpdateProtectedIpInPolicyBodyType
	VIP            UpdateProtectedIpInPolicyBodyType
	CCI            UpdateProtectedIpInPolicyBodyType
	EIP            UpdateProtectedIpInPolicyBodyType
	ELB            UpdateProtectedIpInPolicyBodyType
	REROUTING_IP   UpdateProtectedIpInPolicyBodyType
	SUB_ENI        UpdateProtectedIpInPolicyBodyType
	NET_INTER_FACE UpdateProtectedIpInPolicyBodyType
}

func GetUpdateProtectedIpInPolicyBodyTypeEnum() UpdateProtectedIpInPolicyBodyTypeEnum {
	return UpdateProtectedIpInPolicyBodyTypeEnum{
		VPN: UpdateProtectedIpInPolicyBodyType{
			value: "VPN",
		},
		NAT: UpdateProtectedIpInPolicyBodyType{
			value: "NAT",
		},
		VIP: UpdateProtectedIpInPolicyBodyType{
			value: "VIP",
		},
		CCI: UpdateProtectedIpInPolicyBodyType{
			value: "CCI",
		},
		EIP: UpdateProtectedIpInPolicyBodyType{
			value: "EIP",
		},
		ELB: UpdateProtectedIpInPolicyBodyType{
			value: "ELB",
		},
		REROUTING_IP: UpdateProtectedIpInPolicyBodyType{
			value: "REROUTING_IP",
		},
		SUB_ENI: UpdateProtectedIpInPolicyBodyType{
			value: "SubEni",
		},
		NET_INTER_FACE: UpdateProtectedIpInPolicyBodyType{
			value: "NetInterFace",
		},
	}
}

func (c UpdateProtectedIpInPolicyBodyType) Value() string {
	return c.value
}

func (c UpdateProtectedIpInPolicyBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateProtectedIpInPolicyBodyType) UnmarshalJSON(b []byte) error {
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
