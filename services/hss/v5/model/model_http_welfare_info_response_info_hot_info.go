package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HttpWelfareInfoResponseInfoHotInfo 热点事件信息
type HttpWelfareInfoResponseInfoHotInfo struct {

	// **参数解释**: 热点事件标题 **取值范围**: 字符长度1-256
	Title *string `json:"title,omitempty"`

	// **参数解释**: 热点事件链接，包括中国站，国际站欧洲站，由console根据不同的场景选择跳转 **取值范围**: 字符长度0-65535
	UrlJson *string `json:"url_json,omitempty"`
}

func (o HttpWelfareInfoResponseInfoHotInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpWelfareInfoResponseInfoHotInfo struct{}"
	}

	return strings.Join([]string{"HttpWelfareInfoResponseInfoHotInfo", string(data)}, " ")
}
