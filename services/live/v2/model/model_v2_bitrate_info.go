package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 实时码率
type V2BitrateInfo struct {

	// 域名。
	PublishDomain *string `json:"publish_domain,omitempty" xml:"publish_domain"`

	// 应用名称。
	App *string `json:"app,omitempty" xml:"app"`

	// 流名。
	Stream *string `json:"stream,omitempty" xml:"stream"`

	// 采样开始时间。日期格式按照ISO8601表示法，并使用UTC时间。 格式为：YYYY-MM-DDThh:mm:ssZ。
	StartTime *string `json:"start_time,omitempty" xml:"start_time"`

	// 采样结束时间。日期格式按照ISO8601表示法，并使用UTC时间。 格式为：YYYY-MM-DDThh:mm:ssZ。
	EndTime *string `json:"end_time,omitempty" xml:"end_time"`

	// 码率信息列表，单位为Kbps。
	DataList *[]int64 `json:"data_list,omitempty" xml:"data_list"`
}

func (o V2BitrateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2BitrateInfo struct{}"
	}

	return strings.Join([]string{"V2BitrateInfo", string(data)}, " ")
}
