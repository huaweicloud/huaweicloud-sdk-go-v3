package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LinksItem **参数解释：** 指向当前页或者其他页的链接。当查询需要分页时，需要包含一个next链接指向下一页。 **取值范围：** 不涉及。
type LinksItem struct {

	// **参数解释：** 对应快捷链接。 **取值范围：** 不涉及。
	Href string `json:"href"`

	// **参数解释：** 快捷链接标记名称。 **取值范围：** 不涉及。
	Rel string `json:"rel"`
}

func (o LinksItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LinksItem struct{}"
	}

	return strings.Join([]string{"LinksItem", string(data)}, " ")
}
