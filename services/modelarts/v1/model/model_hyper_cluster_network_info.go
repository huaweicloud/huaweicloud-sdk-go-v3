package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HyperClusterNetworkInfo struct {

	// **参数解释**：hyper cluster的子网名称。 **取值范围**：^[-_.a-zA-Z0-9]{1,64}$。
	HyperClusterSubnetId *string `json:"hyper_cluster_subnet_id,omitempty"`

	// **参数解释**：是否默认。 **约束限制**：不涉及。 **取值范围**： - true：默认网络 - false：非默认网络
	IsDefault *bool `json:"is_default,omitempty"`
}

func (o HyperClusterNetworkInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HyperClusterNetworkInfo struct{}"
	}

	return strings.Join([]string{"HyperClusterNetworkInfo", string(data)}, " ")
}
