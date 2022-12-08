package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 修改网络请求body
type ChangeInstanceNetwork struct {

	// 状态
	Type string `json:"type"`

	// 安全组信息
	SecurityGroups []SecurityGroup `json:"security_groups"`

	// 网卡信息
	Nics []Nics `json:"nics"`
}

func (o ChangeInstanceNetwork) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeInstanceNetwork struct{}"
	}

	return strings.Join([]string{"ChangeInstanceNetwork", string(data)}, " ")
}
