package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIpdIssueCommentResponse Response Object
type UpdateIpdIssueCommentResponse struct {

	// **参数解释**： 返回状态。 **取值范围**： - success：响应成功 - error：响应失败
	Status *string `json:"status,omitempty"`

	// **参数解释**： 信息。 **取值范围**： 不涉及
	Message *string `json:"message,omitempty"`

	Result         *CommentEntity `json:"result,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o UpdateIpdIssueCommentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpdIssueCommentResponse struct{}"
	}

	return strings.Join([]string{"UpdateIpdIssueCommentResponse", string(data)}, " ")
}
