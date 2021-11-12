package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ExportGraphRequest struct {
	// 图ID。

	GraphId string `json:"graph_id"`
	// 图actionId

	ActionId ExportGraphRequestActionId `json:"action_id"`

	Body *ExportGraphReq `json:"body,omitempty"`
}

func (o ExportGraphRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportGraphRequest struct{}"
	}

	return strings.Join([]string{"ExportGraphRequest", string(data)}, " ")
}

type ExportGraphRequestActionId struct {
	value string
}

type ExportGraphRequestActionIdEnum struct {
	EXPORT_GRAPH ExportGraphRequestActionId
}

func GetExportGraphRequestActionIdEnum() ExportGraphRequestActionIdEnum {
	return ExportGraphRequestActionIdEnum{
		EXPORT_GRAPH: ExportGraphRequestActionId{
			value: "export-graph",
		},
	}
}

func (c ExportGraphRequestActionId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportGraphRequestActionId) UnmarshalJSON(b []byte) error {
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
