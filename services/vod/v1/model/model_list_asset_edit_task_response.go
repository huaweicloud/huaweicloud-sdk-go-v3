package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAssetEditTaskResponse Response Object
type ListAssetEditTaskResponse struct {

	// 任务总数
	Total *int64 `json:"total,omitempty"`

	// 任务列表
	Tasks          *[]EditingTaskInfo `json:"tasks,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListAssetEditTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssetEditTaskResponse struct{}"
	}

	return strings.Join([]string{"ListAssetEditTaskResponse", string(data)}, " ")
}
