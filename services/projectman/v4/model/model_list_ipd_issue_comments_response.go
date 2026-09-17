package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIpdIssueCommentsResponse Response Object
type ListIpdIssueCommentsResponse struct {

	// **参数解释**： 响应状态。 **取值范围**： 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 响应信息。 **取值范围**： 不涉及。
	Message *string `json:"message,omitempty"`

	Result         *CommentResult `json:"result,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListIpdIssueCommentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIpdIssueCommentsResponse struct{}"
	}

	return strings.Join([]string{"ListIpdIssueCommentsResponse", string(data)}, " ")
}
