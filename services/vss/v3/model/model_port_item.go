package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PortItem struct {

	// 端口号
	Port *int32 `json:"port,omitempty"`

	// 服务
	Service *string `json:"service,omitempty"`

	// 端口协议:   * TCP   * UDP
	Protocol *PortItemProtocol `json:"protocol,omitempty"`

	// 端口状态:   * filtered - 过滤的   * open - 开放
	Status *PortItemStatus `json:"status,omitempty"`
}

func (o PortItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PortItem struct{}"
	}

	return strings.Join([]string{"PortItem", string(data)}, " ")
}

type PortItemProtocol struct {
	value string
}

type PortItemProtocolEnum struct {
	TCP PortItemProtocol
	UDP PortItemProtocol
}

func GetPortItemProtocolEnum() PortItemProtocolEnum {
	return PortItemProtocolEnum{
		TCP: PortItemProtocol{
			value: "TCP",
		},
		UDP: PortItemProtocol{
			value: "UDP",
		},
	}
}

func (c PortItemProtocol) Value() string {
	return c.value
}

func (c PortItemProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PortItemProtocol) UnmarshalJSON(b []byte) error {
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

type PortItemStatus struct {
	value string
}

type PortItemStatusEnum struct {
	FILTERED PortItemStatus
	OPEN     PortItemStatus
}

func GetPortItemStatusEnum() PortItemStatusEnum {
	return PortItemStatusEnum{
		FILTERED: PortItemStatus{
			value: "filtered",
		},
		OPEN: PortItemStatus{
			value: "open",
		},
	}
}

func (c PortItemStatus) Value() string {
	return c.value
}

func (c PortItemStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PortItemStatus) UnmarshalJSON(b []byte) error {
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
