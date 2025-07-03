package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListFlowRespItem 响应流列表项
type ListFlowRespItem struct {

	// 流名称
	Name string `json:"name"`

	// 区域
	Region string `json:"region"`

	// 流id
	FlowId string `json:"flow_id"`

	// 当前状态，ACTIVE运行中，STANDBY未启动
	Status ListFlowRespItemStatus `json:"status"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 协议，srt-caller，srt-listener
	Protocol ListFlowRespItemProtocol `json:"protocol"`
}

func (o ListFlowRespItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlowRespItem struct{}"
	}

	return strings.Join([]string{"ListFlowRespItem", string(data)}, " ")
}

type ListFlowRespItemStatus struct {
	value string
}

type ListFlowRespItemStatusEnum struct {
	ACTIVE  ListFlowRespItemStatus
	STANDBY ListFlowRespItemStatus
}

func GetListFlowRespItemStatusEnum() ListFlowRespItemStatusEnum {
	return ListFlowRespItemStatusEnum{
		ACTIVE: ListFlowRespItemStatus{
			value: "ACTIVE",
		},
		STANDBY: ListFlowRespItemStatus{
			value: "STANDBY",
		},
	}
}

func (c ListFlowRespItemStatus) Value() string {
	return c.value
}

func (c ListFlowRespItemStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlowRespItemStatus) UnmarshalJSON(b []byte) error {
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

type ListFlowRespItemProtocol struct {
	value string
}

type ListFlowRespItemProtocolEnum struct {
	SRT_CALLER   ListFlowRespItemProtocol
	SRT_LISTENER ListFlowRespItemProtocol
}

func GetListFlowRespItemProtocolEnum() ListFlowRespItemProtocolEnum {
	return ListFlowRespItemProtocolEnum{
		SRT_CALLER: ListFlowRespItemProtocol{
			value: "srt-caller",
		},
		SRT_LISTENER: ListFlowRespItemProtocol{
			value: "srt-listener",
		},
	}
}

func (c ListFlowRespItemProtocol) Value() string {
	return c.value
}

func (c ListFlowRespItemProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlowRespItemProtocol) UnmarshalJSON(b []byte) error {
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
