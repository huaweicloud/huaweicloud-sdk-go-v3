package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NetworkSpec 对外网络的描述。
type NetworkSpec struct {

	// **参数解释**：网络的cidr值。 **取值范围**： - 172.16.0.0/12~24 - 192.168.0.0/16~24
	Cidr string `json:"cidr"`

	Connection *NetworkConnection `json:"connection,omitempty"`
}

func (o NetworkSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkSpec struct{}"
	}

	return strings.Join([]string{"NetworkSpec", string(data)}, " ")
}
