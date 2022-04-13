package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListKeypairsRequest struct {
	// 查询返回keypair列表当前页面的数量。 取值范围：0~1000。

	Limit *int32 `json:"limit,omitempty"`
	// 偏移量。 当前偏移量，默认为0。

	Offset *int32 `json:"offset,omitempty"`
	// 根据名称查询keypair列表。

	Name *string `json:"name,omitempty"`
}

func (o ListKeypairsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKeypairsRequest struct{}"
	}

	return strings.Join([]string{"ListKeypairsRequest", string(data)}, " ")
}
