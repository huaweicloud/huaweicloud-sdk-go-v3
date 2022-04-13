package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AttachBaremetalServerVolumeRequest struct {
	// 裸金属服务器ID

	ServerId string `json:"server_id"`

	Body *AttachVolumeBody `json:"body,omitempty"`
}

func (o AttachBaremetalServerVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachBaremetalServerVolumeRequest struct{}"
	}

	return strings.Join([]string{"AttachBaremetalServerVolumeRequest", string(data)}, " ")
}
