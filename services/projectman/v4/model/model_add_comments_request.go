package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddCommentsRequest 工作项添加评论的请求参数。
type AddCommentsRequest struct {

	// **参数解释：** 工作项id。 **约束限制：** 长度在1位到10位之间的纯数字。 **取值范围：** 最小长度：1，最大长度：10。 **默认取值：** 不涉及。
	Id string `json:"id"`

	// **参数解释：** 工作项的URL编码后的评论内容。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Notes string `json:"notes"`

	// **参数解释：** 工作项的评论内容。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	InnerText *string `json:"innerText,omitempty"`

	// **参数解释**： 项目的32位uuid。 **约束限制**： 由数字和英文组成的32字符串。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ProjectUUId *string `json:"projectUUId,omitempty"`

	// **参数解释**： 工作项所属项目类型。 **约束限制**： 不涉及。 **取值范围**： scrum。 **默认取值**： 不涉及。
	Type *string `json:"type,omitempty"`
}

func (o AddCommentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddCommentsRequest struct{}"
	}

	return strings.Join([]string{"AddCommentsRequest", string(data)}, " ")
}
