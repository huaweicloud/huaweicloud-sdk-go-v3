package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResSceneRequest Request Object
type ShowResSceneRequest struct {

	// 工作空间id。
	WorkspaceId string `json:"workspace_id"`

	// 场景id。
	SceneId string `json:"scene_id"`
}

func (o ShowResSceneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResSceneRequest struct{}"
	}

	return strings.Join([]string{"ShowResSceneRequest", string(data)}, " ")
}
