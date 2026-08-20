package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryIssueAssociatedItemResponse Response Object
type QueryIssueAssociatedItemResponse struct {

	// 返回状态。
	Status *string `json:"status,omitempty"`

	// 信息。
	Message *string `json:"message,omitempty"`

	Result         *IssueListResult `json:"result,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o QueryIssueAssociatedItemResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryIssueAssociatedItemResponse struct{}"
	}

	return strings.Join([]string{"QueryIssueAssociatedItemResponse", string(data)}, " ")
}
