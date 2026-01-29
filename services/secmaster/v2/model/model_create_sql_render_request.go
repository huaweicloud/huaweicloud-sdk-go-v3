package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSqlRenderRequest Request Object
type CreateSqlRenderRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *CreateSqlRenderRequestBody `json:"body,omitempty"`
}

func (o CreateSqlRenderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlRenderRequest struct{}"
	}

	return strings.Join([]string{"CreateSqlRenderRequest", string(data)}, " ")
}
