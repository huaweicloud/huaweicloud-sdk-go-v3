package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StreamForbiddenOnceSetting struct {

	// 直播推流域名
	Domain string `json:"domain"`

	// 应用名称
	AppName string `json:"app_name"`

	// 流名称
	StreamName string `json:"stream_name"`
}

func (o StreamForbiddenOnceSetting) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StreamForbiddenOnceSetting struct{}"
	}

	return strings.Join([]string{"StreamForbiddenOnceSetting", string(data)}, " ")
}
