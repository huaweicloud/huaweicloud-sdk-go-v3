package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowPolicy 工作流策略。
type WorkflowPolicy struct {

	// 使用场景。
	UseScene *string `json:"use_scene,omitempty"`

	// 场景ID。
	SceneId *string `json:"scene_id,omitempty"`

	// 场景。
	Scenes *[]Scene `json:"scenes,omitempty"`
}

func (o WorkflowPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowPolicy struct{}"
	}

	return strings.Join([]string{"WorkflowPolicy", string(data)}, " ")
}
