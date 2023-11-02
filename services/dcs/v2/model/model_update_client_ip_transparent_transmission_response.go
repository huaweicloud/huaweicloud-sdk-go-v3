package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClientIpTransparentTransmissionResponse Response Object
type UpdateClientIpTransparentTransmissionResponse struct {

	// 开启或关闭客户ip透传的任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateClientIpTransparentTransmissionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClientIpTransparentTransmissionResponse struct{}"
	}

	return strings.Join([]string{"UpdateClientIpTransparentTransmissionResponse", string(data)}, " ")
}
