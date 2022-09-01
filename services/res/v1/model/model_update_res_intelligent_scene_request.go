package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateResIntelligentSceneRequest struct {

	// 场景id
	SceneId string `json:"scene_id" xml:"scene_id"`

	// 工作空间id
	WorkspaceId string `json:"workspace_id" xml:"workspace_id"`

	Body *UpdateResIntelligentSceneRequestBody `json:"body,omitempty" xml:"body"`
}

func (o UpdateResIntelligentSceneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResIntelligentSceneRequest struct{}"
	}

	return strings.Join([]string{"UpdateResIntelligentSceneRequest", string(data)}, " ")
}
