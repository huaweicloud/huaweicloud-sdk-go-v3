package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceSpecsPricePageInfo struct {

	// 下一页起始标识，为null代表无下一页
	NextMarker *string `json:"next_marker,omitempty"`
}

func (o ResourceSpecsPricePageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceSpecsPricePageInfo struct{}"
	}

	return strings.Join([]string{"ResourceSpecsPricePageInfo", string(data)}, " ")
}
