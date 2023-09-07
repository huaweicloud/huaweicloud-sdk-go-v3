package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVolumeReq struct {
	ApiVersion *ApiVersionObj `json:"api_version"`

	Kind *VolumeKindObj `json:"kind"`

	Spec *VolumeSpec `json:"spec"`
}

func (o CreateVolumeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVolumeReq struct{}"
	}

	return strings.Join([]string{"CreateVolumeReq", string(data)}, " ")
}
