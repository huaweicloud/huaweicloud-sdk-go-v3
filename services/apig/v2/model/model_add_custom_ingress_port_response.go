package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AddCustomIngressPortResponse Response Object
type AddCustomIngressPortResponse struct {

	// 实例自定义入方向端口协议。 - HTTP：实例自定义入方向端口使用HTTP协议。 - HTTPS：实例自定义入方向端口使用HTTPS协议。
	Protocol *AddCustomIngressPortResponseProtocol `json:"protocol,omitempty"`

	// 实例自定义入方向端口，支持的端口范围为1024~49151。
	IngressPort *int32 `json:"ingress_port,omitempty"`

	// 实例自定义入方向端口ID。
	IngressPortId *string `json:"ingress_port_id,omitempty"`

	// 实例自定义入方向端口的有效状态。 - normal：实例自定义入方向端口状态正常。 - abnormal：实例自定义入方向端口状态异常，无法使用。
	Status         *AddCustomIngressPortResponseStatus `json:"status,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o AddCustomIngressPortResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddCustomIngressPortResponse struct{}"
	}

	return strings.Join([]string{"AddCustomIngressPortResponse", string(data)}, " ")
}

type AddCustomIngressPortResponseProtocol struct {
	value string
}

type AddCustomIngressPortResponseProtocolEnum struct {
	HTTP  AddCustomIngressPortResponseProtocol
	HTTPS AddCustomIngressPortResponseProtocol
}

func GetAddCustomIngressPortResponseProtocolEnum() AddCustomIngressPortResponseProtocolEnum {
	return AddCustomIngressPortResponseProtocolEnum{
		HTTP: AddCustomIngressPortResponseProtocol{
			value: "HTTP",
		},
		HTTPS: AddCustomIngressPortResponseProtocol{
			value: "HTTPS",
		},
	}
}

func (c AddCustomIngressPortResponseProtocol) Value() string {
	return c.value
}

func (c AddCustomIngressPortResponseProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddCustomIngressPortResponseProtocol) UnmarshalJSON(b []byte) error {
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

type AddCustomIngressPortResponseStatus struct {
	value string
}

type AddCustomIngressPortResponseStatusEnum struct {
	NORMAL   AddCustomIngressPortResponseStatus
	ABNORMAL AddCustomIngressPortResponseStatus
}

func GetAddCustomIngressPortResponseStatusEnum() AddCustomIngressPortResponseStatusEnum {
	return AddCustomIngressPortResponseStatusEnum{
		NORMAL: AddCustomIngressPortResponseStatus{
			value: "normal",
		},
		ABNORMAL: AddCustomIngressPortResponseStatus{
			value: "abnormal",
		},
	}
}

func (c AddCustomIngressPortResponseStatus) Value() string {
	return c.value
}

func (c AddCustomIngressPortResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddCustomIngressPortResponseStatus) UnmarshalJSON(b []byte) error {
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
