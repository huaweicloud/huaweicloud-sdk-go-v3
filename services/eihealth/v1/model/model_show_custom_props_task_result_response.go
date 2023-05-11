package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowCustomPropsTaskResultResponse struct {
	Status *AsyncTaskStatus `json:"status,omitempty"`

	TaskData *CustomPropsTaskData `json:"task_data,omitempty"`

	Result         *CustomPropsResult `json:"result,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowCustomPropsTaskResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCustomPropsTaskResultResponse struct{}"
	}

	return strings.Join([]string{"ShowCustomPropsTaskResultResponse", string(data)}, " ")
}
