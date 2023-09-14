package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccessConfigQueryLogInfo 日志接入日志详情
type AccessConfigQueryLogInfo struct {
}

func (o AccessConfigQueryLogInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigQueryLogInfo struct{}"
	}

	return strings.Join([]string{"AccessConfigQueryLogInfo", string(data)}, " ")
}
