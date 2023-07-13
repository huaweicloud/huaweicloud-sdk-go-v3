package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLoadbalancerRequestBody This is a auto create Body Object
type CreateLoadbalancerRequestBody struct {
	Loadbalancer *CreateLoadbalancerReq `json:"loadbalancer"`
}

func (o CreateLoadbalancerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoadbalancerRequestBody struct{}"
	}

	return strings.Join([]string{"CreateLoadbalancerRequestBody", string(data)}, " ")
}
