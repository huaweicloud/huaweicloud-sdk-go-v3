package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ClearGraphRequest struct {
	// 图ID。

	GraphId string `json:"graph_id"`
	// 图actionId

	ActionId ClearGraphRequestActionId `json:"action_id"`
	// 是否清空图关联的元数据。建议清除。

	ClearMetadata *bool `json:"clear-metadata,omitempty"`
}

func (o ClearGraphRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClearGraphRequest struct{}"
	}

	return strings.Join([]string{"ClearGraphRequest", string(data)}, " ")
}

type ClearGraphRequestActionId struct {
	value string
}

type ClearGraphRequestActionIdEnum struct {
	CLEAR_GRAPH ClearGraphRequestActionId
}

func GetClearGraphRequestActionIdEnum() ClearGraphRequestActionIdEnum {
	return ClearGraphRequestActionIdEnum{
		CLEAR_GRAPH: ClearGraphRequestActionId{
			value: "clear-graph",
		},
	}
}

func (c ClearGraphRequestActionId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ClearGraphRequestActionId) UnmarshalJSON(b []byte) error {
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
