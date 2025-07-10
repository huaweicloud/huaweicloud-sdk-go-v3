package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlexibleOriginsEngine **参数解释：** 高级回源，实现根据不同的资源类型或路径回源到不同源站 **约束限制：** 最多配置20条
type FlexibleOriginsEngine struct {

	// **参数解释：** 源站类型 **约束限制：** 不涉及 **取值范围：** - ipaddr: 源站IP - domain: 源站域名 - obs_bucket: OBS桶域名 - third_bucket: 第三方桶域名 **默认取值：** 不涉及
	SourcesType string `json:"sources_type"`

	// **参数解释：** 源站IP或者域名 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	IpOrDomain string `json:"ip_or_domain"`

	// **参数解释：** OBS桶类型 **约束限制：** 源站类型是“OBS桶域名”时需要传该参数 **取值范围：** - private: 私有桶 - public: 公有桶 **默认取值：** public: 公有桶
	ObsBucketType *string `json:"obs_bucket_type,omitempty"`

	// **参数解释：** 第三方对象存储访问密钥 **约束限制：** 源站类型为第三方桶时必填 **取值范围：** 不涉及 **默认取值：** 不涉及
	BucketAccessKey *string `json:"bucket_access_key,omitempty"`

	// **参数解释：** 第三方对象存储密钥 **约束限制：** 源站类型为第三方桶时必填 **取值范围：** 不涉及 **默认取值：** 不涉及
	BucketSecretKey *string `json:"bucket_secret_key,omitempty"`

	// **参数解释：** 第三方对象存储区域 **约束限制：** 源站类型为第三方桶时必填 **取值范围：** 不涉及 **默认取值：** 不涉及
	BucketRegion *string `json:"bucket_region,omitempty"`

	// **参数解释：** 第三方对象存储名称 **约束限制：** 源站类型为第三方桶时必填 **取值范围：** 不涉及 **默认取值：** 不涉及
	BucketName *string `json:"bucket_name,omitempty"`

	// **参数解释：** 回源HOST **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 加速域名
	HostName *string `json:"host_name,omitempty"`

	// **参数解释：** 指定CDN回源时的请求协议 **约束限制：** 不涉及 **取值范围：** - follow: 协议跟随 - http: http协议 - https: https协议 **默认取值：** http: http协议
	OriginProtocol *string `json:"origin_protocol,omitempty"`

	// **参数解释：** HTTP端口 **约束限制：** 不涉及 **取值范围：** 1-65535 **默认取值：** 80
	HttpPort *int32 `json:"http_port,omitempty"`

	// **参数解释：** HTTPS端口 **约束限制：** 不涉及 **取值范围：** 1-65535 **默认取值：** 443
	HttpsPort *int32 `json:"https_port,omitempty"`

	// **参数解释：** 优先级，值越大优先级越高 **约束限制：** 不涉及 **取值范围：** 1-100 **默认取值：** 不涉及
	Priority int32 `json:"priority"`

	// **参数解释：** 权重，值越大回源到该源站的次数越多。多个优先级相同的源站，由权重决定回源到各个源站的比例 **约束限制：** 不涉及 **取值范围：** 1-100 **默认取值：** 不涉及
	Weight int32 `json:"weight"`
}

func (o FlexibleOriginsEngine) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlexibleOriginsEngine struct{}"
	}

	return strings.Join([]string{"FlexibleOriginsEngine", string(data)}, " ")
}
