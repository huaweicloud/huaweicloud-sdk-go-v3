package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPlayDomainStreamInfoRequest Request Object
type ListPlayDomainStreamInfoRequest struct {

	// 播放域名列表，最多支持查询10个域名，多个域名以逗号分隔。  如果不传入域名，则查询租户下所有播放域名的流数据。
	PlayDomains *[]string `json:"play_domains,omitempty"`

	// 查询数据的时间点，精确到分钟。  日期格式按照ISO8601表示法，并使用UTC时间。  格式为：YYYY-MM-DDThh:mm:ssZ，最大查询周期七天。  时间必须为时间粒度整时刻点，如：2024-02-02T08:01:00Z。
	Time string `json:"time"`
}

func (o ListPlayDomainStreamInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPlayDomainStreamInfoRequest struct{}"
	}

	return strings.Join([]string{"ListPlayDomainStreamInfoRequest", string(data)}, " ")
}
