package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueResponseV4Env 缺陷发现环境（仅Bug类型工作项具备该字段）
type IssueResponseV4Env struct {

	// 缺陷发现环境id
	Id *int32 `json:"id,omitempty"`

	// 缺陷发现环境名称
	Name *string `json:"name,omitempty"`
}

func (o IssueResponseV4Env) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueResponseV4Env struct{}"
	}

	return strings.Join([]string{"IssueResponseV4Env", string(data)}, " ")
}
