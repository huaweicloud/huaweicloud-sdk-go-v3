package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateResIntelligentSceneRequest struct {

	// 工作空间id。
	WorkspaceId string `json:"workspace_id" xml:"workspace_id"`

	Body *CreateResIntelligentSceneRequestNBody `json:"body,omitempty" xml:"body"`
}

func (o CreateResIntelligentSceneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResIntelligentSceneRequest struct{}"
	}

	return strings.Join([]string{"CreateResIntelligentSceneRequest", string(data)}, " ")
}
