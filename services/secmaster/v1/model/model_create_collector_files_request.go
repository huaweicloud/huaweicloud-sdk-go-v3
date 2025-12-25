package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCollectorFilesRequest Request Object
type CreateCollectorFilesRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *CreateCollectorFilesRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o CreateCollectorFilesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCollectorFilesRequest struct{}"
	}

	return strings.Join([]string{"CreateCollectorFilesRequest", string(data)}, " ")
}
