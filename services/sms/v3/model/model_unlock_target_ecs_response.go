package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnlockTargetEcsResponse Response Object
type UnlockTargetEcsResponse struct {

	// 解锁指定任务的目的端服务器成功
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UnlockTargetEcsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnlockTargetEcsResponse struct{}"
	}

	return strings.Join([]string{"UnlockTargetEcsResponse", string(data)}, " ")
}
