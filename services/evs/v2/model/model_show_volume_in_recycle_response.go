package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVolumeInRecycleResponse Response Object
type ShowVolumeInRecycleResponse struct {
	Volume         *RecycleBinVolume `json:"volume,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowVolumeInRecycleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVolumeInRecycleResponse struct{}"
	}

	return strings.Join([]string{"ShowVolumeInRecycleResponse", string(data)}, " ")
}
