package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRetrieveScriptRequest Request Object
type CreateRetrieveScriptRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *CreateRetrieveScriptRequestBody `json:"body,omitempty"`
}

func (o CreateRetrieveScriptRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRetrieveScriptRequest struct{}"
	}

	return strings.Join([]string{"CreateRetrieveScriptRequest", string(data)}, " ")
}
