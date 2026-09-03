package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishArtifactsBody 产物发布请求体
type PublishArtifactsBody struct {

	// 工作空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 产物发布请求列表
	PublishArtifacts *[]ArtifactsPublish `json:"publish_artifacts,omitempty"`
}

func (o PublishArtifactsBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishArtifactsBody struct{}"
	}

	return strings.Join([]string{"PublishArtifactsBody", string(data)}, " ")
}
