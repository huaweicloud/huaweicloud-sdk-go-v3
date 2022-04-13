package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteLoadbalancerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLoadbalancerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLoadbalancerResponse struct{}"
	}

	return strings.Join([]string{"DeleteLoadbalancerResponse", string(data)}, " ")
}
