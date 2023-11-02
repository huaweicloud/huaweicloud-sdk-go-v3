package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckResourceInfo 资源信息。
type CheckResourceInfo struct {

	// 企业项目ID。action为createInstance时必填。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 实例数量。action为createInstance时必填。
	InstanceNum *int32 `json:"instance_num,omitempty"`

	// 实例类型，目前仅支持Cluster。action为createInstance时必填。
	Mode *string `json:"mode,omitempty"`

	// 可用区类型，单可用区single或多可用区multi。action为createInstance时必填。
	AvailabilityZoneMode *string `json:"availability_zone_mode,omitempty"`

	// 节点数量。action为createInstance、createReadonlyNode时必填。
	NodeNum *int32 `json:"node_num,omitempty"`

	// 规格码。action为createInstance、resizeFlavor时必填。
	FlavorRef *string `json:"flavor_ref,omitempty"`

	// 可用区码。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 子网ID。action为createInstance时必填。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 实例ID。action为createReadonlyNode、resizeFlavor时必填。
	InstanceId *string `json:"instance_id,omitempty"`
}

func (o CheckResourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckResourceInfo struct{}"
	}

	return strings.Join([]string{"CheckResourceInfo", string(data)}, " ")
}
