package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMergeRequestDiscussionResponse Response Object
type DeleteMergeRequestDiscussionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteMergeRequestDiscussionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMergeRequestDiscussionResponse struct{}"
	}

	return strings.Join([]string{"DeleteMergeRequestDiscussionResponse", string(data)}, " ")
}
