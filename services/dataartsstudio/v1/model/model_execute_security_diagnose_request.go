package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteSecurityDiagnoseRequest Request Object
type ExecuteSecurityDiagnoseRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *DiagnoseTypeDto `json:"body,omitempty"`
}

func (o ExecuteSecurityDiagnoseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteSecurityDiagnoseRequest struct{}"
	}

	return strings.Join([]string{"ExecuteSecurityDiagnoseRequest", string(data)}, " ")
}
