package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DetachBaremetalServerVolumeRequest struct {

	// 裸金属服务器ID
	ServerId string `json:"server_id"`

	// 裸金属服务器的云磁盘ID
	AttachmentId string `json:"attachment_id"`
}

func (o DetachBaremetalServerVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachBaremetalServerVolumeRequest struct{}"
	}

	return strings.Join([]string{"DetachBaremetalServerVolumeRequest", string(data)}, " ")
}
