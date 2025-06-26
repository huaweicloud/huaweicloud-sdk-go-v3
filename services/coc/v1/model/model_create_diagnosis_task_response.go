package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDiagnosisTaskResponse Response Object
type CreateDiagnosisTaskResponse struct {

	// error_code
	ErrorCode *string `json:"error_code,omitempty"`

	// error_msg
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 数据
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDiagnosisTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDiagnosisTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateDiagnosisTaskResponse", string(data)}, " ")
}
