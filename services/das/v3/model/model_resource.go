package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Resource struct {

	// 根据type过滤查询指定类型的配额。
	Type string `json:"type"`

	// 已创建的资源个数。
	Used int64 `json:"used"`

	// 资源的最大配额数。
	Quota int64 `json:"quota"`

	// 允许修改的配额最小值。
	Min int64 `json:"min"`
}

func (o Resource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resource struct{}"
	}

	return strings.Join([]string{"Resource", string(data)}, " ")
}
