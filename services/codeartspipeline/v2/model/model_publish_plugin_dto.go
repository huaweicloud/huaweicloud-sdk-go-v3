package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PublishPluginDto struct {

	// 插件名
	PluginName *string `json:"plugin_name,omitempty"`

	// 版本
	Version *string `json:"version,omitempty"`

	// 发布商ID
	PublisherUniqueId *string `json:"publisher_unique_id,omitempty"`
}

func (o PublishPluginDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishPluginDto struct{}"
	}

	return strings.Join([]string{"PublishPluginDto", string(data)}, " ")
}
