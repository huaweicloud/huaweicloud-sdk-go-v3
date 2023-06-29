package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHealthmonitorResponse Response Object
type UpdateHealthmonitorResponse struct {
	Healthmonitor  *HealthmonitorResp `json:"healthmonitor,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o UpdateHealthmonitorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHealthmonitorResponse struct{}"
	}

	return strings.Join([]string{"UpdateHealthmonitorResponse", string(data)}, " ")
}
