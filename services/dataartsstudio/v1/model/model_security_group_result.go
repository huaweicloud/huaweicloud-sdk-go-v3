package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityGroupResult 安全组诊断结果
type SecurityGroupResult struct {
	Result *DiagnoseResult `json:"result,omitempty"`

	// kerberos信息
	SecurityGroup *[]SecurityGroupStatus `json:"security_group,omitempty"`
}

func (o SecurityGroupResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityGroupResult struct{}"
	}

	return strings.Join([]string{"SecurityGroupResult", string(data)}, " ")
}
