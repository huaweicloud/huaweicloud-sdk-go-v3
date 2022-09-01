package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateResSceneRequest struct {

	// 工作空间id。
	WorkspaceId string `json:"workspace_id" xml:"workspace_id"`

	Body *CreateResSceneRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CreateResSceneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResSceneRequest struct{}"
	}

	return strings.Join([]string{"CreateResSceneRequest", string(data)}, " ")
}
