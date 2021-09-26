package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateTestCaseInPlanResponse struct {
	// 接口调用失败错误码

	ErrorCode *string `json:"error_code,omitempty"`
	// 接口调用失败错误信息

	ErrorMsg *string `json:"error_msg,omitempty"`

	ErrorDetail    *ErrorDetailInfo `json:"error_detail,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CreateTestCaseInPlanResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateTestCaseInPlanResponse struct{}"
	}

	return strings.Join([]string{"CreateTestCaseInPlanResponse", string(data)}, " ")
}