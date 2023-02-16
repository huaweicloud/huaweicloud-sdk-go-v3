package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 搜索的相关信息。
type SearchRestInfo struct {

	// 搜索结果列表。
	Items *[]SearchItem `json:"items,omitempty"`

	SearchInfo *SearchInfo `json:"search_info,omitempty"`

	ImageInfo *SearchRestInfoImageInfo `json:"image_info,omitempty"`
}

func (o SearchRestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchRestInfo struct{}"
	}

	return strings.Join([]string{"SearchRestInfo", string(data)}, " ")
}
