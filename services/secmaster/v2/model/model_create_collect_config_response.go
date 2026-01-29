package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCollectConfigResponse Response Object
type CreateCollectConfigResponse struct {

	// 自动转告警的开关
	Alert *bool `json:"alert,omitempty"`

	// 能否开自动转告警
	AllowAlert *bool `json:"allow_alert,omitempty"`

	// 是否支持lts
	AllowLts *bool `json:"allow_lts,omitempty"`

	// 云服务描述
	CsvcDisplay *string `json:"csvc_display,omitempty"`

	// 局点
	Region *string `json:"region,omitempty"`

	// 数据源描述
	SourceDisplay *string `json:"source_display,omitempty"`

	// 是否创建成功
	Success *bool `json:"success,omitempty"`

	// 条数
	Total *int32 `json:"total,omitempty"`

	// 关联的数据集列表
	Datasets       *[]DatasetInfo `json:"datasets,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o CreateCollectConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCollectConfigResponse struct{}"
	}

	return strings.Join([]string{"CreateCollectConfigResponse", string(data)}, " ")
}
