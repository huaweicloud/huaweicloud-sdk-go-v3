package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StreamTranscodingTemplate struct {
	// 推流域名

	Domain string `json:"domain"`
	// 应用名称。 默认为“live”，若您需要自定义应用名称，请先提交工单申请。

	AppName string `json:"app_name"`
	// 视频质量信息

	QualityInfo []QualityInfo `json:"quality_info"`
}

func (o StreamTranscodingTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StreamTranscodingTemplate struct{}"
	}

	return strings.Join([]string{"StreamTranscodingTemplate", string(data)}, " ")
}
