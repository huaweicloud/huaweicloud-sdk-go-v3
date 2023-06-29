package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVolumeTypesResponse Response Object
type ShowVolumeTypesResponse struct {

	// 硬盘类型列表。
	VolumeTypes    *[]VolumeType `json:"volume_types,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowVolumeTypesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVolumeTypesResponse struct{}"
	}

	return strings.Join([]string{"ShowVolumeTypesResponse", string(data)}, " ")
}
