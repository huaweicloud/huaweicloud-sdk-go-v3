package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAdhocQueryRequest Request Object
type CreateAdhocQueryRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *CreateAdhocQueryRequestBody `json:"body,omitempty"`
}

func (o CreateAdhocQueryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAdhocQueryRequest struct{}"
	}

	return strings.Join([]string{"CreateAdhocQueryRequest", string(data)}, " ")
}
