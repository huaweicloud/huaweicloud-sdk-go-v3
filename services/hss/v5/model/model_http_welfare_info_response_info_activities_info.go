package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HttpWelfareInfoResponseInfoActivitiesInfo 活动大促信息
type HttpWelfareInfoResponseInfoActivitiesInfo struct {

	// **参数解释**: 活动大促标题 **取值范围**: 字符长度1-256
	Title *string `json:"title,omitempty"`

	// **参数解释**: 活动大促链接，包括中国站，国际站欧洲站，由console根据不同的场景选择跳转 **取值范围**: 字符长度0-65535
	UrlJson *string `json:"url_json,omitempty"`
}

func (o HttpWelfareInfoResponseInfoActivitiesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpWelfareInfoResponseInfoActivitiesInfo struct{}"
	}

	return strings.Join([]string{"HttpWelfareInfoResponseInfoActivitiesInfo", string(data)}, " ")
}
