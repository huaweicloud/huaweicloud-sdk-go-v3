package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWhitelistsResponse Response Object
type ListWhitelistsResponse struct {

	// 总条数
	Total *int32 `json:"total,omitempty"`

	// 白名单列表
	Whitelists     *[]WhitelistBean `json:"whitelists,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListWhitelistsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWhitelistsResponse struct{}"
	}

	return strings.Join([]string{"ListWhitelistsResponse", string(data)}, " ")
}
