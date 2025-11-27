package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountResourcesOfResourceViewRequest Request Object
type CountResourcesOfResourceViewRequest struct {

	// **参数解释：** 视图ID。 **约束限制：** 不涉及。 **取值范围：** 自定义生成 长度1~32之间。 **默认取值：** 不涉及。
	ViewId string `json:"view_id"`

	// **参数解释：** 云服务名称。 **约束限制：** 不涉及。 **取值范围：** 自定义，可选esc，cce等资源。 **默认取值：** 不涉及。
	Provider *string `json:"provider,omitempty"`

	// **参数解释：** 资源类型名称。 **约束限制：** 不涉及。 **取值范围：** 自定义，云资源类型。 **默认取值：** 不涉及。
	Type *string `json:"type,omitempty"`
}

func (o CountResourcesOfResourceViewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountResourcesOfResourceViewRequest struct{}"
	}

	return strings.Join([]string{"CountResourcesOfResourceViewRequest", string(data)}, " ")
}
