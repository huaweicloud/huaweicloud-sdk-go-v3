package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateAlarmSubRequest struct {
	Body *AlarmSubRequest `json:"body,omitempty"`
}

func (o CreateAlarmSubRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlarmSubRequest struct{}"
	}

	return strings.Join([]string{"CreateAlarmSubRequest", string(data)}, " ")
}
