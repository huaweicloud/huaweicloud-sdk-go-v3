package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCommentsRequest 工作项更新评论的请求参数。
type UpdateCommentsRequest struct {

	// **参数解释：** 工作项id，可通过[高级查询工作项](ListIssuesV4.xml)接口获取，响应消息体中的**id**字段的值就是工作项id。 **约束限制：** 长度在1位到10位之间的纯数字。 **取值范围：** 最小长度：1，最大长度：10。 **默认取值：** 不涉及。
	Id int32 `json:"id"`

	// **参数解释：** 工作项的评论URL编码内容。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Notes string `json:"notes"`

	// **参数解释：** 工作项的评论内容。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	InnerText *string `json:"innerText,omitempty"`

	// **参数解释**： 项目的32位uuid，项目唯一标识。 **约束限制**： 由数字和英文组成的32字符串。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ProjectUUId string `json:"projectUUId"`

	// **参数解释**： 工作项所属项目类型。 **约束限制**： 不涉及。 **取值范围**： scrum。 **默认取值**： 不涉及。
	Type string `json:"type"`

	// **参数解释：** 工作项的评论id。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	NoteId int32 `json:"noteId"`
}

func (o UpdateCommentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCommentsRequest struct{}"
	}

	return strings.Join([]string{"UpdateCommentsRequest", string(data)}, " ")
}
