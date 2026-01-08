package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeLoadbalancerTagsResponse Response Object
type ChangeLoadbalancerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeLoadbalancerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeLoadbalancerTagsResponse struct{}"
	}

	return strings.Join([]string{"ChangeLoadbalancerTagsResponse", string(data)}, " ")
}
