package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDiagnoseFeedbackResponse Response Object
type CreateDiagnoseFeedbackResponse struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDiagnoseFeedbackResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDiagnoseFeedbackResponse struct{}"
	}

	return strings.Join([]string{"CreateDiagnoseFeedbackResponse", string(data)}, " ")
}
