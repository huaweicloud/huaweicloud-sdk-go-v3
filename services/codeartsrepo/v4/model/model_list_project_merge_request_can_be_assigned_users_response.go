package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectMergeRequestCanBeAssignedUsersResponse Response Object
type ListProjectMergeRequestCanBeAssignedUsersResponse struct {
	Body           *[]MergeRequestVoteReviewerDto `json:"body,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListProjectMergeRequestCanBeAssignedUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectMergeRequestCanBeAssignedUsersResponse struct{}"
	}

	return strings.Join([]string{"ListProjectMergeRequestCanBeAssignedUsersResponse", string(data)}, " ")
}
