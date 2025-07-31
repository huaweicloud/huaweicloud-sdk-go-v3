package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PageLink **参数解释：** 指向当前页或者其他页的链接。当查询需要分页时，需要包含一个next链接指向下一页。 **取值范围：** 不涉及。
type PageLink struct {

	// **参数解释：** 当前页面的链接。 **取值范围：** 不涉及。
	Self *string `json:"self,omitempty"`

	// **参数解释：** 下一页的链接。 **取值范围：** 不涉及。
	Next *string `json:"next,omitempty"`
}

func (o PageLink) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageLink struct{}"
	}

	return strings.Join([]string{"PageLink", string(data)}, " ")
}
