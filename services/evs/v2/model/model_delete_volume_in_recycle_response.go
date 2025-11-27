package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVolumeInRecycleResponse Response Object
type DeleteVolumeInRecycleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVolumeInRecycleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVolumeInRecycleResponse struct{}"
	}

	return strings.Join([]string{"DeleteVolumeInRecycleResponse", string(data)}, " ")
}
