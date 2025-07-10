package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSubnetBandwidthResponse Response Object
type UpdateSubnetBandwidthResponse struct {

	// 任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateSubnetBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubnetBandwidthResponse struct{}"
	}

	return strings.Join([]string{"UpdateSubnetBandwidthResponse", string(data)}, " ")
}
