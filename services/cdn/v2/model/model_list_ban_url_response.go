package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBanUrlResponse Response Object
type ListBanUrlResponse struct {

	// 查询结果总数。
	Total *int32 `json:"total,omitempty"`

	// url信息。
	BanUrls        *[]BanUrlList `json:"ban_urls,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListBanUrlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBanUrlResponse struct{}"
	}

	return strings.Join([]string{"ListBanUrlResponse", string(data)}, " ")
}
