package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnlargeVolume struct {
	EnlargeVolume *EnlargeVolumeObject `json:"enlarge_volume"`
}

func (o EnlargeVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnlargeVolume struct{}"
	}

	return strings.Join([]string{"EnlargeVolume", string(data)}, " ")
}
