package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CdmClusterFlavor 集群规格。
type CdmClusterFlavor struct {

	// CPU。
	Cpu *int32 `json:"cpu,omitempty"`

	// 内存。
	Ram *int32 `json:"ram,omitempty"`

	// 规格名称。
	Name *string `json:"name,omitempty"`

	// region。
	Region *string `json:"region,omitempty"`

	// 类型名称。
	Typename *string `json:"typename,omitempty"`

	// 集群模式。
	ClusterMode *string `json:"clusterMode,omitempty"`

	// 规格状态。
	Status *string `json:"status,omitempty"`

	// 规格ID。
	StrId *string `json:"str_id,omitempty"`
}

func (o CdmClusterFlavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CdmClusterFlavor struct{}"
	}

	return strings.Join([]string{"CdmClusterFlavor", string(data)}, " ")
}
