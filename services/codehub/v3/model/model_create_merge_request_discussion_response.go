package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMergeRequestDiscussionResponse Response Object
type CreateMergeRequestDiscussionResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *MergeRequestDiscussionDto `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateMergeRequestDiscussionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMergeRequestDiscussionResponse struct{}"
	}

	return strings.Join([]string{"CreateMergeRequestDiscussionResponse", string(data)}, " ")
}
