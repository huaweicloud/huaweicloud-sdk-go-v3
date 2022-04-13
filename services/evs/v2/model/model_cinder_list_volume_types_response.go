package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CinderListVolumeTypesResponse struct {
	VolumeTypes    *[]VolumeType `json:"volume_types,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o CinderListVolumeTypesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderListVolumeTypesResponse struct{}"
	}

	return strings.Join([]string{"CinderListVolumeTypesResponse", string(data)}, " ")
}
