package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateVariableRequest struct {
	// 测试工程id

	TestSuiteId int32 `json:"test_suite_id"`

	Body *[]CreateVariableRequestBody `json:"body,omitempty"`
}

func (o CreateVariableRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateVariableRequest struct{}"
	}

	return strings.Join([]string{"CreateVariableRequest", string(data)}, " ")
}
