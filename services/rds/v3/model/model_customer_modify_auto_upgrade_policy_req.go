package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomerModifyAutoUpgradePolicyReq struct {

	// 自动小版本升级开关选项 true：打开自动小版本升级 false：关闭自动小版本升级
	SwitchOption bool `json:"switch_option"`
}

func (o CustomerModifyAutoUpgradePolicyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomerModifyAutoUpgradePolicyReq struct{}"
	}

	return strings.Join([]string{"CustomerModifyAutoUpgradePolicyReq", string(data)}, " ")
}
