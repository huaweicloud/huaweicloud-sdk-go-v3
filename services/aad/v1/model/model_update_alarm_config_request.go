package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAlarmConfigRequest Request Object
type UpdateAlarmConfigRequest struct {
	Body *AlarmBody `json:"body,omitempty"`
}

func (o UpdateAlarmConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateAlarmConfigRequest", string(data)}, " ")
}
