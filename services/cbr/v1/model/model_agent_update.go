package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

//
type AgentUpdate struct {

	// 客户端状态，当前只支持卸载，由客户端被卸载时自动触发
	Status AgentUpdateStatus `json:"status"`
}

func (o AgentUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentUpdate struct{}"
	}

	return strings.Join([]string{"AgentUpdate", string(data)}, " ")
}

type AgentUpdateStatus struct {
	value string
}

type AgentUpdateStatusEnum struct {
	UNINSTALL AgentUpdateStatus
}

func GetAgentUpdateStatusEnum() AgentUpdateStatusEnum {
	return AgentUpdateStatusEnum{
		UNINSTALL: AgentUpdateStatus{
			value: "uninstall",
		},
	}
}

func (c AgentUpdateStatus) Value() string {
	return c.value
}

func (c AgentUpdateStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AgentUpdateStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
