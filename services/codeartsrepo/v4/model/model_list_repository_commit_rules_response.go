package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRepositoryCommitRulesResponse Response Object
type ListRepositoryCommitRulesResponse struct {

	// 仓库提交规则列表
	Body *[]CommitRuleDto `json:"body,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRepositoryCommitRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryCommitRulesResponse struct{}"
	}

	return strings.Join([]string{"ListRepositoryCommitRulesResponse", string(data)}, " ")
}
