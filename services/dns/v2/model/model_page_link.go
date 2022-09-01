package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 分页信息
type PageLink struct {

	// 当前资源的链接。
	Self *string `json:"self,omitempty" xml:"self"`

	// 下一页资源的链接。
	Next *string `json:"next,omitempty" xml:"next"`
}

func (o PageLink) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageLink struct{}"
	}

	return strings.Join([]string{"PageLink", string(data)}, " ")
}
