package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteResSceneRequest Request Object
type DeleteResSceneRequest struct {

	// 工作空间id。
	WorkspaceId string `json:"workspace_id"`

	// 场景id。
	SceneId string `json:"scene_id"`
}

func (o DeleteResSceneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResSceneRequest struct{}"
	}

	return strings.Join([]string{"DeleteResSceneRequest", string(data)}, " ")
}
