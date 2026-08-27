package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUpgradeTasksResponse Response Object
type ListUpgradeTasksResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 任务列表
	Tasks          *[]UpgradeScheduledTaskVo `json:"tasks,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListUpgradeTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUpgradeTasksResponse struct{}"
	}

	return strings.Join([]string{"ListUpgradeTasksResponse", string(data)}, " ")
}
