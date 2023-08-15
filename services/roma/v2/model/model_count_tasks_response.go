package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountTasksResponse Response Object
type CountTasksResponse struct {
	CommonTaskDetails *TaskStatisticDetails `json:"common_task_details,omitempty"`

	CdcTaskDetails *TaskStatisticDetails `json:"cdc_task_details,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o CountTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountTasksResponse struct{}"
	}

	return strings.Join([]string{"CountTasksResponse", string(data)}, " ")
}
