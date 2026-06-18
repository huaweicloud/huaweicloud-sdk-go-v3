package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupMergeRequestValidAssignedCandidatesResponse Response Object
type ListGroupMergeRequestValidAssignedCandidatesResponse struct {
	Body *[]UserBasicDto `json:"body,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListGroupMergeRequestValidAssignedCandidatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupMergeRequestValidAssignedCandidatesResponse struct{}"
	}

	return strings.Join([]string{"ListGroupMergeRequestValidAssignedCandidatesResponse", string(data)}, " ")
}
