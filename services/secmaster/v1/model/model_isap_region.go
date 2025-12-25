package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsapRegion 区域
type IsapRegion struct {
}

func (o IsapRegion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapRegion struct{}"
	}

	return strings.Join([]string{"IsapRegion", string(data)}, " ")
}
