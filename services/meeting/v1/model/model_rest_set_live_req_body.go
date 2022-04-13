package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 启停直播消息体。
type RestSetLiveReqBody struct {
	// 默认值为0。 - 0: 停止会议直播。 - 1: 启动会议直播。

	IsLive int32 `json:"isLive"`
}

func (o RestSetLiveReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestSetLiveReqBody struct{}"
	}

	return strings.Join([]string{"RestSetLiveReqBody", string(data)}, " ")
}
