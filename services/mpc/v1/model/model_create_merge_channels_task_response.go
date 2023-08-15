package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMergeChannelsTaskResponse Response Object
type CreateMergeChannelsTaskResponse struct {

	// 任务ID。 如果返回值为200 OK，为接受任务后产生的任务ID。
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateMergeChannelsTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMergeChannelsTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateMergeChannelsTaskResponse", string(data)}, " ")
}
