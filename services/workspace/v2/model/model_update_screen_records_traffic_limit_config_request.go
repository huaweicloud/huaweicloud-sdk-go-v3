package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScreenRecordsTrafficLimitConfigRequest Request Object
type UpdateScreenRecordsTrafficLimitConfigRequest struct {
	Body *UpdateScreenRecordsTrafficLimitConfigRequestBody `json:"body,omitempty"`
}

func (o UpdateScreenRecordsTrafficLimitConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScreenRecordsTrafficLimitConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateScreenRecordsTrafficLimitConfigRequest", string(data)}, " ")
}
