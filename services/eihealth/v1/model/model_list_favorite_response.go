package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFavoriteResponse Response Object
type ListFavoriteResponse struct {

	// 收藏列表。
	Favorites *[]FavoriteDto `json:"favorites,omitempty"`

	// 收藏总数。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListFavoriteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFavoriteResponse struct{}"
	}

	return strings.Join([]string{"ListFavoriteResponse", string(data)}, " ")
}
