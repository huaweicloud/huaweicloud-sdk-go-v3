package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AgentSwitchRequest struct {

	// 审计agent的ID。可在查询数据库agent列表接口ID字段获取。
	AgentId string `json:"agent_id"`

	// Agent开关状态 - 1：开启 - 0：关闭
	Status AgentSwitchRequestStatus `json:"status"`
}

func (o AgentSwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentSwitchRequest struct{}"
	}

	return strings.Join([]string{"AgentSwitchRequest", string(data)}, " ")
}

type AgentSwitchRequestStatus struct {
	value int32
}

type AgentSwitchRequestStatusEnum struct {
	E_0 AgentSwitchRequestStatus
	E_1 AgentSwitchRequestStatus
}

func GetAgentSwitchRequestStatusEnum() AgentSwitchRequestStatusEnum {
	return AgentSwitchRequestStatusEnum{
		E_0: AgentSwitchRequestStatus{
			value: 0,
		}, E_1: AgentSwitchRequestStatus{
			value: 1,
		},
	}
}

func (c AgentSwitchRequestStatus) Value() int32 {
	return c.value
}

func (c AgentSwitchRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AgentSwitchRequestStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
