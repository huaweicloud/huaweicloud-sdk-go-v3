package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新后端云服务器组的请求体
type UpdatePoolReq struct {

	// 后端云服务器组的负载均衡算法，取值：ROUND_ROBIN：加权轮询算法；LEAST_CONNECTIONS：加权最少连接算法；SOURCE_IP：源IP算法；当该字段的取值为SOURCE_IP时，后端云服务器组绑定的后端云服务器的weight字段无效。
	LbAlgorithm *string `json:"lb_algorithm,omitempty" xml:"lb_algorithm"`

	// 后端云服务器组的名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 后端云服务器组的描述信息
	Description *string `json:"description,omitempty" xml:"description"`

	// 后端云服务器组的管理状态；该字段为预留字段，暂未启用。只支持更新为true。
	AdminStateUp *bool `json:"admin_state_up,omitempty" xml:"admin_state_up"`

	SessionPersistence *SessionPersistence `json:"session_persistence,omitempty" xml:"session_persistence"`
}

func (o UpdatePoolReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePoolReq struct{}"
	}

	return strings.Join([]string{"UpdatePoolReq", string(data)}, " ")
}
