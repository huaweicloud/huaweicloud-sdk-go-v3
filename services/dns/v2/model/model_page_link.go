package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 分页信息
type PageLink struct {
	// 当前资源的链接。

	Self *string `json:"self,omitempty"`
	// 下一页资源的链接。

	Next *string `json:"next,omitempty"`
}

func (o PageLink) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageLink struct{}"
	}

	return strings.Join([]string{"PageLink", string(data)}, " ")
}
