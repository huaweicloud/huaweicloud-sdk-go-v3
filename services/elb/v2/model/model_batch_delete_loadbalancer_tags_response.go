package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchDeleteLoadbalancerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteLoadbalancerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteLoadbalancerTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteLoadbalancerTagsResponse", string(data)}, " ")
}
