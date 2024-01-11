package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StatusStatistics 资源统计信息
type StatusStatistics struct {

	// 活跃资源
	Active *int64 `json:"active,omitempty"`

	// 总资源
	Total *int64 `json:"total,omitempty"`
}

func (o StatusStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatusStatistics struct{}"
	}

	return strings.Join([]string{"StatusStatistics", string(data)}, " ")
}
