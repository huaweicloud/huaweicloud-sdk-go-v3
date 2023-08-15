package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStreamCountRequest Request Object
type ShowStreamCountRequest struct {

	// 推流域名列表，最多支持查询100个域名，多个域名以逗号分隔，若查询多个域名，则返回的是多个域名合并数据。
	PublishDomains []string `json:"publish_domains"`

	// 起始时间。日期格式按照ISO8601表示法，并使用UTC时间。  格式为：YYYY-MM-DDThh:mm:ssZ。最大查询跨度31天，最大查询周期1年。  若参数为空，默认查询7天数据。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间。日期格式按照ISO8601表示法，并使用UTC时间。 格式为：YYYY-MM-DDThh:mm:ssZ。  若参数为空，默认为当前时间，最大查询跨度31天， 最大查询跨度31天，最大查询周期1年。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ShowStreamCountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStreamCountRequest struct{}"
	}

	return strings.Join([]string{"ShowStreamCountRequest", string(data)}, " ")
}
