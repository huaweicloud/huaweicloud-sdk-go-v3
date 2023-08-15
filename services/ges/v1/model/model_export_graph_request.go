package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExportGraphRequest Request Object
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

func (c ExportGraphRequestActionId) Value() string {
	return c.value
}

func (c ExportGraphRequestActionId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportGraphRequestActionId) UnmarshalJSON(b []byte) error {
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
