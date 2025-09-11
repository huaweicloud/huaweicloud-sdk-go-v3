package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IsFavoriteItem struct {

	// **参数解释** 监控看板是否标记收藏 **取值范围** - true: 收藏, - false: 未收藏
	IsFavorite *bool `json:"is_favorite,omitempty"`
}

func (o IsFavoriteItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsFavoriteItem struct{}"
	}

	return strings.Join([]string{"IsFavoriteItem", string(data)}, " ")
}
