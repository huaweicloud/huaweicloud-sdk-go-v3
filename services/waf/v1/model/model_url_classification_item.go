package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UrlClassificationItem struct {

	// UrlItem的总数量
	Total *int32 `json:"total,omitempty" xml:"total"`

	// UrlItem详细信息
	Items *[]UrlItem `json:"items,omitempty" xml:"items"`
}

func (o UrlClassificationItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UrlClassificationItem struct{}"
	}

	return strings.Join([]string{"UrlClassificationItem", string(data)}, " ")
}
