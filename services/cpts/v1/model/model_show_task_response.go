package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowTaskResponse struct {

	// code
	Code *string `json:"code,omitempty" xml:"code"`

	// message
	Message *string `json:"message,omitempty" xml:"message"`

	Taskinfo       *TaskInfo `json:"taskinfo,omitempty" xml:"taskinfo"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskResponse", string(data)}, " ")
}
