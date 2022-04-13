package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowHealthmonitorsResponse struct {
	Healthmonitor  *HealthmonitorResp `json:"healthmonitor,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowHealthmonitorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHealthmonitorsResponse struct{}"
	}

	return strings.Join([]string{"ShowHealthmonitorsResponse", string(data)}, " ")
}
