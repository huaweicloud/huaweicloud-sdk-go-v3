package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CommentCreateVo struct {

	// **参数解释**： 评论类型。 **取值范围**： - comment：评论 - reply：回复 - operation：系统操作（不支持创建）。 **默认取值**： 不涉及。
	Category *CommentCreateVoCategory `json:"category,omitempty"`

	// **参数解释**： 评论关联的工作项类型。 **默认取值**： 不涉及。
	IssueCategory *string `json:"issue_category,omitempty"`

	// **参数解释**： 评论内容，使用html标记语言。 **默认取值**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 评论的父ID，取值为需要回复的评论的ID。 **约束限制**： 回复评论时必填。 **取值范围**： 只支持CR。 **默认取值**： 不涉及。
	ParentId *string `json:"parent_id,omitempty"`

	// **参数解释**： 评论的根ID，取值为需要回复的首层评论的ID。 **约束限制**： 回复评论时必填，创建评论时不能填。 **默认取值**： 不涉及。
	RootId *string `json:"root_id,omitempty"`

	// **参数解释**： 评论时@他人的用户ID，填写此参数后会通知被@的用户，通知形式在需求管理-设置-工作项设置-通知设置中配置。 **默认取值**： 不涉及。
	At *string `json:"at,omitempty"`
}

func (o CommentCreateVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommentCreateVo struct{}"
	}

	return strings.Join([]string{"CommentCreateVo", string(data)}, " ")
}

type CommentCreateVoCategory struct {
	value string
}

type CommentCreateVoCategoryEnum struct {
	COMMENT CommentCreateVoCategory
	REPLY   CommentCreateVoCategory
}

func GetCommentCreateVoCategoryEnum() CommentCreateVoCategoryEnum {
	return CommentCreateVoCategoryEnum{
		COMMENT: CommentCreateVoCategory{
			value: "comment",
		},
		REPLY: CommentCreateVoCategory{
			value: "reply",
		},
	}
}

func (c CommentCreateVoCategory) Value() string {
	return c.value
}

func (c CommentCreateVoCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CommentCreateVoCategory) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
