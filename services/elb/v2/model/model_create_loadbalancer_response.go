package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateLoadbalancerResponse struct {
	Loadbalancer   *LoadbalancerResp `json:"loadbalancer,omitempty" xml:"loadbalancer"`
	HttpStatusCode int               `json:"-"`
}

func (o CreateLoadbalancerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoadbalancerResponse struct{}"
	}

	return strings.Join([]string{"CreateLoadbalancerResponse", string(data)}, " ")
}
