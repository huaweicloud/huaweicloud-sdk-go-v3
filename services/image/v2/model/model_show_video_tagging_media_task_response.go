package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVideoTaggingMediaTaskResponse Response Object
type ShowVideoTaggingMediaTaskResponse struct {

	// 任务id
	TaskId *string `json:"task_id,omitempty"`

	// 任务创建时间，格式为ISO8601：YYYY-MM-DDThh:mm:ssZ
	CreateTime *string `json:"create_time,omitempty"`

	// 任务更新时间，格式为ISO8601：YYYY-MM-DDThh:mm:ssZ
	UpdateTime *string `json:"update_time,omitempty"`

	// 任务当前的状态，分别为SUCCEEDED（运行成功），FAILED（运行失败），RUNNING（运行中）。
	State *string `json:"state,omitempty"`

	Input *VideoTaggingMediaTaskInput `json:"input,omitempty"`

	Output *TaskOutput `json:"output,omitempty"`

	Config *VideoTaggingTaskConfig `json:"config,omitempty"`

	Callback *TaskCallback `json:"callback,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowVideoTaggingMediaTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVideoTaggingMediaTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowVideoTaggingMediaTaskResponse", string(data)}, " ")
}
