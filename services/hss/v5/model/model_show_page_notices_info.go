package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowPageNoticesInfo struct {

	// 页面位置
	PageLocation *string `json:"page_location,omitempty"`

	// 通知类型，包含如下两种。 - links :超链接 - text  :文本
	Type *string `json:"type,omitempty"`

	// 通知内容
	Content *string `json:"content,omitempty"`

	// 通知标题
	Title *string `json:"title,omitempty"`

	// 超链接
	Url *string `json:"url,omitempty"`

	// 通知等级，包含如下3种。 - error :紧急 - warn :重要 - prompt :提示
	Level *string `json:"level,omitempty"`
}

func (o ShowPageNoticesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPageNoticesInfo struct{}"
	}

	return strings.Join([]string{"ShowPageNoticesInfo", string(data)}, " ")
}
