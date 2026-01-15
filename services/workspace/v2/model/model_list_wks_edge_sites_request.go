package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWksEdgeSitesRequest Request Object
type ListWksEdgeSitesRequest struct {

	// 根据边缘小站名称查询。
	Name *string `json:"name,omitempty"`

	// 根据边缘可用区ID查询。
	AvailabilityZoneId *string `json:"availability_zone_id,omitempty"`

	// 根据边缘小站部署状态查询。
	Status *string `json:"status,omitempty"`

	// 每页数量，范围0-1000，默认1000。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，默认0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListWksEdgeSitesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWksEdgeSitesRequest struct{}"
	}

	return strings.Join([]string{"ListWksEdgeSitesRequest", string(data)}, " ")
}
