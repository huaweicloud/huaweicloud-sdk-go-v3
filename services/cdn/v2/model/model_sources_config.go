package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SourcesConfig **参数解释：** 配置加速域名的源站，CDN通过加速域名将源站上需要加速的内容缓存到CDN节点，达到加速访问资源的效果 **约束限制：** 不涉及
type SourcesConfig struct {

	// **参数解释：** 源站类型 **约束限制：** 不涉及 **取值范围：** - ipaddr: 源站IP - domain: 源站域名 - obs_bucket: OBS桶域名 - third_bucket: 第三方桶，使用该源站类型需要提交工单申请  **默认取值：** 不涉及
	OriginType string `json:"origin_type"`

	// **参数解释：** 源站IP或者域名 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	OriginAddr string `json:"origin_addr"`

	// **参数解释：** 源站优先级 **约束限制：** 不涉及 **取值范围：** 1-1000 **默认取值：** 不涉及
	Priority int32 `json:"priority"`

	// **参数解释：** 权重 **约束限制：** 不涉及 **取值范围：** 1-100 **默认取值：** 不涉及
	Weight *int32 `json:"weight,omitempty"`

	// **参数解释：** 是否开启静态网站托管功能，如果您的源站是OBS桶域名且配置了静态网站托管，开启该功能后，客户端即可访问静态网站托管页面 > 开启该功能后将自动给桶源站地址加上website字段，如果地址中已有website字段，该参数不生效 **约束限制：** - 源站类型为obs_bucket时需要传该参数 - 只有当源站OBS桶已经开启静态网站托管时，才可开启，否则将导致业务异常  **取值范围：** - on: 开启 - off: 关闭  **默认取值：** 不涉及
	ObsWebHostingStatus *string `json:"obs_web_hosting_status,omitempty"`

	// **参数解释：** HTTP端口 **约束限制：** 不涉及 **取值范围：** 1-65535 **默认取值：** 80
	HttpPort *int32 `json:"http_port,omitempty"`

	// **参数解释：** HTTPS端口 **约束限制：** 不涉及 **取值范围：** 1-65535 **默认取值：** 443
	HttpsPort *int32 `json:"https_port,omitempty"`

	// **参数解释：** 回源HOST **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 加速域名
	HostName *string `json:"host_name,omitempty"`

	// **参数解释：** OBS回源鉴权开关，适用于OBS桶源站场景，如果桶需要带鉴权访问，则需要打开OBS回源鉴权开关。如此，CDN节点回源时会携带鉴权信息，鉴权信息与OBS桶不匹配时，回源失败，避免非授权的请求盗刷流量 **约束限制：** 源站类型是“OBS桶域名”时需要传该参数 **取值范围：** - private: 开启回源鉴权开关 - public: 关闭回源鉴权开关  **默认取值：** public: 关闭回源鉴权开关
	ObsBucketType *string `json:"obs_bucket_type,omitempty"`

	// **参数解释：** 第三方对象存储访问密钥ID，即AK **约束限制：** 源站类型为第三方桶时必填 **取值范围：** 不涉及 **默认取值：** 不涉及
	BucketAccessKey *string `json:"bucket_access_key,omitempty"`

	// **参数解释：** 第三方对象存储私密访问密钥，即SK **约束限制：** 源站类型为第三方桶时必填 **取值范围：** 不涉及 **默认取值：** 不涉及
	BucketSecretKey *string `json:"bucket_secret_key,omitempty"`

	// **参数解释：** 第三方对象存储区域 **约束限制：** 源站类型为第三方桶时必填 **取值范围：** 不涉及 **默认取值：** 不涉及
	BucketRegion *string `json:"bucket_region,omitempty"`

	// **参数解释：** 第三方对象存储名称 **约束限制：** 源站类型为第三方桶时必填 **取值范围：** 不涉及 **默认取值：** 不涉及
	BucketName *string `json:"bucket_name,omitempty"`
}

func (o SourcesConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourcesConfig struct{}"
	}

	return strings.Join([]string{"SourcesConfig", string(data)}, " ")
}
