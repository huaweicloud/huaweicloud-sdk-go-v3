package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommentUpdateVo struct {

	// **参数解释**： 评论内容，使用html标记语言。 **默认取值**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 评论时@他人的用户ID，填写此参数后会通知被@的用户，通知形式在需求管理-设置-工作项设置-通知设置中配置。 **默认取值**： 不涉及。
	At *string `json:"at,omitempty"`
}

func (o CommentUpdateVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommentUpdateVo struct{}"
	}

	return strings.Join([]string{"CommentUpdateVo", string(data)}, " ")
}
