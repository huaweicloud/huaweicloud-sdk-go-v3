package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBandwidthOption 更新带宽参数
type UpdateBandwidthOption struct {

	// 带宽名称。
	Name *string `json:"name,omitempty"`

	// 带宽大小。
	Size *int32 `json:"size,omitempty"`
}

func (o UpdateBandwidthOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBandwidthOption struct{}"
	}

	return strings.Join([]string{"UpdateBandwidthOption", string(data)}, " ")
}
