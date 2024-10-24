package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReservedResource 通用预留资源
type ReservedResource struct {
	Apu *ResourceDef `json:"apu,omitempty"`

	Dpu *ResourceDef `json:"dpu,omitempty"`

	Mu *ResourceDemand `json:"mu,omitempty"`
}

func (o ReservedResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReservedResource struct{}"
	}

	return strings.Join([]string{"ReservedResource", string(data)}, " ")
}
