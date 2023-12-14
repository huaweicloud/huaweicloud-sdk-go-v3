package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMonitorSystemResponse Response Object
type UpdateMonitorSystemResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateMonitorSystemResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMonitorSystemResponse struct{}"
	}

	return strings.Join([]string{"UpdateMonitorSystemResponse", string(data)}, " ")
}
