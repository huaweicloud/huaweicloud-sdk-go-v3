package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListClustersRequest struct {
	// 分页参数，列表的偏移量，默认值为0

	Offset *int32 `json:"offset,omitempty"`
	// 分页参数，列表当前分页的数量限制，默认为10。

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListClustersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClustersRequest struct{}"
	}

	return strings.Join([]string{"ListClustersRequest", string(data)}, " ")
}
