package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScreenRecordsTrafficLimitConfigResponse Response Object
type UpdateScreenRecordsTrafficLimitConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateScreenRecordsTrafficLimitConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScreenRecordsTrafficLimitConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateScreenRecordsTrafficLimitConfigResponse", string(data)}, " ")
}
