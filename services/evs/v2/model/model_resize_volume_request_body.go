package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type ResizeVolumeRequestBody struct {
	BssParam *BssParamForResizeVolume `json:"bssParam,omitempty" xml:"bssParam"`

	OsExtend *OsExtend `json:"os-extend" xml:"os-extend"`
}

func (o ResizeVolumeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeVolumeRequestBody struct{}"
	}

	return strings.Join([]string{"ResizeVolumeRequestBody", string(data)}, " ")
}
