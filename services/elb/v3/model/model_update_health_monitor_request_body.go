package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHealthMonitorRequestBody This is a auto create Body Object
type UpdateHealthMonitorRequestBody struct {
	Healthmonitor *UpdateHealthMonitorOption `json:"healthmonitor"`
}

func (o UpdateHealthMonitorRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHealthMonitorRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateHealthMonitorRequestBody", string(data)}, " ")
}
