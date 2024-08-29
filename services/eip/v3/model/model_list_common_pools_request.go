package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCommonPoolsRequest Request Object
type ListCommonPoolsRequest struct {

	// 显示，形式为\"fields=id&fields=name&...\"  支持字段：id/name/status/type/used/allow_share_bandwidth_types/public_border_group/description
	Fields *[]string `json:"fields,omitempty"`

	// 每页返回的个数取值范围：0~[2000]，其中2000为局点差异项，具体取值由局点决定
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询起始的资源序号
	Offset *int32 `json:"offset,omitempty"`

	// 公共池名称
	Name *string `json:"name,omitempty"`

	// 公共池位于中心还是边缘
	PublicBorderGroup *string `json:"public_border_group,omitempty"`
}

func (o ListCommonPoolsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCommonPoolsRequest struct{}"
	}

	return strings.Join([]string{"ListCommonPoolsRequest", string(data)}, " ")
}
