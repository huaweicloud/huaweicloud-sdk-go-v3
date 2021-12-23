package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListEdgeCloudRequest struct {
	// 偏移量。 当前偏移量，默认为0。

	Offset *int32 `json:"offset,omitempty"`
	// 查询返回边缘业务列表当前页面的数量。 取值范围：0~1000。

	Limit *int32 `json:"limit,omitempty"`
	// 边缘业务名称。

	Name *string `json:"name,omitempty"`
	// 边缘业务ID。

	Id *string `json:"id,omitempty"`
}

func (o ListEdgeCloudRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEdgeCloudRequest struct{}"
	}

	return strings.Join([]string{"ListEdgeCloudRequest", string(data)}, " ")
}
