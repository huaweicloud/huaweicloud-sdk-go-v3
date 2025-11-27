package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListManagedClustersRequest Request Object
type ListManagedClustersRequest struct {

	// 是否已导入到ucs
	Unimported *bool `json:"unimported,omitempty"`

	// 分页获取列表时，页的大小，默认为-1
	Limit *int32 `json:"limit,omitempty"`

	// 分页获取列表时，起始偏移量，默认为0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListManagedClustersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListManagedClustersRequest struct{}"
	}

	return strings.Join([]string{"ListManagedClustersRequest", string(data)}, " ")
}
