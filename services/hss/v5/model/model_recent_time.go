package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecentTime 发生时间，毫秒
type RecentTime struct {
}

func (o RecentTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecentTime struct{}"
	}

	return strings.Join([]string{"RecentTime", string(data)}, " ")
}
