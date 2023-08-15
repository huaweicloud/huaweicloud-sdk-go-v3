package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDiskUsageResponse Response Object
type ShowDiskUsageResponse struct {

	// 磁盘信息列表
	Volumes        *[]DiskVolumes `json:"volumes,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowDiskUsageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDiskUsageResponse struct{}"
	}

	return strings.Join([]string{"ShowDiskUsageResponse", string(data)}, " ")
}
