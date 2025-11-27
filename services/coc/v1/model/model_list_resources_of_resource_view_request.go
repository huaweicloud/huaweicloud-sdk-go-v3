package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourcesOfResourceViewRequest Request Object
type ListResourcesOfResourceViewRequest struct {

	// **参数解释：** 用于分页查询，查询数量，最大的返回数量。 **约束限制：** 不涉及。 **取值范围：** 自定义，在1-500范围。 **默认取值：** 不涉及。
	Limit int32 `json:"limit"`

	// **参数解释：** 用于分页查询。 **约束限制：** 不涉及。 **取值范围：** 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页。 **默认取值：** 不涉及。
	Offset *string `json:"offset,omitempty"`

	// **参数解释：** 视图ID。 **约束限制：** 不涉及。 **取值范围：** 自定义生成 长度1~32之间。 **默认取值：** 不涉及。
	ViewId string `json:"view_id"`

	// **参数解释：** 云服务名称。 **约束限制：** 不涉及。 **取值范围：** 自定义，可选esc，cce等资源。 **默认取值：** 不涉及。
	Provider *string `json:"provider,omitempty"`

	// **参数解释：** 资源类型名称。 **约束限制：** 不涉及。 **取值范围：** 自定义，云资源类型。 **默认取值：** 不涉及。
	Type *string `json:"type,omitempty"`
}

func (o ListResourcesOfResourceViewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcesOfResourceViewRequest struct{}"
	}

	return strings.Join([]string{"ListResourcesOfResourceViewRequest", string(data)}, " ")
}
