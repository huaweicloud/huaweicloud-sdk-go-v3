package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIpsWhitelistsResponse Response Object
type ListIpsWhitelistsResponse struct {
	Data           *ListIpsWhiteListsVo `json:"data,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListIpsWhitelistsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIpsWhitelistsResponse struct{}"
	}

	return strings.Join([]string{"ListIpsWhitelistsResponse", string(data)}, " ")
}
