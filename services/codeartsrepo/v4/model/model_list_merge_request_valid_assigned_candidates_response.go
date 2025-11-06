package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMergeRequestValidAssignedCandidatesResponse Response Object
type ListMergeRequestValidAssignedCandidatesResponse struct {
	Body           *[]MrVoteReviewerDto `json:"body,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListMergeRequestValidAssignedCandidatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMergeRequestValidAssignedCandidatesResponse struct{}"
	}

	return strings.Join([]string{"ListMergeRequestValidAssignedCandidatesResponse", string(data)}, " ")
}
