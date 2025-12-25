package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDisasterRecoveryReq **参数解释**： 更新容灾配置请求体。 **取值范围**： 不涉及。
type UpdateDisasterRecoveryReq struct {

	// **参数解释**： 容灾同步周期。 **取值范围**： 不涉及。
	DrSyncPeriod *string `json:"dr_sync_period,omitempty"`

	// **参数解释**： 是否发送请求。 **取值范围**： 不涉及。
	SendRequest *int32 `json:"send_request,omitempty"`

	// **参数解释**： 主集群角色。 **取值范围**： 不涉及。
	PrimaryToRole *string `json:"primary_to_role,omitempty"`

	// **参数解释**： 设置容灾动作。 **取值范围**： 不涉及。
	ResetAction *string `json:"reset_action,omitempty"`

	// **参数解释**： 备集群角色。 **取值范围**： 不涉及。
	StandbyToRole *string `json:"standby_to_role,omitempty"`

	// **参数解释**： 容灾状态。 **取值范围**： - creating，容灾创建中。 - create_failed，容灾创建失败。 - unstart，容灾未启动。 - starting，容灾启动中。 - start_failed，容灾启动失败。 - running，容灾运行中。 - stopping，容灾停止中。 - stop_failed，容灾停止失败。 - switchovering，灾备切换中。 - abnormal，容灾异常。 - deleting，容灾删除中。 - deleted，容灾已删除。
	DrStatus *string `json:"dr_status,omitempty"`
}

func (o UpdateDisasterRecoveryReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDisasterRecoveryReq struct{}"
	}

	return strings.Join([]string{"UpdateDisasterRecoveryReq", string(data)}, " ")
}
