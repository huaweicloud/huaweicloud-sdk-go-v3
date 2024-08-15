package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RedirectPoolsStickySessionConfig 参数解释：配置转发策略关联的服务器组之间会话保持。负载均衡器会根据客户端第一个请求生成一个cookie，后续所有包含这个cookie值的请求都会由同一个pool处理。  [约束限制：共享型负载均衡器下的转发策略不支持该字段，传入会报错。](tag:hws,hws_hk,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb,fcs,dt)  [不支持该字段，请勿使用。](tag:hcso_dt)  [荷兰region不支持该字段，请勿使用。](tag:dt)
type RedirectPoolsStickySessionConfig struct {

	// 参数解释：转发策略主机组会话保持开启的开关。  默认取值：false，表示关闭主机组会话保持。
	Enable *bool `json:"enable,omitempty"`

	// 参数解释：会话保持的时间。  取值范围：1-1440（分钟）  默认取值：1440  [荷兰region不支持QUIC。](tag:dt,dt_test)
	Timeout *int32 `json:"timeout,omitempty"`
}

func (o RedirectPoolsStickySessionConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedirectPoolsStickySessionConfig struct{}"
	}

	return strings.Join([]string{"RedirectPoolsStickySessionConfig", string(data)}, " ")
}
