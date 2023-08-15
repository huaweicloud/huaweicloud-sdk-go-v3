package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVolumeTypesRequest Request Object
type ShowVolumeTypesRequest struct {
}

func (o ShowVolumeTypesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVolumeTypesRequest struct{}"
	}

	return strings.Join([]string{"ShowVolumeTypesRequest", string(data)}, " ")
}
