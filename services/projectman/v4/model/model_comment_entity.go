package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CommentEntity 评论实体对象
type CommentEntity struct {

	// **参数解释**： 评论ID。 **默认取值**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 评论类型。 **取值范围**： - comment：评论 - reply：回复 - operation：系统操作。 **默认取值**： 不涉及。
	Category *string `json:"category,omitempty"`

	// **参数解释**： 评论元数据类型。 **取值范围**： 固定为comment。 **默认取值**： 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 是否显示在置顶区域。 **取值范围**： - true：显示。 - false： 不显示。 **默认取值**： 不涉及。
	Top *bool `json:"top,omitempty"`

	// **参数解释**： 置顶时间的unix时间戳，单位：毫秒。当有多条置顶评论时，最后置顶的评论显示在最上层。 **默认取值**： 不涉及。
	TopTime *string `json:"top_time,omitempty"`

	// **参数解释**： 评论内容，表现形式为html标签。 **默认取值**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 评论关联的工作项ID。 **默认取值**： 不涉及。
	IssueId *string `json:"issue_id,omitempty"`

	// **参数解释**： 当前评论是否被置顶。 **取值范围**： - true：置顶。 - false： 不置顶。 **默认取值**： 不涉及。
	TopFlag *bool `json:"top_flag,omitempty"`

	// **参数解释**： 评论创建人ID。 **默认取值**： 不涉及。
	CreatedBy *string `json:"created_by,omitempty"`

	// **参数解释**： 评论创建时间。 **默认取值**： 不涉及。
	CreatedDate *string `json:"created_date,omitempty"`

	CreatorInfo *UserVo `json:"creator_info,omitempty"`

	// **参数解释**： 评论的一些扩展属性，表现为json字符串。 **默认取值**： 不涉及。
	ExtendAttribute *string `json:"extend_attribute,omitempty"`

	ExtendAttributeObj *CommentExtendAttribute `json:"extend_attribute_obj,omitempty"`

	// **参数解释**： 评论的扩展属性对象数组。 **默认取值**： 不涉及。
	ExtendAttributeObjs *[]CommentExtendAttribute `json:"extend_attribute_objs,omitempty"`
}

func (o CommentEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommentEntity struct{}"
	}

	return strings.Join([]string{"CommentEntity", string(data)}, " ")
}
