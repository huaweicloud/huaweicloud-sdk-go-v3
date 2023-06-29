package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletefavoriteRequest Request Object
type DeletefavoriteRequest struct {

	// 收藏资源id
	FavResId string `json:"fav_res_id"`
}

func (o DeletefavoriteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletefavoriteRequest struct{}"
	}

	return strings.Join([]string{"DeletefavoriteRequest", string(data)}, " ")
}
