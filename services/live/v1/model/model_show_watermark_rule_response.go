package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWatermarkRuleResponse Response Object
type ShowWatermarkRuleResponse struct {

	// 水印规则名称，如果不填会使用默认名称。默认名称的构造规则为“域名:应用名:流名”，示例“example.com:live:stream”。
	RuleName *string `json:"rule_name,omitempty"`

	// 水印模板ID，只包含数字字母中划线，创建模板时生成
	TemplateId *string `json:"template_id,omitempty"`

	// 域名
	Domain *string `json:"domain,omitempty"`

	// APP名。须知：云直播场景是可选配置，媒体直播场景为必选配置。
	App *string `json:"app,omitempty"`

	// 流名OTT场景下，可以不填
	Stream *string `json:"stream,omitempty"`

	Location *WatermarkLocation `json:"location,omitempty"`

	// OTT场景使用，填对应频道的频ID
	ChannelId *string `json:"channel_id,omitempty"`

	// OTT场景时，填频道对应的转码模板名称
	TranscodeTemplateName *string `json:"transcode_template_name,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowWatermarkRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWatermarkRuleResponse struct{}"
	}

	return strings.Join([]string{"ShowWatermarkRuleResponse", string(data)}, " ")
}
