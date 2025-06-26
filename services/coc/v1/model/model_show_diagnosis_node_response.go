package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDiagnosisNodeResponse Response Object
type ShowDiagnosisNodeResponse struct {

	// error_code
	ErrorCode *string `json:"error_code,omitempty"`

	// error_msg
	ErrorMsg *string `json:"error_msg,omitempty"`

	Data           *DiagnosisTaskNodeDetail `json:"data,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowDiagnosisNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDiagnosisNodeResponse struct{}"
	}

	return strings.Join([]string{"ShowDiagnosisNodeResponse", string(data)}, " ")
}
