package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVolumeReq struct {

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion string `json:"api_version"`

	// API类型，固定值“Volume”，该值不可修改。
	Kind string `json:"kind"`

	Spec *VolumeSpec `json:"spec"`
}

func (o CreateVolumeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVolumeReq struct{}"
	}

	return strings.Join([]string{"CreateVolumeReq", string(data)}, " ")
}
