package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SupabaseManagedInstanceInfo 纳管实例信息。
type SupabaseManagedInstanceInfo struct {
	InstanceId *string `json:"instance_id,omitempty"`

	Status *string `json:"status,omitempty"`
}

func (o SupabaseManagedInstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupabaseManagedInstanceInfo struct{}"
	}

	return strings.Join([]string{"SupabaseManagedInstanceInfo", string(data)}, " ")
}
