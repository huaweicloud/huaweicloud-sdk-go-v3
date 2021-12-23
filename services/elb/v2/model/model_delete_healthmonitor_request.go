package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteHealthmonitorRequest struct {
	// 健康检查id

	HealthmonitorId string `json:"healthmonitor_id"`
}

func (o DeleteHealthmonitorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHealthmonitorRequest struct{}"
	}

	return strings.Join([]string{"DeleteHealthmonitorRequest", string(data)}, " ")
}
