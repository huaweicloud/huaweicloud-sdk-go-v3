package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRoutetablesRequest Request Object
type ListRoutetablesRequest struct {

	// 路由表ID
	Id *string `json:"id,omitempty"`

	// 每页的最大数
	Limit *int32 `json:"limit,omitempty"`

	// 路由表名称。
	Name *string `json:"name,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// vpc的ID
	VpcId *string `json:"vpc_id,omitempty"`
}

func (o ListRoutetablesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRoutetablesRequest struct{}"
	}

	return strings.Join([]string{"ListRoutetablesRequest", string(data)}, " ")
}
