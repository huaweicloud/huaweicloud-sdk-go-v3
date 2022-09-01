package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CountTasksResponse struct {
	CommonTaskDetails *TaskStatisticDetails `json:"common_task_details,omitempty" xml:"common_task_details"`

	CdcTaskDetails *TaskStatisticDetails `json:"cdc_task_details,omitempty" xml:"cdc_task_details"`
	HttpStatusCode int                   `json:"-"`
}

func (o CountTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountTasksResponse struct{}"
	}

	return strings.Join([]string{"CountTasksResponse", string(data)}, " ")
}
