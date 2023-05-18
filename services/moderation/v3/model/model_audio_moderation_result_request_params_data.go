package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建作业时传的data参数
type AudioModerationResultRequestParamsData struct {
	Url *string `json:"url,omitempty"`
}

func (o AudioModerationResultRequestParamsData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AudioModerationResultRequestParamsData struct{}"
	}

	return strings.Join([]string{"AudioModerationResultRequestParamsData", string(data)}, " ")
}
