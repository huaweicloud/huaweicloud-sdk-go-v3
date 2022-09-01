package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 热度策略的分群配置
type Category struct {

	// 用户特征。
	UserMetaList *[]string `json:"user_meta_list,omitempty" xml:"user_meta_list"`

	// 物品特征。
	ItemMetaList *[]string `json:"item_meta_list,omitempty" xml:"item_meta_list"`
}

func (o Category) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Category struct{}"
	}

	return strings.Join([]string{"Category", string(data)}, " ")
}
