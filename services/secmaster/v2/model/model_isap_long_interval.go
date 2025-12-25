package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsapLongInterval 长整型间隔时长
type IsapLongInterval struct {
}

func (o IsapLongInterval) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapLongInterval struct{}"
	}

	return strings.Join([]string{"IsapLongInterval", string(data)}, " ")
}
