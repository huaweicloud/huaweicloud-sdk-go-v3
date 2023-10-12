package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteOneClickAlarmsRequest Request Object
type BatchDeleteOneClickAlarmsRequest struct {
	Body *BatchDeleteOneClickAlarmsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteOneClickAlarmsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteOneClickAlarmsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteOneClickAlarmsRequest", string(data)}, " ")
}
