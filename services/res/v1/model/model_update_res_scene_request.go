package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateResSceneRequest Request Object
type UpdateResSceneRequest struct {

	// 场景id。
	SceneId string `json:"scene_id"`

	// 工作空间id。
	WorkspaceId string `json:"workspace_id"`

	Body *UpdateResSceneRequestBody `json:"body,omitempty"`
}

func (o UpdateResSceneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResSceneRequest struct{}"
	}

	return strings.Join([]string{"UpdateResSceneRequest", string(data)}, " ")
}
