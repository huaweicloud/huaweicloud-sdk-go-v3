package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeVolumeRequestBody This is a auto create Body Object
type ResizeVolumeRequestBody struct {
	BssParam *BssParamForResizeVolume `json:"bssParam,omitempty"`

	OsExtend *OsExtend `json:"os-extend"`
}

func (o ResizeVolumeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeVolumeRequestBody struct{}"
	}

	return strings.Join([]string{"ResizeVolumeRequestBody", string(data)}, " ")
}
