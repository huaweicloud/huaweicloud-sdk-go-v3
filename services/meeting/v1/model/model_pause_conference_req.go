package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PauseConferenceReq 主持人暂停/取消暂停会议请求消息体
type PauseConferenceReq struct {

	// 主持人暂停/取消暂停会议 0：会议正常 1：会议暂停
	Pause int32 `json:"pause"`
}

func (o PauseConferenceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PauseConferenceReq struct{}"
	}

	return strings.Join([]string{"PauseConferenceReq", string(data)}, " ")
}
