package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateResSceneRequest struct {

	// 场景id。
	SceneId string `json:"scene_id" xml:"scene_id"`

	// 工作空间id。
	WorkspaceId string `json:"workspace_id" xml:"workspace_id"`

	Body *UpdateResSceneRequestBody `json:"body,omitempty" xml:"body"`
}

func (o UpdateResSceneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResSceneRequest struct{}"
	}

	return strings.Join([]string{"UpdateResSceneRequest", string(data)}, " ")
}
