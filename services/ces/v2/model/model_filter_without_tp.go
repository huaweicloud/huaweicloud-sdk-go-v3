package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FilterWithoutTp 聚合方式, 支持的值为(average|min|max|sum)
type FilterWithoutTp struct {
}

func (o FilterWithoutTp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FilterWithoutTp struct{}"
	}

	return strings.Join([]string{"FilterWithoutTp", string(data)}, " ")
}
