package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CinderListVolumeTypesRequest struct {
}

func (o CinderListVolumeTypesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderListVolumeTypesRequest struct{}"
	}

	return strings.Join([]string{"CinderListVolumeTypesRequest", string(data)}, " ")
}
