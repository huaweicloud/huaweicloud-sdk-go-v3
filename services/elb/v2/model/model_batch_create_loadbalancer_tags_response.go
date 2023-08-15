package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateLoadbalancerTagsResponse Response Object
type BatchCreateLoadbalancerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateLoadbalancerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateLoadbalancerTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateLoadbalancerTagsResponse", string(data)}, " ")
}
