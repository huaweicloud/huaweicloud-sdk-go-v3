package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMergeRequestReviewersResponse Response Object
type UpdateMergeRequestReviewersResponse struct {
	Body           *[]ApproverBasicDto `json:"body,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o UpdateMergeRequestReviewersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMergeRequestReviewersResponse struct{}"
	}

	return strings.Join([]string{"UpdateMergeRequestReviewersResponse", string(data)}, " ")
}
