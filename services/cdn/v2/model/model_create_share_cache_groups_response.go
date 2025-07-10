package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateShareCacheGroupsResponse Response Object
type CreateShareCacheGroupsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateShareCacheGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateShareCacheGroupsResponse struct{}"
	}

	return strings.Join([]string{"CreateShareCacheGroupsResponse", string(data)}, " ")
}
