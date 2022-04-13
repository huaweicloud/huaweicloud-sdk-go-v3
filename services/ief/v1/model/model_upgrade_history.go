package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 节点安装或升级记录
type UpgradeHistory struct {
	// 节点升级或安装历史版本id

	Id *int32 `json:"id,omitempty"`
	// 节点升级前节点上edgecore版本号，形式如2.1.0，其中每一位都是整数

	FromVersion *string `json:"from_version,omitempty"`
	// 节点升级或安装后节点行edgecore版本号，形式如2.1.0，其中每一位都是整数

	ToVersion *string `json:"to_version,omitempty"`
	// 节点升级或安装的十位时间戳

	UpgradeTime *int32 `json:"upgrade_time,omitempty"`
	// 节点升级或安装状态，包含 - install_success：边缘节点安装成功 - upgrade_success：边缘节点升级成功 - install_failed：边缘节点安装失败 - upgrade_failed：边缘节点升级失败 - upgrade_failed_rollback_success：边缘节点升级失败回滚成功 - upgrade_failed_rollback_failed：边缘节点升级失败回滚失败

	Result *string `json:"result,omitempty"`
	// 节点升级所消耗的时间

	DurTime *int32 `json:"dur_time,omitempty"`
}

func (o UpgradeHistory) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeHistory struct{}"
	}

	return strings.Join([]string{"UpgradeHistory", string(data)}, " ")
}
