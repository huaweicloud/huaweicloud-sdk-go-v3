package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageBatchSyncReqImageTextConfig 图文审核黑白词库配置。
type ImageBatchSyncReqImageTextConfig struct {

	// 检测时使用的自定义黑名单词库列表。
	BlackGlossaryNames *[]string `json:"black_glossary_names,omitempty"`

	// 检测时使用的自定义白名单词库列表。
	WhiteGlossaryNames *[]string `json:"white_glossary_names,omitempty"`
}

func (o ImageBatchSyncReqImageTextConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageBatchSyncReqImageTextConfig struct{}"
	}

	return strings.Join([]string{"ImageBatchSyncReqImageTextConfig", string(data)}, " ")
}
