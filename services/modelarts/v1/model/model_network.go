package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Network 网络的详细信息。
type Network struct {

	// **参数解释**：资源的API版本。 **取值范围**：可选值如下： - v1：当前资源版本为v1。
	ApiVersion string `json:"apiVersion"`

	// **参数解释**：资源的类型。 **取值范围**：可选值如下： - Network：网络。
	Kind string `json:"kind"`

	Metadata *NetworkMetadata `json:"metadata"`

	Spec *NetworkSpec `json:"spec"`

	Status *NetworkStatus `json:"status,omitempty"`
}

func (o Network) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Network struct{}"
	}

	return strings.Join([]string{"Network", string(data)}, " ")
}
