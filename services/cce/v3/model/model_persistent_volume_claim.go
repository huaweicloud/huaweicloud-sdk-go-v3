package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type PersistentVolumeClaim struct {

	// API版本，固定值**v1**
	ApiVersion string `json:"apiVersion" xml:"apiVersion"`

	// API类型，固定值**PersistentVolumeClaim**
	Kind string `json:"kind" xml:"kind"`

	Metadata *PersistentVolumeClaimMetadata `json:"metadata" xml:"metadata"`

	Spec *PersistentVolumeClaimSpec `json:"spec" xml:"spec"`

	Status *PersistentVolumeClaimStatus `json:"status,omitempty" xml:"status"`
}

func (o PersistentVolumeClaim) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistentVolumeClaim struct{}"
	}

	return strings.Join([]string{"PersistentVolumeClaim", string(data)}, " ")
}
