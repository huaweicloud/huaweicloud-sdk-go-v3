package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateAlarmRuleStatusResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAlarmRuleStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmRuleStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdateAlarmRuleStatusResponse", string(data)}, " ")
}
