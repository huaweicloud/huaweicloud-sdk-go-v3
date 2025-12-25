package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsapIntInterval 整型间隔时长
type IsapIntInterval struct {
}

func (o IsapIntInterval) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapIntInterval struct{}"
	}

	return strings.Join([]string{"IsapIntInterval", string(data)}, " ")
}
