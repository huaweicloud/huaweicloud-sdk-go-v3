package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NetworkRequestBody 检查云堡垒机实例网络请求信息。
type NetworkRequestBody struct {

	// 云堡垒机实例状态，枚举值如下： - create  创建 - renewals  更新 - change  变更 状态为renewals或change时server_id必传。
	Type string `json:"type"`

	// 云堡垒升级实例所在安全组信息。
	SecurityGroups []SecurityGroup `json:"security_groups"`

	// 云堡垒机实例的网卡信息。
	Nics []Nics `json:"nics"`

	// 云堡垒机实例ID。
	ServerId *string `json:"server_id,omitempty"`
}

func (o NetworkRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkRequestBody struct{}"
	}

	return strings.Join([]string{"NetworkRequestBody", string(data)}, " ")
}
