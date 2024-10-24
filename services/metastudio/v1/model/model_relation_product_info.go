package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RelationProductInfo 关联商品
type RelationProductInfo struct {

	// 关联商品ID。如果配置，则段落切换回调中会携带该信息。美团平台对应goodsId
	ProductId *string `json:"product_id,omitempty"`

	// 关联商品标题/名称。如果配置，则段落切换回调中会携带该信息。美团平台对应goodsTitle
	ProductTitle *string `json:"product_title,omitempty"`
}

func (o RelationProductInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RelationProductInfo struct{}"
	}

	return strings.Join([]string{"RelationProductInfo", string(data)}, " ")
}
