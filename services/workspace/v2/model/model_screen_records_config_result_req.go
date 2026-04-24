package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScreenRecordsConfigResultReq 录屏记录配置结果请求体。
type ScreenRecordsConfigResultReq struct {

	// 录屏记录。
	Configs *[]ScreenRecordsConfigResultReqConfigs `json:"configs,omitempty"`
}

func (o ScreenRecordsConfigResultReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScreenRecordsConfigResultReq struct{}"
	}

	return strings.Join([]string{"ScreenRecordsConfigResultReq", string(data)}, " ")
}
