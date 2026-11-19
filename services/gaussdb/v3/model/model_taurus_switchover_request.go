package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaurusSwitchoverRequest struct {

	// 只读节点ID，倒换后为主节点。
	NodeId string `json:"node_id"`

	// **参数解释**：  是否进行强制倒换。强制倒换将导致主节点重启，未提交的事务会被中断，请在业务低峰期执行该操作。  **约束限制**：  不涉及。  **取值范围**：  - true：强制倒换。 - false：普通倒换。  **默认取值**：  false。
	Force *bool `json:"force,omitempty"`
}

func (o TaurusSwitchoverRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaurusSwitchoverRequest struct{}"
	}

	return strings.Join([]string{"TaurusSwitchoverRequest", string(data)}, " ")
}
