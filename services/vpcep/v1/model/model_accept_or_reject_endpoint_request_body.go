package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AcceptOrRejectEndpointRequestBody 连接终端节点列表请求结构体
type AcceptOrRejectEndpointRequestBody struct {

	// 允许或拒绝连接。  - receive：允许连接。  - reject：拒绝连接。
	Action AcceptOrRejectEndpointRequestBodyAction `json:"action"`

	// 终端节点ID列表。 每次请求目前支持单条endpoint的接受或拒绝。
	Endpoints []string `json:"endpoints"`
}

func (o AcceptOrRejectEndpointRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AcceptOrRejectEndpointRequestBody struct{}"
	}

	return strings.Join([]string{"AcceptOrRejectEndpointRequestBody", string(data)}, " ")
}

type AcceptOrRejectEndpointRequestBodyAction struct {
	value string
}

type AcceptOrRejectEndpointRequestBodyActionEnum struct {
	RECEIVE AcceptOrRejectEndpointRequestBodyAction
	REJECT  AcceptOrRejectEndpointRequestBodyAction
}

func GetAcceptOrRejectEndpointRequestBodyActionEnum() AcceptOrRejectEndpointRequestBodyActionEnum {
	return AcceptOrRejectEndpointRequestBodyActionEnum{
		RECEIVE: AcceptOrRejectEndpointRequestBodyAction{
			value: "receive",
		},
		REJECT: AcceptOrRejectEndpointRequestBodyAction{
			value: "reject",
		},
	}
}

func (c AcceptOrRejectEndpointRequestBodyAction) Value() string {
	return c.value
}

func (c AcceptOrRejectEndpointRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AcceptOrRejectEndpointRequestBodyAction) UnmarshalJSON(b []byte) error {
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
