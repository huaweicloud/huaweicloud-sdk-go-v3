package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApprovalMergeRequestResponse Response Object
type ApprovalMergeRequestResponse struct {
	Body           *[]ApproverBasicDto `json:"body,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ApprovalMergeRequestResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApprovalMergeRequestResponse struct{}"
	}

	return strings.Join([]string{"ApprovalMergeRequestResponse", string(data)}, " ")
}
