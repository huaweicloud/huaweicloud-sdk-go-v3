package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FavouriteResponseResult 结果
type FavouriteResponseResult struct {

	// 是否收藏任务
	Favorite *bool `json:"favorite,omitempty"`
}

func (o FavouriteResponseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FavouriteResponseResult struct{}"
	}

	return strings.Join([]string{"FavouriteResponseResult", string(data)}, " ")
}
