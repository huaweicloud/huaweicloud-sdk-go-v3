package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRelatedCommitsResponse Response Object
type ListRelatedCommitsResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *RelatedCommitListVo `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRelatedCommitsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRelatedCommitsResponse struct{}"
	}

	return strings.Join([]string{"ListRelatedCommitsResponse", string(data)}, " ")
}
