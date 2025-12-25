package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResourceConfigRequest Request Object
type CreateResourceConfigRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *CreateCloudLogResourceRequestBody `json:"body,omitempty"`
}

func (o CreateResourceConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceConfigRequest struct{}"
	}

	return strings.Join([]string{"CreateResourceConfigRequest", string(data)}, " ")
}
