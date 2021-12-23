package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteLoadbalancerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLoadbalancerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLoadbalancerTagsResponse struct{}"
	}

	return strings.Join([]string{"DeleteLoadbalancerTagsResponse", string(data)}, " ")
}
