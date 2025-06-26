package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDiagnosisTaskResponse Response Object
type ShowDiagnosisTaskResponse struct {

	// error_code
	ErrorCode *string `json:"error_code,omitempty"`

	// error_msg
	ErrorMsg *string `json:"error_msg,omitempty"`

	Data           *DiagnosisTaskDetail `json:"data,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowDiagnosisTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDiagnosisTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowDiagnosisTaskResponse", string(data)}, " ")
}
