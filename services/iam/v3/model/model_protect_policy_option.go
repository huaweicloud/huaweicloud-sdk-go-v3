package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type ProtectPolicyOption struct {
	// 是否开启操作保护，开启为\"true\"，未开启为\"false\"。

	OperationProtection bool `json:"operation_protection"`
}

func (o ProtectPolicyOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectPolicyOption struct{}"
	}

	return strings.Join([]string{"ProtectPolicyOption", string(data)}, " ")
}
