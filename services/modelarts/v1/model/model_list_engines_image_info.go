package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnginesImageInfo 引擎具体信息。
type ListEnginesImageInfo struct {

	// cpu规格下对应镜像。
	CpuImageUrl *string `json:"cpu_image_url,omitempty"`

	// gpu[或者Ascend](tag:hc,hk,fcs_super)规格下对应镜像。
	GpuImageUrl *string `json:"gpu_image_url,omitempty"`

	// 镜像版本。
	ImageVersion *string `json:"image_version,omitempty"`
}

func (o ListEnginesImageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnginesImageInfo struct{}"
	}

	return strings.Join([]string{"ListEnginesImageInfo", string(data)}, " ")
}
