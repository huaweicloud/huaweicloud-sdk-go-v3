package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResIntelligentSceneRequest Request Object
type CreateResIntelligentSceneRequest struct {

	// 工作空间id。
	WorkspaceId string `json:"workspace_id"`

	Body *CreateResIntelligentSceneRequestNBody `json:"body,omitempty"`
}

func (o CreateResIntelligentSceneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResIntelligentSceneRequest struct{}"
	}

	return strings.Join([]string{"CreateResIntelligentSceneRequest", string(data)}, " ")
}
