package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskCommentInfo Information of task comment
type TaskCommentInfo struct {

	// **参数解释**: 待办评论ID **取值范围**: 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**: 待办评论信息 **取值范围**: 不涉及
	Message string `json:"message"`

	// **参数解释**: 创建待办评论的用户ID **取值范围**: 不涉及
	UserId *string `json:"user_id,omitempty"`

	// **参数解释**: 创建待办评论的用户名称 **取值范围**: 不涉及
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**: 创建待办评论的时间 **取值范围**: 不涉及
	Time *int64 `json:"time,omitempty"`
}

func (o TaskCommentInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskCommentInfo struct{}"
	}

	return strings.Join([]string{"TaskCommentInfo", string(data)}, " ")
}
