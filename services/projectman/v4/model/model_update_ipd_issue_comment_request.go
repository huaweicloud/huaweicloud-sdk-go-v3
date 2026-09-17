package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIpdIssueCommentRequest Request Object
type UpdateIpdIssueCommentRequest struct {

	// 项目32位ID，项目唯一标识。通过查询IPD项目列表获取，响应消息体中的id字段的值就是项目ID。
	ProjectId string `json:"project_id"`

	// 工作项唯一ID。可以通过查询工作项列表或者查询树状工作项接口获取，响应消息体中的id字段的值就是工作项ID。
	IssueId string `json:"issue_id"`

	// **参数解释**： 评论ID。评论唯一标识。 **默认取值**： 不涉及。
	CommentId string `json:"comment_id"`

	Body *CommentUpdateVo `json:"body,omitempty"`
}

func (o UpdateIpdIssueCommentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpdIssueCommentRequest struct{}"
	}

	return strings.Join([]string{"UpdateIpdIssueCommentRequest", string(data)}, " ")
}
