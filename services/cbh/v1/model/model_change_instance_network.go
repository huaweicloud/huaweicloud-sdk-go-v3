package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 修改网络请求body。
type ChangeInstanceNetwork struct {

	// 云堡垒机实例状态，枚举值如下： - create  创建 - renewals  续费 - change  扩容
	Type string `json:"type"`

	// 云堡垒机实例修改后的安全组信息。
	SecurityGroups []SecurityGroup `json:"security_groups"`

	// 云堡垒机实例修改后的网卡信息。
	Nics []Nics `json:"nics"`

	// 云堡垒机实例ID。云堡垒机实例状态为renewals或change时必传。
	ServerId *string `json:"server_id,omitempty"`
}

func (o ChangeInstanceNetwork) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeInstanceNetwork struct{}"
	}

	return strings.Join([]string{"ChangeInstanceNetwork", string(data)}, " ")
}
