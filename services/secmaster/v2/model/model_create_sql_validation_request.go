package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSqlValidationRequest Request Object
type CreateSqlValidationRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *CreateSqlValidationRequestBody `json:"body,omitempty"`
}

func (o CreateSqlValidationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlValidationRequest struct{}"
	}

	return strings.Join([]string{"CreateSqlValidationRequest", string(data)}, " ")
}
