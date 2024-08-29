package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SmartChatJobsReq 创建交互任务
type SmartChatJobsReq struct {

	// 扩展参数，按照Json格式携带 * city:所在城市
	ExtendParam *string `json:"extend_param,omitempty"`
}

func (o SmartChatJobsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartChatJobsReq struct{}"
	}

	return strings.Join([]string{"SmartChatJobsReq", string(data)}, " ")
}
