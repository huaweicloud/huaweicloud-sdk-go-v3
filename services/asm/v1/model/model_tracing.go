package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Tracing struct {

	// tracing采样率
	RandomSamplingPercentage *float32 `json:"randomSamplingPercentage,omitempty"`

	// tracing默认上报的provider名称，必须与extensionProviders中的name字段匹配，或使用ASM预设的provider \"apm-otel\"。 如果使用\"apm-otel\"，需确认当前region已支持APM2.0且网格版本大于1.18。
	DefaultProviders *[]string `json:"defaultProviders,omitempty"`

	// 用户自配置provider，目前支持zipkin协议。 如果用户配置zipkin协议的provider，请保证网格版本大于等于1.15。
	ExtensionProviders *[]TracingExtensionProvider `json:"extensionProviders,omitempty"`
}

func (o Tracing) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tracing struct{}"
	}

	return strings.Join([]string{"Tracing", string(data)}, " ")
}
