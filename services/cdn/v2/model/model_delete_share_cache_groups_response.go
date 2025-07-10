package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteShareCacheGroupsResponse Response Object
type DeleteShareCacheGroupsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteShareCacheGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteShareCacheGroupsResponse struct{}"
	}

	return strings.Join([]string{"DeleteShareCacheGroupsResponse", string(data)}, " ")
}
