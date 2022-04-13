package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 裸金属服务器挂载云硬盘接口请求结构体
type AttachVolumeBody struct {
	VolumeAttachment *VolumeAttachment `json:"volumeAttachment"`
}

func (o AttachVolumeBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachVolumeBody struct{}"
	}

	return strings.Join([]string{"AttachVolumeBody", string(data)}, " ")
}
