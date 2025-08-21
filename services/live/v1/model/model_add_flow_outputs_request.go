package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AddFlowOutputsRequest struct {

	// ip白名单，最大20条ip白名单
	CidrWhitelist *[]string `json:"cidr_whitelist,omitempty"`

	// 推流地址，支持ip和域名，最大值64
	Destination *string `json:"destination,omitempty"`

	// output名称，有效字符：大小写字母，数字，中划线，下划线;且只能以字母开头
	Name string `json:"name"`

	Encryption *FlowSourceDecryption `json:"encryption,omitempty"`

	// 输出状态，ENABLED：启用，DISABLED：禁用
	OutputStatus *AddFlowOutputsRequestOutputStatus `json:"output_status,omitempty"`

	// 端口，如果ssrt-listener为监听端口，那么srt-caller为对端端口
	Port *int32 `json:"port,omitempty"`

	// 协议，srt-caller，srt-listener
	Protocol AddFlowOutputsRequestProtocol `json:"protocol"`

	// srt-caller模式支持设置streamid
	StreamId *string `json:"stream_id,omitempty"`

	// 输出描述
	Description *string `json:"description,omitempty"`

	// 最小时延，单位毫秒，默认值2000，最小值10，最大值15000
	MinLatency *int32 `json:"min_latency,omitempty"`
}

func (o AddFlowOutputsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddFlowOutputsRequest struct{}"
	}

	return strings.Join([]string{"AddFlowOutputsRequest", string(data)}, " ")
}

type AddFlowOutputsRequestOutputStatus struct {
	value string
}

type AddFlowOutputsRequestOutputStatusEnum struct {
	ENABLED  AddFlowOutputsRequestOutputStatus
	DISABLED AddFlowOutputsRequestOutputStatus
}

func GetAddFlowOutputsRequestOutputStatusEnum() AddFlowOutputsRequestOutputStatusEnum {
	return AddFlowOutputsRequestOutputStatusEnum{
		ENABLED: AddFlowOutputsRequestOutputStatus{
			value: "ENABLED",
		},
		DISABLED: AddFlowOutputsRequestOutputStatus{
			value: "DISABLED",
		},
	}
}

func (c AddFlowOutputsRequestOutputStatus) Value() string {
	return c.value
}

func (c AddFlowOutputsRequestOutputStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddFlowOutputsRequestOutputStatus) UnmarshalJSON(b []byte) error {
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

type AddFlowOutputsRequestProtocol struct {
	value string
}

type AddFlowOutputsRequestProtocolEnum struct {
	SRT_LISTENER AddFlowOutputsRequestProtocol
	SRT_CALLER   AddFlowOutputsRequestProtocol
}

func GetAddFlowOutputsRequestProtocolEnum() AddFlowOutputsRequestProtocolEnum {
	return AddFlowOutputsRequestProtocolEnum{
		SRT_LISTENER: AddFlowOutputsRequestProtocol{
			value: "srt-listener",
		},
		SRT_CALLER: AddFlowOutputsRequestProtocol{
			value: "srt-caller",
		},
	}
}

func (c AddFlowOutputsRequestProtocol) Value() string {
	return c.value
}

func (c AddFlowOutputsRequestProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddFlowOutputsRequestProtocol) UnmarshalJSON(b []byte) error {
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
