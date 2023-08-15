package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLoadbalancerTagsResponse Response Object
type CreateLoadbalancerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateLoadbalancerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoadbalancerTagsResponse struct{}"
	}

	return strings.Join([]string{"CreateLoadbalancerTagsResponse", string(data)}, " ")
}
