package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplySubnetBandwidthResponse Response Object
type ApplySubnetBandwidthResponse struct {

	// 任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ApplySubnetBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplySubnetBandwidthResponse struct{}"
	}

	return strings.Join([]string{"ApplySubnetBandwidthResponse", string(data)}, " ")
}
