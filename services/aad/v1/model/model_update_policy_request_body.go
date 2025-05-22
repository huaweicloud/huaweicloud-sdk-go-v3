package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdatePolicyRequestBody 更新策略的请求体
type UpdatePolicyRequestBody struct {

	// 策略名
	Name *string `json:"name,omitempty"`

	// 清洗阈值
	Threshold *int32 `json:"threshold,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// udp协议设置。block：封禁，unblock：不封禁，limiting：限速
	Udp *UpdatePolicyRequestBodyUdp `json:"udp,omitempty"`

	// tcp协议设置。block：封禁，unblock：不封禁，limiting：限速
	Tcp *UpdatePolicyRequestBodyTcp `json:"tcp,omitempty"`

	// icmp协议设置。block：封禁，unblock：不封禁，limiting：限速
	Icmp *UpdatePolicyRequestBodyIcmp `json:"icmp,omitempty"`

	// other协议设置。block：封禁，unblock：不封禁，limiting：限速
	Other *UpdatePolicyRequestBodyOther `json:"other,omitempty"`

	// icmp自定义限速值，icmp取值limiting情况下，如果该值为空表示不限速
	IcmpTrafficLimiting *int64 `json:"icmp_traffic_limiting,omitempty"`

	// udp自定义限速值，udp取值limiting情况下，如果该值为空表示不限速
	UdpTrafficLimiting *int64 `json:"udp_traffic_limiting,omitempty"`

	// udp分片自定义限速值，udp取值limiting情况下，如果该值为空表示不限速
	UdpFragmentRateLimiting *int64 `json:"udp_fragment_rate_limiting,omitempty"`

	// other自定义限速值，other取值limiting情况下，如果该值为空表示不限速
	OtherTrafficLimiting *int64 `json:"other_traffic_limiting,omitempty"`

	// tcp自定义限速值，tcp取值limiting情况下，如果该值为空表示不限速
	TcpTrafficLimiting *int64 `json:"tcp_traffic_limiting,omitempty"`

	// tcp分片自定义限速值，tcp取值limiting情况下，如果该值为空表示不限速
	TcpFragmentRateLimiting *int64 `json:"tcp_fragment_rate_limiting,omitempty"`
}

func (o UpdatePolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePolicyRequestBody", string(data)}, " ")
}

type UpdatePolicyRequestBodyUdp struct {
	value string
}

type UpdatePolicyRequestBodyUdpEnum struct {
	BLOCK    UpdatePolicyRequestBodyUdp
	UNBLOCK  UpdatePolicyRequestBodyUdp
	LIMITING UpdatePolicyRequestBodyUdp
}

func GetUpdatePolicyRequestBodyUdpEnum() UpdatePolicyRequestBodyUdpEnum {
	return UpdatePolicyRequestBodyUdpEnum{
		BLOCK: UpdatePolicyRequestBodyUdp{
			value: "block",
		},
		UNBLOCK: UpdatePolicyRequestBodyUdp{
			value: "unblock",
		},
		LIMITING: UpdatePolicyRequestBodyUdp{
			value: "limiting",
		},
	}
}

func (c UpdatePolicyRequestBodyUdp) Value() string {
	return c.value
}

func (c UpdatePolicyRequestBodyUdp) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePolicyRequestBodyUdp) UnmarshalJSON(b []byte) error {
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

type UpdatePolicyRequestBodyTcp struct {
	value string
}

type UpdatePolicyRequestBodyTcpEnum struct {
	BLOCK    UpdatePolicyRequestBodyTcp
	UNBLOCK  UpdatePolicyRequestBodyTcp
	LIMITING UpdatePolicyRequestBodyTcp
}

func GetUpdatePolicyRequestBodyTcpEnum() UpdatePolicyRequestBodyTcpEnum {
	return UpdatePolicyRequestBodyTcpEnum{
		BLOCK: UpdatePolicyRequestBodyTcp{
			value: "block",
		},
		UNBLOCK: UpdatePolicyRequestBodyTcp{
			value: "unblock",
		},
		LIMITING: UpdatePolicyRequestBodyTcp{
			value: "limiting",
		},
	}
}

func (c UpdatePolicyRequestBodyTcp) Value() string {
	return c.value
}

func (c UpdatePolicyRequestBodyTcp) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePolicyRequestBodyTcp) UnmarshalJSON(b []byte) error {
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

type UpdatePolicyRequestBodyIcmp struct {
	value string
}

type UpdatePolicyRequestBodyIcmpEnum struct {
	BLOCK    UpdatePolicyRequestBodyIcmp
	UNBLOCK  UpdatePolicyRequestBodyIcmp
	LIMITING UpdatePolicyRequestBodyIcmp
}

func GetUpdatePolicyRequestBodyIcmpEnum() UpdatePolicyRequestBodyIcmpEnum {
	return UpdatePolicyRequestBodyIcmpEnum{
		BLOCK: UpdatePolicyRequestBodyIcmp{
			value: "block",
		},
		UNBLOCK: UpdatePolicyRequestBodyIcmp{
			value: "unblock",
		},
		LIMITING: UpdatePolicyRequestBodyIcmp{
			value: "limiting",
		},
	}
}

func (c UpdatePolicyRequestBodyIcmp) Value() string {
	return c.value
}

func (c UpdatePolicyRequestBodyIcmp) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePolicyRequestBodyIcmp) UnmarshalJSON(b []byte) error {
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

type UpdatePolicyRequestBodyOther struct {
	value string
}

type UpdatePolicyRequestBodyOtherEnum struct {
	BLOCK    UpdatePolicyRequestBodyOther
	UNBLOCK  UpdatePolicyRequestBodyOther
	LIMITING UpdatePolicyRequestBodyOther
}

func GetUpdatePolicyRequestBodyOtherEnum() UpdatePolicyRequestBodyOtherEnum {
	return UpdatePolicyRequestBodyOtherEnum{
		BLOCK: UpdatePolicyRequestBodyOther{
			value: "block",
		},
		UNBLOCK: UpdatePolicyRequestBodyOther{
			value: "unblock",
		},
		LIMITING: UpdatePolicyRequestBodyOther{
			value: "limiting",
		},
	}
}

func (c UpdatePolicyRequestBodyOther) Value() string {
	return c.value
}

func (c UpdatePolicyRequestBodyOther) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePolicyRequestBodyOther) UnmarshalJSON(b []byte) error {
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
