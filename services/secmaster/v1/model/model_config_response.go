package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigResponse struct {

	// 自动转告警的开关
	Alert *bool `json:"alert,omitempty"`

	// 能否开自动转告警
	AllowAlert *bool `json:"allow_alert,omitempty"`

	// 是否允许长期存储
	AllowLts *bool `json:"allow_lts,omitempty"`

	// 云服务描述
	CsvcDisplay *string `json:"csvc_display,omitempty"`

	// 关联的数据集列表
	Datasets *[]DatasetInfo `json:"datasets,omitempty"`

	// 是否按区域分片采集
	Region *bool `json:"region,omitempty"`

	// 数据源描述
	SourceDisplay *string `json:"source_display,omitempty"`

	// 是否创建成功
	Success *bool `json:"success,omitempty"`

	// 当前已采集的日志条数
	Total *int32 `json:"total,omitempty"`
}

func (o ConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigResponse struct{}"
	}

	return strings.Join([]string{"ConfigResponse", string(data)}, " ")
}
