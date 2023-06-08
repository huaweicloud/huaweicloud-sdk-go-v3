package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangeServerChargeModeRequestBody struct {

	// 云服务器ID列表
	ServerIds []string `json:"server_ids"`

	// 更换后的计费模式：prePaid包周期计费模式
	ChargeMode string `json:"charge_mode"`

	PrepaidOptions *ChangeServerChargeModePrepaidOption `json:"prepaid_options,omitempty"`

	// 是否预先校验此次请求。 true: 发送检查请求，不触发真正的计费转换操作 false: 发送正常请求，触发计费转换操作。 默认值为false
	DryRun *bool `json:"dry_run,omitempty"`
}

func (o ChangeServerChargeModeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeServerChargeModeRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeServerChargeModeRequestBody", string(data)}, " ")
}
