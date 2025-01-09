package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigKeyStringSet 查询配置的key。
type ConfigKeyStringSet struct {

	// 查询企业配置请求的key，一次请求数量区间 [0, 100]。
	Items []string `json:"items"`
}

func (o ConfigKeyStringSet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigKeyStringSet struct{}"
	}

	return strings.Join([]string{"ConfigKeyStringSet", string(data)}, " ")
}
