package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDisasterRecoveryReq 更新容灾配置请求体
type UpdateDisasterRecoveryReq struct {

	// 容灾同步周期
	DrSyncPeriod *string `json:"dr_sync_period,omitempty"`

	// 是否发送请求
	SendRequest *int32 `json:"send_request,omitempty"`

	// 主集群角色
	PrimaryToRole *string `json:"primary_to_role,omitempty"`

	// 设置容灾动作
	ResetAction *string `json:"reset_action,omitempty"`

	// 备集群角色
	StandbyToRole *string `json:"standby_to_role,omitempty"`

	// 容灾状态
	DrStatus *string `json:"dr_status,omitempty"`
}

func (o UpdateDisasterRecoveryReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDisasterRecoveryReq struct{}"
	}

	return strings.Join([]string{"UpdateDisasterRecoveryReq", string(data)}, " ")
}
