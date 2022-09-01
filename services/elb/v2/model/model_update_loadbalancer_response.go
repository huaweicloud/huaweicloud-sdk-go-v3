package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateLoadbalancerResponse struct {
	Loadbalancer   *LoadbalancerResp `json:"loadbalancer,omitempty" xml:"loadbalancer"`
	HttpStatusCode int               `json:"-"`
}

func (o UpdateLoadbalancerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLoadbalancerResponse struct{}"
	}

	return strings.Join([]string{"UpdateLoadbalancerResponse", string(data)}, " ")
}
