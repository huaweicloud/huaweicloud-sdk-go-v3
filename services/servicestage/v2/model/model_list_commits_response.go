package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListCommitsResponse struct {
	// 提交记录列表。

	Commits        *[]CommitsCommits `json:"commits,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListCommitsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCommitsResponse struct{}"
	}

	return strings.Join([]string{"ListCommitsResponse", string(data)}, " ")
}
