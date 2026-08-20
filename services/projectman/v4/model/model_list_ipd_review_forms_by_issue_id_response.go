package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIpdReviewFormsByIssueIdResponse Response Object
type ListIpdReviewFormsByIssueIdResponse struct {

	// 返回状态。
	Status *string `json:"status,omitempty"`

	// 请求信息。
	Message *string `json:"message,omitempty"`

	Result         *ShowIpdProcessInstancesResponseResult `json:"result,omitempty"`
	HttpStatusCode int                                    `json:"-"`
}

func (o ListIpdReviewFormsByIssueIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIpdReviewFormsByIssueIdResponse struct{}"
	}

	return strings.Join([]string{"ListIpdReviewFormsByIssueIdResponse", string(data)}, " ")
}
