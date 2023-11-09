package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteConnectionMonitorResponse Response Object
type DeleteConnectionMonitorResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteConnectionMonitorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConnectionMonitorResponse struct{}"
	}

	return strings.Join([]string{"DeleteConnectionMonitorResponse", string(data)}, " ")
}
