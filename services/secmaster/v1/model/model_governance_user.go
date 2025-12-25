package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GovernanceUser 资产治理责任人
type GovernanceUser struct {

	// 资产治理责任人类型
	Type *string `json:"type,omitempty"`

	// 资产治理责任人名称，为空则忽略，不存在则创建
	Name *string `json:"name,omitempty"`
}

func (o GovernanceUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GovernanceUser struct{}"
	}

	return strings.Join([]string{"GovernanceUser", string(data)}, " ")
}
