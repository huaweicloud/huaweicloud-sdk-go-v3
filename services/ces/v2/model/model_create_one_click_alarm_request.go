package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOneClickAlarmRequest Request Object
type CreateOneClickAlarmRequest struct {
	Body *EnableOneClickAlarmRequestBody `json:"body,omitempty"`
}

func (o CreateOneClickAlarmRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOneClickAlarmRequest struct{}"
	}

	return strings.Join([]string{"CreateOneClickAlarmRequest", string(data)}, " ")
}
