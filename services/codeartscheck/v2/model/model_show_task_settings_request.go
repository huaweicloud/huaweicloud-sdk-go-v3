package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskSettingsRequest Request Object
type ShowTaskSettingsRequest struct {

	// 项目ID
	ProjectId string `json:"project_id"`

	// 任务ID
	TaskId string `json:"task_id"`

	// 分页索引，偏移量，非必填
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的数量，非必填
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowTaskSettingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskSettingsRequest struct{}"
	}

	return strings.Join([]string{"ShowTaskSettingsRequest", string(data)}, " ")
}
