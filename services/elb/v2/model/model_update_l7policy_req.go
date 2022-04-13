package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新转发策略对象请求体
type UpdateL7policyReq struct {
	// 转发策略名称

	Name *string `json:"name,omitempty"`
	// 转发策略的管理状态；该字段为预留字段，暂未启用。默认为true。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 转发策略额描述信息

	Description *string `json:"description,omitempty"`
	// 转发到的listener的ID，当action为REDIRECT_TO_LISTENER时生效。当action为REDIRECT_TO_LISTENER时必选

	RedirectListenerId *string `json:"redirect_listener_id,omitempty"`
	// 转发到pool的ID。当action为REDIRECT_TO_POOL时生效。使用说明：redirect_pool不能是listener的default_pool，不能已经被其他listener的l7policy所使用。当action为REDIRECT_TO_LISTENER时，不可指定。不允许更新为空。

	RedirectPoolId *string `json:"redirect_pool_id,omitempty"`
}

func (o UpdateL7policyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateL7policyReq struct{}"
	}

	return strings.Join([]string{"UpdateL7policyReq", string(data)}, " ")
}
