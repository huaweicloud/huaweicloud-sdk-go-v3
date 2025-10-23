package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskResponse Response Object
type ShowTaskResponse struct {
	TaskBasicInfo *TaskEntity `json:"task_basic_info,omitempty"`

	TaskAdditionalInfo *TaskAdditionalInfo `json:"task_additional_info,omitempty"`
	HttpStatusCode     int                 `json:"-"`
}

func (o ShowTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskResponse", string(data)}, " ")
}
