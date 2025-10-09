package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVpnGatewayJobRequest Request Object
type DeleteVpnGatewayJobRequest struct {

	// 任务ID
	JobId string `json:"job_id"`
}

func (o DeleteVpnGatewayJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVpnGatewayJobRequest struct{}"
	}

	return strings.Join([]string{"DeleteVpnGatewayJobRequest", string(data)}, " ")
}
