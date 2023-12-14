package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LtslogInfo LTS日志信息
type LtslogInfo struct {

	// 配置状态
	Status string `json:"status"`

	// 日志ID
	Id string `json:"id"`

	// 日志类型
	LogType string `json:"log_type"`

	// 日志描述
	LogDesc string `json:"log_desc"`

	// LTS日志访问URL
	AccessUrl string `json:"access_url"`
}

func (o LtslogInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtslogInfo struct{}"
	}

	return strings.Join([]string{"LtslogInfo", string(data)}, " ")
}
