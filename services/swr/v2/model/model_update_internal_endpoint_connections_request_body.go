package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateInternalEndpointConnectionsRequestBody struct {

	// VPC终端节点ID列表
	Endpoints []string `json:"endpoints"`

	// 允许或拒绝连接 取值范围: - receive:允许连接 - reject:拒绝连接
	Action UpdateInternalEndpointConnectionsRequestBodyAction `json:"action"`
}

func (o UpdateInternalEndpointConnectionsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInternalEndpointConnectionsRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateInternalEndpointConnectionsRequestBody", string(data)}, " ")
}

type UpdateInternalEndpointConnectionsRequestBodyAction struct {
	value string
}

type UpdateInternalEndpointConnectionsRequestBodyActionEnum struct {
	RECEIVE UpdateInternalEndpointConnectionsRequestBodyAction
	REJECT  UpdateInternalEndpointConnectionsRequestBodyAction
}

func GetUpdateInternalEndpointConnectionsRequestBodyActionEnum() UpdateInternalEndpointConnectionsRequestBodyActionEnum {
	return UpdateInternalEndpointConnectionsRequestBodyActionEnum{
		RECEIVE: UpdateInternalEndpointConnectionsRequestBodyAction{
			value: "receive",
		},
		REJECT: UpdateInternalEndpointConnectionsRequestBodyAction{
			value: "reject",
		},
	}
}

func (c UpdateInternalEndpointConnectionsRequestBodyAction) Value() string {
	return c.value
}

func (c UpdateInternalEndpointConnectionsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateInternalEndpointConnectionsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
