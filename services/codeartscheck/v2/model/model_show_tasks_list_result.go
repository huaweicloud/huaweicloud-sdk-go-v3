package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTasksListResult **参数解释：** 任务列表响应。
type ShowTasksListResult struct {

	// **参数解释：** 任务列表。
	Info *[]TaskBasicInfo `json:"info,omitempty"`

	// **参数解释：** 代码检查任务总数。 **取值范围：** 不涉及。
	Total *int32 `json:"total,omitempty"`
}

func (o ShowTasksListResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTasksListResult struct{}"
	}

	return strings.Join([]string{"ShowTasksListResult", string(data)}, " ")
}
