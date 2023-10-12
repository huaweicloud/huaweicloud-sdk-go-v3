package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RedisTableDetail 重分布表详情
type RedisTableDetail struct {

	// 具体数据
	Data *[]RedisTable `json:"data,omitempty"`

	// 总条数
	Total *int32 `json:"total,omitempty"`
}

func (o RedisTableDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedisTableDetail struct{}"
	}

	return strings.Join([]string{"RedisTableDetail", string(data)}, " ")
}
