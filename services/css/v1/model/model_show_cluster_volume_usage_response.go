package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterVolumeUsageResponse Response Object
type ShowClusterVolumeUsageResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowClusterVolumeUsageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterVolumeUsageResponse struct{}"
	}

	return strings.Join([]string{"ShowClusterVolumeUsageResponse", string(data)}, " ")
}
