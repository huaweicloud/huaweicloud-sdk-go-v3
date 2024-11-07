package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBandwidthPackageLevelsRequest Request Object
type ListBandwidthPackageLevelsRequest struct {

	// 根据带宽包等级进行查询
	Level *string `json:"level,omitempty"`

	// 根据名称进行模糊查询
	Name *string `json:"name,omitempty"`
}

func (o ListBandwidthPackageLevelsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBandwidthPackageLevelsRequest struct{}"
	}

	return strings.Join([]string{"ListBandwidthPackageLevelsRequest", string(data)}, " ")
}
