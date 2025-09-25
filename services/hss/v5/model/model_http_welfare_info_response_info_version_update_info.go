package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HttpWelfareInfoResponseInfoVersionUpdateInfo 版本更新信息
type HttpWelfareInfoResponseInfoVersionUpdateInfo struct {

	// **参数解释**: 版本更新标题 **取值范围**: 字符长度1-256
	Title *string `json:"title,omitempty"`

	// **参数解释**: 版本更新链接，包括中国站，国际站欧洲站，由console根据不同的场景选择跳转 **取值范围**: 字符长度0-65535
	UrlJson *string `json:"url_json,omitempty"`
}

func (o HttpWelfareInfoResponseInfoVersionUpdateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpWelfareInfoResponseInfoVersionUpdateInfo struct{}"
	}

	return strings.Join([]string{"HttpWelfareInfoResponseInfoVersionUpdateInfo", string(data)}, " ")
}
