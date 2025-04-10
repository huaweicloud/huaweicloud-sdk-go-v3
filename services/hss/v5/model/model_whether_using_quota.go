package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WhetherUsingQuota 是否使用病毒查杀按次计费配额
type WhetherUsingQuota struct {
}

func (o WhetherUsingQuota) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WhetherUsingQuota struct{}"
	}

	return strings.Join([]string{"WhetherUsingQuota", string(data)}, " ")
}
