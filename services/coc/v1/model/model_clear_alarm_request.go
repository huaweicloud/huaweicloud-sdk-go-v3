package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClearAlarmRequest Request Object
type ClearAlarmRequest struct {
	Body *ClearAlarmRequestBody `json:"body,omitempty"`
}

func (o ClearAlarmRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClearAlarmRequest struct{}"
	}

	return strings.Join([]string{"ClearAlarmRequest", string(data)}, " ")
}
