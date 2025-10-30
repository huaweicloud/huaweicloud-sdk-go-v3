package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWatermarkRuleRequest Request Object
type ListWatermarkRuleRequest struct {

	// 水印模板ID
	TemplateId *string `json:"template_id,omitempty"`

	// 推流域名
	Domain *string `json:"domain,omitempty"`

	// 推流appname
	App *string `json:"app,omitempty"`

	// OTT场景，频道ID
	ChannelId *string `json:"channel_id,omitempty"`

	// OTT场景，填转码模板ID，云直播填流名
	Stream *string `json:"stream,omitempty"`

	// 偏移量，表示从此偏移量开始查询，offset大于等于0
	Offset *int32 `json:"offset,omitempty"`

	// 每页记录数，取值范围[1,100]，默认值10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListWatermarkRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWatermarkRuleRequest struct{}"
	}

	return strings.Join([]string{"ListWatermarkRuleRequest", string(data)}, " ")
}
