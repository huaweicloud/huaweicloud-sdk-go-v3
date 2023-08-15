package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstancesResourceDetail
type InstancesResourceDetail struct {
	Vault *Vault `json:"vault,omitempty"`
}

func (o InstancesResourceDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstancesResourceDetail struct{}"
	}

	return strings.Join([]string{"InstancesResourceDetail", string(data)}, " ")
}
