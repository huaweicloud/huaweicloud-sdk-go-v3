package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// IngressPortInfo 实例自定义入方向端口信息。
type IngressPortInfo struct {

	// 实例自定义入方向端口协议。 - HTTP：实例自定义入方向端口使用HTTP协议。 - HTTPS：实例自定义入方向端口使用HTTPS协议。
	Protocol *IngressPortInfoProtocol `json:"protocol,omitempty"`

	// 实例自定义入方向端口，支持的端口范围为1024~49151。
	IngressPort *int32 `json:"ingress_port,omitempty"`

	// 实例自定义入方向端口ID。
	IngressPortId *string `json:"ingress_port_id,omitempty"`

	// 实例自定义入方向端口的有效状态。 - normal：实例自定义入方向端口状态正常。 - abnormal：实例自定义入方向端口状态异常，无法使用。
	Status *IngressPortInfoStatus `json:"status,omitempty"`
}

func (o IngressPortInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IngressPortInfo struct{}"
	}

	return strings.Join([]string{"IngressPortInfo", string(data)}, " ")
}

type IngressPortInfoProtocol struct {
	value string
}

type IngressPortInfoProtocolEnum struct {
	HTTP  IngressPortInfoProtocol
	HTTPS IngressPortInfoProtocol
}

func GetIngressPortInfoProtocolEnum() IngressPortInfoProtocolEnum {
	return IngressPortInfoProtocolEnum{
		HTTP: IngressPortInfoProtocol{
			value: "HTTP",
		},
		HTTPS: IngressPortInfoProtocol{
			value: "HTTPS",
		},
	}
}

func (c IngressPortInfoProtocol) Value() string {
	return c.value
}

func (c IngressPortInfoProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IngressPortInfoProtocol) UnmarshalJSON(b []byte) error {
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

type IngressPortInfoStatus struct {
	value string
}

type IngressPortInfoStatusEnum struct {
	NORMAL   IngressPortInfoStatus
	ABNORMAL IngressPortInfoStatus
}

func GetIngressPortInfoStatusEnum() IngressPortInfoStatusEnum {
	return IngressPortInfoStatusEnum{
		NORMAL: IngressPortInfoStatus{
			value: "normal",
		},
		ABNORMAL: IngressPortInfoStatus{
			value: "abnormal",
		},
	}
}

func (c IngressPortInfoStatus) Value() string {
	return c.value
}

func (c IngressPortInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IngressPortInfoStatus) UnmarshalJSON(b []byte) error {
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
