package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 锁定会场视频源请求体
type RestLockSiteViewReqBody struct {
	// - 0: 取消锁定。 - 1: 锁定。

	Status int32 `json:"status"`
	// 被锁定视频源的与会者标识。

	ParticipantID string `json:"participantID"`
}

func (o RestLockSiteViewReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestLockSiteViewReqBody struct{}"
	}

	return strings.Join([]string{"RestLockSiteViewReqBody", string(data)}, " ")
}
