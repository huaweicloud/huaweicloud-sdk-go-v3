package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CdmClusterAvailabilityZone 集群的可用区。
type CdmClusterAvailabilityZone struct {

	// 可用区ID。
	AvailableZoneId *string `json:"availableZoneId,omitempty"`

	// 可用区名称。
	AvailableZoneName *string `json:"availableZoneName,omitempty"`

	// 可用区码。
	AvailableZoneCode *string `json:"availableZoneCode,omitempty"`

	// 可用区状态。
	AzStatus *string `json:"azStatus,omitempty"`

	// 可用区类型。
	Type *string `json:"type,omitempty"`

	// 可用区tag。
	Tags *interface{} `json:"tags,omitempty"`
}

func (o CdmClusterAvailabilityZone) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CdmClusterAvailabilityZone struct{}"
	}

	return strings.Join([]string{"CdmClusterAvailabilityZone", string(data)}, " ")
}
