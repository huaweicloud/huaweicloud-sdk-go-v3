package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchSgcComputeDimensionsResponse Response Object
type SearchSgcComputeDimensionsResponse struct {

	// 总数
	Count *int64 `json:"count,omitempty"`

	// 计算成本实例列表
	Computes       *[]ComputeDimension `json:"computes,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o SearchSgcComputeDimensionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchSgcComputeDimensionsResponse struct{}"
	}

	return strings.Join([]string{"SearchSgcComputeDimensionsResponse", string(data)}, " ")
}
