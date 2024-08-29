package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetypeVolumeRequestBody This is a auto create Body Object
type RetypeVolumeRequestBody struct {
	BssParam *BssParamForRetypeVolume `json:"bssParam,omitempty"`

	OsRetype *RetypeVolume `json:"os-retype"`
}

func (o RetypeVolumeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetypeVolumeRequestBody struct{}"
	}

	return strings.Join([]string{"RetypeVolumeRequestBody", string(data)}, " ")
}
