package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestSetRecordReqBody 启停会议录制请求。
type RestSetRecordReqBody struct {

	// 录制启停开关。默认值为0。 - 0: 停止会议录制 - 1: 启动会议录制
	IsRecord int32 `json:"isRecord"`
}

func (o RestSetRecordReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestSetRecordReqBody struct{}"
	}

	return strings.Join([]string{"RestSetRecordReqBody", string(data)}, " ")
}
