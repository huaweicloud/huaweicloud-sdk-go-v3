package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowLoadbalancersStatusResponse struct {
	Statuses       *StatusResp `json:"statuses,omitempty" xml:"statuses"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowLoadbalancersStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLoadbalancersStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowLoadbalancersStatusResponse", string(data)}, " ")
}
