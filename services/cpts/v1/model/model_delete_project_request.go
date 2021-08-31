package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteProjectRequest struct {
	// 测试工程id

	TestSuiteId int32 `json:"test_suite_id"`
}

func (o DeleteProjectRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteProjectRequest struct{}"
	}

	return strings.Join([]string{"DeleteProjectRequest", string(data)}, " ")
}
