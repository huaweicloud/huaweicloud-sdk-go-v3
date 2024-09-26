package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMergeRequestReviewersResponse Response Object
type ListMergeRequestReviewersResponse struct {
	Result *MergeRequestApprovalReviewersDto `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListMergeRequestReviewersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMergeRequestReviewersResponse struct{}"
	}

	return strings.Join([]string{"ListMergeRequestReviewersResponse", string(data)}, " ")
}
