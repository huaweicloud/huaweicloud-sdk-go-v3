package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssignCaseInfo 用例关联信息
type AssignCaseInfo struct {

	// 用例URI
	CaseUri *string `json:"case_uri,omitempty"`

	// 是否可用
	IsAvailable *bool `json:"is_available,omitempty"`
}

func (o AssignCaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssignCaseInfo struct{}"
	}

	return strings.Join([]string{"AssignCaseInfo", string(data)}, " ")
}
