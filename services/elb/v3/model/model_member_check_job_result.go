package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MemberCheckJobResult 参数解释：后端服务器检测任务各检查项的检测结果。
type MemberCheckJobResult struct {
	Config *MemberCheckJobResultGroup `json:"config,omitempty"`

	Acl *MemberCheckJobResultGroup `json:"acl,omitempty"`

	SecurityGroup *MemberCheckJobResultGroup `json:"security_group,omitempty"`
}

func (o MemberCheckJobResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberCheckJobResult struct{}"
	}

	return strings.Join([]string{"MemberCheckJobResult", string(data)}, " ")
}
