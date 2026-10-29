package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskListByProjectIdResponse Response Object
type ShowTaskListByProjectIdResponse struct {

	// **参数解释：** 任务信息。
	Tasks *[]SimpleTaskInfoV2 `json:"tasks,omitempty"`

	// **参数解释：** 项目下的任务总数。 **取值范围：** 不涉及。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowTaskListByProjectIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskListByProjectIdResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskListByProjectIdResponse", string(data)}, " ")
}
