package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsapListCount 计数
type IsapListCount struct {
}

func (o IsapListCount) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapListCount struct{}"
	}

	return strings.Join([]string{"IsapListCount", string(data)}, " ")
}
