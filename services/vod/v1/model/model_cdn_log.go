package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CdnLog struct {

	// 域名名称。
	DomainName *string `json:"domain_name,omitempty" xml:"domain_name"`

	// 查询起始时间。
	StartTime *string `json:"start_time,omitempty" xml:"start_time"`

	// 查询结束时间。
	EndTime *string `json:"end_time,omitempty" xml:"end_time"`

	// 日志名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 日志大小。  单位：byte。
	Size *int64 `json:"size,omitempty" xml:"size"`

	// 日志下载链接,日志文件[参数说明](https://support.huaweicloud.com/usermanual-cdn/zh-cn_topic_0073337424.html)。
	Link *string `json:"link,omitempty" xml:"link"`
}

func (o CdnLog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CdnLog struct{}"
	}

	return strings.Join([]string{"CdnLog", string(data)}, " ")
}
