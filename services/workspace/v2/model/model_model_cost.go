package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelCost 模型费用信息。
type ModelCost struct {

	// 每百万输入Token费用。
	Input *float64 `json:"input,omitempty"`

	// 每百万输出Token费用。
	Output *float64 `json:"output,omitempty"`

	// 每百万缓存读取Token费用。
	CacheRead *float64 `json:"cache_read,omitempty"`

	// 每百万缓存写入Token费用。
	CacheWrite *float64 `json:"cache_write,omitempty"`
}

func (o ModelCost) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelCost struct{}"
	}

	return strings.Join([]string{"ModelCost", string(data)}, " ")
}
