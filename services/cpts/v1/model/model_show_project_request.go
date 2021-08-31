package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowProjectRequest struct {
	// 测试工程id

	TestSuiteId int32 `json:"test_suite_id"`
}

func (o ShowProjectRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowProjectRequest struct{}"
	}

	return strings.Join([]string{"ShowProjectRequest", string(data)}, " ")
}
