package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLoadbalancerResponse Response Object
type ShowLoadbalancerResponse struct {
	Loadbalancer   *LoadbalancerResp `json:"loadbalancer,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowLoadbalancerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLoadbalancerResponse struct{}"
	}

	return strings.Join([]string{"ShowLoadbalancerResponse", string(data)}, " ")
}
