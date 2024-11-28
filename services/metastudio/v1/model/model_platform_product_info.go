package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PlatformProductInfo 第三方直播平台商品详情
type PlatformProductInfo struct {

	// 商品id
	ProductId *string `json:"product_id,omitempty"`

	// 商品标题
	ProductTitle *string `json:"product_title,omitempty"`
}

func (o PlatformProductInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlatformProductInfo struct{}"
	}

	return strings.Join([]string{"PlatformProductInfo", string(data)}, " ")
}
