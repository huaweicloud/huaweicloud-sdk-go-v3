package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Filter 聚合方式, 支持的值为(average|min|max|sum)
type Filter struct {
}

func (o Filter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Filter struct{}"
	}

	return strings.Join([]string{"Filter", string(data)}, " ")
}
