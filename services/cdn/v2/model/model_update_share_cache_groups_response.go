package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateShareCacheGroupsResponse Response Object
type UpdateShareCacheGroupsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateShareCacheGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateShareCacheGroupsResponse struct{}"
	}

	return strings.Join([]string{"UpdateShareCacheGroupsResponse", string(data)}, " ")
}
