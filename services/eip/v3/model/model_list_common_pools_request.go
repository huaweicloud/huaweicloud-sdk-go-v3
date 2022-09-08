package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListCommonPoolsRequest struct {

	// 显示，形式为\"fields=id&fields=name&...\"  支持字段：id/name/status/type/used/allow_share_bandwidth_types/public_border_group
	Fields *string `json:"fields,omitempty"`

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
