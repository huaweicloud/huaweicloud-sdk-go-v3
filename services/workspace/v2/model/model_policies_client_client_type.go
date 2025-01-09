package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesClientClientType 终端类型登录管控。
type PoliciesClientClientType struct {

	// 是否开启终端类型登录管控。 false：表示关闭。 true：表示开启。
	ClientTypeLimit *bool `json:"client_type_limit,omitempty"`

	Options *PoliciesClientClientTypeOptions `json:"options,omitempty"`
}

func (o PoliciesClientClientType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesClientClientType struct{}"
	}

	return strings.Join([]string{"PoliciesClientClientType", string(data)}, " ")
}
