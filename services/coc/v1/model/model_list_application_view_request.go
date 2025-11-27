package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApplicationViewRequest Request Object
type ListApplicationViewRequest struct {

	// **参数解释：** 名称模糊匹配，支持模糊查询。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	NameLike *string `json:"name_like,omitempty"`

	// **参数解释：** 应用、组件、分组code组成。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	CodeList *[]string `json:"code_list,omitempty"`

	// **参数解释：** 分页参数，上一页请求最后一个id。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Marker *string `json:"marker,omitempty"`

	// **参数解释：** 分页查询每页显示的条目数量。 **约束限制：** 不涉及。 **取值范围：** 自定义，在1-500范围。 **默认取值：** 不涉及。
	Limit int32 `json:"limit"`

	// **参数解释：** 分页页码。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	PageNo *int32 `json:"page_no,omitempty"`

	// **参数解释：** 是否收藏。 **约束限制：** 不涉及。 **取值范围：** - true：在我的收藏去查询应用、组件、分组，默认为true。 - false：在全部应用中查询应用、组件、分组，可以不传。 **默认取值：** 默认未收藏。
	IsCollection *bool `json:"is_collection,omitempty"`
}

func (o ListApplicationViewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationViewRequest struct{}"
	}

	return strings.Join([]string{"ListApplicationViewRequest", string(data)}, " ")
}
