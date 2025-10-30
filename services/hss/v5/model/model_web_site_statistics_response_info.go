package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WebSiteStatisticsResponseInfo **参数解释** Web站点统计信息列表
type WebSiteStatisticsResponseInfo struct {

	// **参数解释**: Web站点域名 **取值范围**: 字符长度0-256
	Domain *string `json:"domain,omitempty"`

	// **参数解释** Web站点统计信息总数 **取值范围** 最小值0，最大值300000
	Num *int32 `json:"num,omitempty"`
}

func (o WebSiteStatisticsResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebSiteStatisticsResponseInfo struct{}"
	}

	return strings.Join([]string{"WebSiteStatisticsResponseInfo", string(data)}, " ")
}
