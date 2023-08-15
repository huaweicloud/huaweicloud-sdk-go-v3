package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCpiTaskResultResponse Response Object
type ShowCpiTaskResultResponse struct {
	Status *AsyncTaskStatus `json:"status,omitempty"`

	TaskData *CpiTaskData `json:"task_data,omitempty"`

	Result         *CpiResult `json:"result,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowCpiTaskResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCpiTaskResultResponse struct{}"
	}

	return strings.Join([]string{"ShowCpiTaskResultResponse", string(data)}, " ")
}
