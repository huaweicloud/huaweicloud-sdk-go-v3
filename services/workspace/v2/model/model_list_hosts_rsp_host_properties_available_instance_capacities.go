package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListHostsRspHostPropertiesAvailableInstanceCapacities struct {

	// 规格id。
	Flavor *string `json:"flavor,omitempty"`
}

func (o ListHostsRspHostPropertiesAvailableInstanceCapacities) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostsRspHostPropertiesAvailableInstanceCapacities struct{}"
	}

	return strings.Join([]string{"ListHostsRspHostPropertiesAvailableInstanceCapacities", string(data)}, " ")
}
