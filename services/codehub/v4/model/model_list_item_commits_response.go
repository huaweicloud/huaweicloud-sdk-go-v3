package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListItemCommitsResponse Response Object
type ListItemCommitsResponse struct {

	// 工作项关联的提交信息列表
	Body *[]ItemCommitDto `json:"body,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListItemCommitsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListItemCommitsResponse struct{}"
	}

	return strings.Join([]string{"ListItemCommitsResponse", string(data)}, " ")
}
