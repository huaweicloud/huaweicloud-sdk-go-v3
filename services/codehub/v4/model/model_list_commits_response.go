package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCommitsResponse Response Object
type ListCommitsResponse struct {

	// 提交列表
	Commits *[]CommitDto `json:"commits,omitempty"`

	// 按天统计信息
	DateTitle      *[]DateTitleDto `json:"date_title,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListCommitsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCommitsResponse struct{}"
	}

	return strings.Join([]string{"ListCommitsResponse", string(data)}, " ")
}
