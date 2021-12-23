package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowTaskDefectsRequest struct {
	// 任务ID

	TaskId string `json:"task_id"`
	// 分页索引，偏移量

	Offset *int32 `json:"offset,omitempty"`
	// 每页显示的数量,每页最多显示100条

	Limit *int32 `json:"limit,omitempty"`
	// 问题状态筛选

	StatusIds *string `json:"status_ids,omitempty"`
	// 严重级别，0致命，1严重，2一般，3提示

	Severity *string `json:"severity,omitempty"`
}

func (o ShowTaskDefectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskDefectsRequest struct{}"
	}

	return strings.Join([]string{"ShowTaskDefectsRequest", string(data)}, " ")
}
