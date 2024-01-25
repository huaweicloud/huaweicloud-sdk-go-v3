package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVolumeRequest Request Object
type ListVolumeRequest struct {

	// 查询硬盘列表当前页面的数量。 取值范围：0~1000
	Limit *int32 `json:"limit,omitempty"`

	// 根据名称查询硬盘列表。
	Name *string `json:"name,omitempty"`

	// 查询的偏移量。默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 根据状态查询硬盘列表。
	Status *string `json:"status,omitempty"`
}

func (o ListVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVolumeRequest struct{}"
	}

	return strings.Join([]string{"ListVolumeRequest", string(data)}, " ")
}
