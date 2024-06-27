package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CaseInfo 更新用例信息列表
type CaseInfo struct {

	// 用例id
	CaseId *string `json:"case_id,omitempty"`

	// 脚本路径
	ScriptUrl *string `json:"script_url,omitempty"`
}

func (o CaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CaseInfo struct{}"
	}

	return strings.Join([]string{"CaseInfo", string(data)}, " ")
}
