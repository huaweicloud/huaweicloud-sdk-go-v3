package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskResponse Response Object
type ShowTaskResponse struct {

	// code
	Code *string `json:"code,omitempty"`

	// message
	Message *string `json:"message,omitempty"`

	TaskInfo       *TaskInfo `json:"taskInfo,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskResponse", string(data)}, " ")
}
