package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteCloudPersistentVolumeClaimsResponse struct {

	// API版本，固定值**v1**
	ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion"`

	// API类型，固定值**PersistentVolumeClaim**
	Kind *string `json:"kind,omitempty" xml:"kind"`

	Metadata *PersistentVolumeClaimMetadata `json:"metadata,omitempty" xml:"metadata"`

	Spec *PersistentVolumeClaimSpec `json:"spec,omitempty" xml:"spec"`

	Status         *PersistentVolumeClaimStatus `json:"status,omitempty" xml:"status"`
	HttpStatusCode int                          `json:"-"`
}

func (o DeleteCloudPersistentVolumeClaimsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCloudPersistentVolumeClaimsResponse struct{}"
	}

	return strings.Join([]string{"DeleteCloudPersistentVolumeClaimsResponse", string(data)}, " ")
}
