package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ReplaceNodeRequest struct {

	// 只读实例ID。
	InstanceId string `json:"instance_id"`

	// 只读实例的节点ID。
	NodeId string `json:"node_id"`

	// 替换动作，取值范围： REPLACE: 节点顶替 REPLACE_ROLLBACK: 节点顶替回切
	ReplaceAction ReplaceNodeRequestReplaceAction `json:"replace_action"`
}

func (o ReplaceNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReplaceNodeRequest struct{}"
	}

	return strings.Join([]string{"ReplaceNodeRequest", string(data)}, " ")
}

type ReplaceNodeRequestReplaceAction struct {
	value string
}

type ReplaceNodeRequestReplaceActionEnum struct {
	REPLACE          ReplaceNodeRequestReplaceAction
	REPLACE_ROLLBACK ReplaceNodeRequestReplaceAction
}

func GetReplaceNodeRequestReplaceActionEnum() ReplaceNodeRequestReplaceActionEnum {
	return ReplaceNodeRequestReplaceActionEnum{
		REPLACE: ReplaceNodeRequestReplaceAction{
			value: "REPLACE",
		},
		REPLACE_ROLLBACK: ReplaceNodeRequestReplaceAction{
			value: "REPLACE_ROLLBACK",
		},
	}
}

func (c ReplaceNodeRequestReplaceAction) Value() string {
	return c.value
}

func (c ReplaceNodeRequestReplaceAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReplaceNodeRequestReplaceAction) UnmarshalJSON(b []byte) error {
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
