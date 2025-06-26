package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDiagnosisTasksResponse Response Object
type ListDiagnosisTasksResponse struct {

	// error_code
	ErrorCode *string `json:"error_code,omitempty"`

	// error_msg
	ErrorMsg *string `json:"error_msg,omitempty"`

	Data           *DiagnosisTaskPage `json:"data,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListDiagnosisTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDiagnosisTasksResponse struct{}"
	}

	return strings.Join([]string{"ListDiagnosisTasksResponse", string(data)}, " ")
}
