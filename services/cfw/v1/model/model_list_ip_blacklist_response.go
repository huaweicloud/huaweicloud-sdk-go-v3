package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIpBlacklistResponse Response Object
type ListIpBlacklistResponse struct {
	Data           *PageDataIpBlacklistsVo `json:"data,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListIpBlacklistResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIpBlacklistResponse struct{}"
	}

	return strings.Join([]string{"ListIpBlacklistResponse", string(data)}, " ")
}
