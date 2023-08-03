package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueDetailResponseV4Env 缺陷发现环境（仅Bug类型工作项具备该字段）
type IssueDetailResponseV4Env struct {

	// 缺陷发现环境id
	Id *int32 `json:"id,omitempty"`

	// 缺陷发现环境名称
	Name *string `json:"name,omitempty"`
}

func (o IssueDetailResponseV4Env) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueDetailResponseV4Env struct{}"
	}

	return strings.Join([]string{"IssueDetailResponseV4Env", string(data)}, " ")
}
