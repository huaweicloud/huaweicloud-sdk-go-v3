package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WebSiteStatisticsResponseInfo Web站点统计信息列表
type WebSiteStatisticsResponseInfo struct {

	// Web站点域名
	Domain *string `json:"domain,omitempty"`

	// Web站点统计信息总数
	Num *int32 `json:"num,omitempty"`
}

func (o WebSiteStatisticsResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebSiteStatisticsResponseInfo struct{}"
	}

	return strings.Join([]string{"WebSiteStatisticsResponseInfo", string(data)}, " ")
}
