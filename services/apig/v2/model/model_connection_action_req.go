package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ConnectionActionReq struct {

	// 允许或拒绝连接 - receive 接受 - reject 拒绝
	Action ConnectionActionReqAction `json:"action"`

	// 终端节点列表
	Endpoints []string `json:"endpoints"`
}

func (o ConnectionActionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionActionReq struct{}"
	}

	return strings.Join([]string{"ConnectionActionReq", string(data)}, " ")
}

type ConnectionActionReqAction struct {
	value string
}

type ConnectionActionReqActionEnum struct {
	RECEIVE ConnectionActionReqAction
	REJECT  ConnectionActionReqAction
}

func GetConnectionActionReqActionEnum() ConnectionActionReqActionEnum {
	return ConnectionActionReqActionEnum{
		RECEIVE: ConnectionActionReqAction{
			value: "receive",
		},
		REJECT: ConnectionActionReqAction{
			value: "reject",
		},
	}
}

func (c ConnectionActionReqAction) Value() string {
	return c.value
}

func (c ConnectionActionReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConnectionActionReqAction) UnmarshalJSON(b []byte) error {
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
