package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteP2cVpnGatewayJobRequest Request Object
type DeleteP2cVpnGatewayJobRequest struct {

	// 任务ID
	JobId string `json:"job_id"`
}

func (o DeleteP2cVpnGatewayJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteP2cVpnGatewayJobRequest struct{}"
	}

	return strings.Join([]string{"DeleteP2cVpnGatewayJobRequest", string(data)}, " ")
}
