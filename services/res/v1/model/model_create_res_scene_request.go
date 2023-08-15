package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResSceneRequest Request Object
type CreateResSceneRequest struct {

	// 工作空间id。
	WorkspaceId string `json:"workspace_id"`

	Body *CreateResSceneRequestBody `json:"body,omitempty"`
}

func (o CreateResSceneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResSceneRequest struct{}"
	}

	return strings.Join([]string{"CreateResSceneRequest", string(data)}, " ")
}
