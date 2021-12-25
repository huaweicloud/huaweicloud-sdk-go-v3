package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListCommitsResponse struct {
	Error *Error `json:"error,omitempty"`
	// 提交列表

	Result *[]CommitInfo `json:"result,omitempty"`
	// 响应状态

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListCommitsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCommitsResponse struct{}"
	}

	return strings.Join([]string{"ListCommitsResponse", string(data)}, " ")
}
