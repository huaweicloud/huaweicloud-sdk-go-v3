package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 邀请共享请求。
type InviteShareDto struct {

	// 邀请标志。 * 0：取消邀请 * 1：邀请
	Share int32 `json:"share"`
}

func (o InviteShareDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InviteShareDto struct{}"
	}

	return strings.Join([]string{"InviteShareDto", string(data)}, " ")
}
