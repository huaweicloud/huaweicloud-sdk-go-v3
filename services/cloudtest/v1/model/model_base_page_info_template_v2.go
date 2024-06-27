package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BasePageInfoTemplateV2 struct {

	// 每页显示的条目数量，最大支持200条
	Limit *int32 `json:"limit,omitempty"`

	List *[]TemplateV2 `json:"list,omitempty"`

	// 起始偏移量，表示从此偏移量开始查询，offset大于等于0，小于等于100000
	Offset *int32 `json:"offset,omitempty"`

	Pages *int32 `json:"pages,omitempty"`

	Size *int32 `json:"size,omitempty"`

	Total *int64 `json:"total,omitempty"`
}

func (o BasePageInfoTemplateV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BasePageInfoTemplateV2 struct{}"
	}

	return strings.Join([]string{"BasePageInfoTemplateV2", string(data)}, " ")
}
