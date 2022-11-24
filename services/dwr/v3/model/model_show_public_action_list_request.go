package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowPublicActionListRequest struct {

	// 模板名前缀。
	Prefix *string `json:"prefix,omitempty"`

	// Action模板的分类。
	Category *string `json:"category,omitempty"`

	// 查询的起始位置
	Offset *int32 `json:"offset,omitempty"`

	// 一次查询返回的最大条数
	Limit *string `json:"limit,omitempty"`
}

func (o ShowPublicActionListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublicActionListRequest struct{}"
	}

	return strings.Join([]string{"ShowPublicActionListRequest", string(data)}, " ")
}
