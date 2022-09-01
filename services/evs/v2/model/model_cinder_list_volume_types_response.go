package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CinderListVolumeTypesResponse struct {
	VolumeTypes    *[]VolumeType `json:"volume_types,omitempty" xml:"volume_types"`
	HttpStatusCode int           `json:"-"`
}

func (o CinderListVolumeTypesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderListVolumeTypesResponse struct{}"
	}

	return strings.Join([]string{"CinderListVolumeTypesResponse", string(data)}, " ")
}
