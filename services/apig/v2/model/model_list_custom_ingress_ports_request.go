package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListCustomIngressPortsRequest Request Object
type ListCustomIngressPortsRequest struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty"`

	// 每页显示的条目数量，条目数量小于等于0时，自动转换为20，条目数量大于500时，自动转换为500
	Limit *int32 `json:"limit,omitempty"`

	// 入方向端口的请求协议。 - HTTP: 入方向端口为HTTP协议。 - HTTPS: 入方向端口为HTTPS协议。
	Protocol *ListCustomIngressPortsRequestProtocol `json:"protocol,omitempty"`

	// 入方向端口的端口号，支持的端口范围为1024~49151。
	IngressPort *int32 `json:"ingress_port,omitempty"`
}

func (o ListCustomIngressPortsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomIngressPortsRequest struct{}"
	}

	return strings.Join([]string{"ListCustomIngressPortsRequest", string(data)}, " ")
}

type ListCustomIngressPortsRequestProtocol struct {
	value string
}

type ListCustomIngressPortsRequestProtocolEnum struct {
	HTTP  ListCustomIngressPortsRequestProtocol
	HTTPS ListCustomIngressPortsRequestProtocol
}

func GetListCustomIngressPortsRequestProtocolEnum() ListCustomIngressPortsRequestProtocolEnum {
	return ListCustomIngressPortsRequestProtocolEnum{
		HTTP: ListCustomIngressPortsRequestProtocol{
			value: "HTTP",
		},
		HTTPS: ListCustomIngressPortsRequestProtocol{
			value: "HTTPS",
		},
	}
}

func (c ListCustomIngressPortsRequestProtocol) Value() string {
	return c.value
}

func (c ListCustomIngressPortsRequestProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCustomIngressPortsRequestProtocol) UnmarshalJSON(b []byte) error {
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
