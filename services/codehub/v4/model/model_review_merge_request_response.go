package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReviewMergeRequestResponse Response Object
type ReviewMergeRequestResponse struct {
	Body           *[]ApproverBasicDto `json:"body,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ReviewMergeRequestResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReviewMergeRequestResponse struct{}"
	}

	return strings.Join([]string{"ReviewMergeRequestResponse", string(data)}, " ")
}
