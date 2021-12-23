package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResizeInstanceVolumeRequestBody struct {
	Volume *ResizeInstanceVolumeOption `json:"volume"`
}

func (o ResizeInstanceVolumeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeInstanceVolumeRequestBody struct{}"
	}

	return strings.Join([]string{"ResizeInstanceVolumeRequestBody", string(data)}, " ")
}
