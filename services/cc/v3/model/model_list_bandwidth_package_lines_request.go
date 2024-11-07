package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBandwidthPackageLinesRequest Request Object
type ListBandwidthPackageLinesRequest struct {

	// 根据带宽包等级进行查询
	Level *string `json:"level,omitempty"`

	// 根据名称进行模糊查询
	Name *string `json:"name,omitempty"`
}

func (o ListBandwidthPackageLinesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBandwidthPackageLinesRequest struct{}"
	}

	return strings.Join([]string{"ListBandwidthPackageLinesRequest", string(data)}, " ")
}
