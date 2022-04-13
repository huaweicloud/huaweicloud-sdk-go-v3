package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateHealthmonitorRequest struct {
	// 健康检查id

	HealthmonitorId string `json:"healthmonitor_id"`

	Body *UpdateHealthmonitorRequestBody `json:"body,omitempty"`
}

func (o UpdateHealthmonitorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHealthmonitorRequest struct{}"
	}

	return strings.Join([]string{"UpdateHealthmonitorRequest", string(data)}, " ")
}
