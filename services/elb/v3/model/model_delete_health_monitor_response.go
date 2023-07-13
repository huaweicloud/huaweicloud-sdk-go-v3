package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHealthMonitorResponse Response Object
type DeleteHealthMonitorResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteHealthMonitorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHealthMonitorResponse struct{}"
	}

	return strings.Join([]string{"DeleteHealthMonitorResponse", string(data)}, " ")
}
