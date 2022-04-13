package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type UpdateLoadbalancerRequestBody struct {
	Loadbalancer *UpdateLoadbalancerReq `json:"loadbalancer"`
}

func (o UpdateLoadbalancerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLoadbalancerRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateLoadbalancerRequestBody", string(data)}, " ")
}
