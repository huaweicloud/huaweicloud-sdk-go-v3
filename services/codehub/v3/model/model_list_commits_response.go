package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListCommitsResponse struct {
	Error *Error `json:"error,omitempty" xml:"error"`

	// 提交列表
	Result *[]CommitInfo `json:"result,omitempty" xml:"result"`

	// 响应状态
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o ListCommitsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCommitsResponse struct{}"
	}

	return strings.Join([]string{"ListCommitsResponse", string(data)}, " ")
}
