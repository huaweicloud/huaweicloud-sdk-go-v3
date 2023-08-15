package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSearchTaskResultResponse Response Object
type ShowSearchTaskResultResponse struct {
	Status *AsyncTaskStatus `json:"status,omitempty"`

	TaskData *SearchTaskData `json:"task_data,omitempty"`

	Result         *SearchResult `json:"result,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowSearchTaskResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSearchTaskResultResponse struct{}"
	}

	return strings.Join([]string{"ShowSearchTaskResultResponse", string(data)}, " ")
}
