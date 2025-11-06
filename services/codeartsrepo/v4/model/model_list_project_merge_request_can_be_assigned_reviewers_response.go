package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectMergeRequestCanBeAssignedReviewersResponse Response Object
type ListProjectMergeRequestCanBeAssignedReviewersResponse struct {
	Body           *[]MergeRequestVoteReviewerDto `json:"body,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListProjectMergeRequestCanBeAssignedReviewersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectMergeRequestCanBeAssignedReviewersResponse struct{}"
	}

	return strings.Join([]string{"ListProjectMergeRequestCanBeAssignedReviewersResponse", string(data)}, " ")
}
