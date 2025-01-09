package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSubnetBandwidthResponse Response Object
type DeleteSubnetBandwidthResponse struct {

	// 任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteSubnetBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubnetBandwidthResponse struct{}"
	}

	return strings.Join([]string{"DeleteSubnetBandwidthResponse", string(data)}, " ")
}
