package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateResIntelligentSceneRequest Request Object
type UpdateResIntelligentSceneRequest struct {

	// 场景id
	SceneId string `json:"scene_id"`

	// 工作空间id
	WorkspaceId string `json:"workspace_id"`

	Body *UpdateResIntelligentSceneRequestBody `json:"body,omitempty"`
}

func (o UpdateResIntelligentSceneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResIntelligentSceneRequest struct{}"
	}

	return strings.Join([]string{"UpdateResIntelligentSceneRequest", string(data)}, " ")
}
