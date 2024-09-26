package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MergeRequestApprovalReviewersDto struct {
	ApprovalMergeRequestReviewers *[]ApprovalUserDto `json:"approval_merge_request_reviewers,omitempty"`
}

func (o MergeRequestApprovalReviewersDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MergeRequestApprovalReviewersDto struct{}"
	}

	return strings.Join([]string{"MergeRequestApprovalReviewersDto", string(data)}, " ")
}
