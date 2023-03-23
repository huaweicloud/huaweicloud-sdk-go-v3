package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowGenerationTaskResultResponse struct {
	Status *AsyncTaskStatus `json:"status,omitempty"`

	TaskData *GenerationTaskData `json:"task_data,omitempty"`

	Result         *GenerationResult `json:"result,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowGenerationTaskResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGenerationTaskResultResponse struct{}"
	}

	return strings.Join([]string{"ShowGenerationTaskResultResponse", string(data)}, " ")
}
