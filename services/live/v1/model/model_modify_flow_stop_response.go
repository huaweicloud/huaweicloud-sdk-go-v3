package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ModifyFlowStopResponse Response Object
type ModifyFlowStopResponse struct {

	// 状态，ACTIVE状态为运行状态,STANDBY状态为未启动状态
	Status *ModifyFlowStopResponseStatus `json:"status,omitempty"`

	// 流id
	FlowId         *string `json:"flow_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ModifyFlowStopResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyFlowStopResponse struct{}"
	}

	return strings.Join([]string{"ModifyFlowStopResponse", string(data)}, " ")
}

type ModifyFlowStopResponseStatus struct {
	value string
}

type ModifyFlowStopResponseStatusEnum struct {
	ACTIVE  ModifyFlowStopResponseStatus
	STANDBY ModifyFlowStopResponseStatus
}

func GetModifyFlowStopResponseStatusEnum() ModifyFlowStopResponseStatusEnum {
	return ModifyFlowStopResponseStatusEnum{
		ACTIVE: ModifyFlowStopResponseStatus{
			value: "ACTIVE",
		},
		STANDBY: ModifyFlowStopResponseStatus{
			value: "STANDBY",
		},
	}
}

func (c ModifyFlowStopResponseStatus) Value() string {
	return c.value
}

func (c ModifyFlowStopResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyFlowStopResponseStatus) UnmarshalJSON(b []byte) error {
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
