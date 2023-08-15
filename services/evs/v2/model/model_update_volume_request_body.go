package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVolumeRequestBody This is a auto create Body Object
type UpdateVolumeRequestBody struct {
	Volume *UpdateVolumeOption `json:"volume"`
}

func (o UpdateVolumeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVolumeRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateVolumeRequestBody", string(data)}, " ")
}
