package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProductTextInfo 商品文本信息
type ProductTextInfo struct {

	// 文本标题
	Title *string `json:"title,omitempty"`

	// 文本
	Text *string `json:"text,omitempty"`
}

func (o ProductTextInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductTextInfo struct{}"
	}

	return strings.Join([]string{"ProductTextInfo", string(data)}, " ")
}
