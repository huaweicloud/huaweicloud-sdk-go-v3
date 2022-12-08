package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CancelScheduleTaskRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	Body *CancelScheduleTask `json:"body,omitempty"`
}

func (o CancelScheduleTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelScheduleTaskRequest struct{}"
	}

	return strings.Join([]string{"CancelScheduleTaskRequest", string(data)}, " ")
}
