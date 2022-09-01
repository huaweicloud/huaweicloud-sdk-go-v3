package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateAlarmRequest struct {
	Body *CreateAlarmRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CreateAlarmRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlarmRequest struct{}"
	}

	return strings.Join([]string{"CreateAlarmRequest", string(data)}, " ")
}
