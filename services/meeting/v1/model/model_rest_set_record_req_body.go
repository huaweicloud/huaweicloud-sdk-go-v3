package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 启停录制消息体。
type RestSetRecordReqBody struct {

	// 默认值为0。 - 0: 停止会议录制。 - 1: 启动会议录制。
	IsRecord int32 `json:"isRecord" xml:"isRecord"`
}

func (o RestSetRecordReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestSetRecordReqBody struct{}"
	}

	return strings.Join([]string{"RestSetRecordReqBody", string(data)}, " ")
}
