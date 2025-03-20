package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MaxSessionDuration 委托或信任委托最大会话时长，默认为3600秒。
type MaxSessionDuration struct {
}

func (o MaxSessionDuration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MaxSessionDuration struct{}"
	}

	return strings.Join([]string{"MaxSessionDuration", string(data)}, " ")
}
