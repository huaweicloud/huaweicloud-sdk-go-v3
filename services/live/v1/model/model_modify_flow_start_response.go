package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ModifyFlowStartResponse Response Object
type ModifyFlowStartResponse struct {

	// 状态，ACTIVE状态为运行状态,STANDBY状态为未启动状态
	Status *ModifyFlowStartResponseStatus `json:"status,omitempty"`

	// 流id
	FlowId         *string `json:"flow_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ModifyFlowStartResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyFlowStartResponse struct{}"
	}

	return strings.Join([]string{"ModifyFlowStartResponse", string(data)}, " ")
}

type ModifyFlowStartResponseStatus struct {
	value string
}

type ModifyFlowStartResponseStatusEnum struct {
	ACTIVE  ModifyFlowStartResponseStatus
	STANDBY ModifyFlowStartResponseStatus
}

func GetModifyFlowStartResponseStatusEnum() ModifyFlowStartResponseStatusEnum {
	return ModifyFlowStartResponseStatusEnum{
		ACTIVE: ModifyFlowStartResponseStatus{
			value: "ACTIVE",
		},
		STANDBY: ModifyFlowStartResponseStatus{
			value: "STANDBY",
		},
	}
}

func (c ModifyFlowStartResponseStatus) Value() string {
	return c.value
}

func (c ModifyFlowStartResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyFlowStartResponseStatus) UnmarshalJSON(b []byte) error {
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
