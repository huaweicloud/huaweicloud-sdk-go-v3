package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteResSceneRequest struct {

	// 工作空间id。
	WorkspaceId string `json:"workspace_id" xml:"workspace_id"`

	// 场景id。
	SceneId string `json:"scene_id" xml:"scene_id"`
}

func (o DeleteResSceneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResSceneRequest struct{}"
	}

	return strings.Join([]string{"DeleteResSceneRequest", string(data)}, " ")
}
