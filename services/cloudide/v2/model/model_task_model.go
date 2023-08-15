package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskModel struct {

	// the docker_id
	DockerId *string `json:"docker_id,omitempty"`

	// exception
	Exception *string `json:"exception,omitempty"`

	// the generated_snippet
	GeneratedSnippet *string `json:"generated_snippet,omitempty"`

	// code language
	Language *string `json:"language,omitempty"`

	// model_id
	ModelId *string `json:"model_id,omitempty"`

	// record_time
	RecordTime *sdktime.SdkTime `json:"record_time,omitempty"`

	// the unique id of request
	RequestId *string `json:"request_id,omitempty"`

	// the snippet of code
	Snippet *string `json:"snippet,omitempty"`

	// start_time
	StartTime *sdktime.SdkTime `json:"start_time,omitempty"`

	// status
	Status *string `json:"status,omitempty"`

	// task_id
	TaskId *int32 `json:"task_id,omitempty"`

	// the time_consuming
	TimeConsuming float32 `json:"time_consuming,omitempty"`
}

func (o TaskModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskModel struct{}"
	}

	return strings.Join([]string{"TaskModel", string(data)}, " ")
}
