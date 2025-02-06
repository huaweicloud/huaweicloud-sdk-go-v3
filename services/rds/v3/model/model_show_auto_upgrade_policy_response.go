package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoUpgradePolicyResponse Response Object
type ShowAutoUpgradePolicyResponse struct {

	// 自动小版本升级开关选项 true：打开自动小版本升级 false：关闭自动小版本升级
	SwitchOption   *bool `json:"switch_option,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowAutoUpgradePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoUpgradePolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowAutoUpgradePolicyResponse", string(data)}, " ")
}
