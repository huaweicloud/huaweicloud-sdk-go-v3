package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NetworkStatus 资源池状态信息。
type NetworkStatus struct {

	// **参数解释**：网络资源的当前状态。 **取值范围**：可选值如下： - Creating：网络创建中。 - Active：网络正常。 - Abnormal：网络异常。
	Phase string `json:"phase"`

	ConnectionStatus *NetworkConnectionStatus `json:"connectionStatus,omitempty"`
}

func (o NetworkStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkStatus struct{}"
	}

	return strings.Join([]string{"NetworkStatus", string(data)}, " ")
}
