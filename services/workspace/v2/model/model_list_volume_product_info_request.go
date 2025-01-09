package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVolumeProductInfoRequest Request Object
type ListVolumeProductInfoRequest struct {

	// 可用分区
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 磁盘类型（多个磁盘类型用逗号隔开）： - SATA: 普通IO磁盘 - SAS：高IO磁盘 - SSD：超高IO磁盘
	VolumeType *string `json:"volume_type,omitempty"`
}

func (o ListVolumeProductInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVolumeProductInfoRequest struct{}"
	}

	return strings.Join([]string{"ListVolumeProductInfoRequest", string(data)}, " ")
}
