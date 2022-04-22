package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ImportGraphRequest struct {

	// 图ID。
	GraphId string `json:"graph_id"`

	// 图actionId
	ActionId ImportGraphRequestActionId `json:"action_id"`

	Body *ImportGraphReq `json:"body,omitempty"`
}

func (o ImportGraphRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportGraphRequest struct{}"
	}

	return strings.Join([]string{"ImportGraphRequest", string(data)}, " ")
}

type ImportGraphRequestActionId struct {
	value string
}

type ImportGraphRequestActionIdEnum struct {
	IMPORT_GRAPH ImportGraphRequestActionId
}

func GetImportGraphRequestActionIdEnum() ImportGraphRequestActionIdEnum {
	return ImportGraphRequestActionIdEnum{
		IMPORT_GRAPH: ImportGraphRequestActionId{
			value: "import-graph",
		},
	}
}

func (c ImportGraphRequestActionId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImportGraphRequestActionId) UnmarshalJSON(b []byte) error {
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
