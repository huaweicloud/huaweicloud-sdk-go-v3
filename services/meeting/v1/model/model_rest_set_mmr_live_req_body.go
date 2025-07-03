package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestSetMmrLiveReqBody rest设置Mmr直播请求体
type RestSetMmrLiveReqBody struct {

	// 0：停止Mmr会议直播 1：启动Mmr会议直播
	LiveState *int32 `json:"liveState,omitempty"`
}

func (o RestSetMmrLiveReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestSetMmrLiveReqBody struct{}"
	}

	return strings.Join([]string{"RestSetMmrLiveReqBody", string(data)}, " ")
}
