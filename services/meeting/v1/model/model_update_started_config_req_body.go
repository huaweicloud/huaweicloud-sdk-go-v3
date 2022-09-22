package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 修改会议配置请求。
type UpdateStartedConfigReqBody struct {

	// 锁定共享标志位。 * 0: 不锁定 * 1: 锁定
	LockSharing *int32 `json:"lockSharing,omitempty"`

	// 允许加入会议的范围。 - 0: 所有用户 - 2: 企业内用户 - 3: 被邀请用户
	CallInRestriction *int32 `json:"callInRestriction,omitempty"`
}

func (o UpdateStartedConfigReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStartedConfigReqBody struct{}"
	}

	return strings.Join([]string{"UpdateStartedConfigReqBody", string(data)}, " ")
}
