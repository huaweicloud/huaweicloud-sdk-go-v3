package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SourcesRequestBody 源站信息
type SourcesRequestBody struct {

	// **参数解释：** 加速域名ID，系统自动生成，创建域名时无需传入此参数 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释：** 源站IP或者域名 **约束限制：** 不能配置内网IP **取值范围：** 不涉及 **默认取值：** 不涉及
	IpOrDomain string `json:"ip_or_domain"`

	// **参数解释：** 源站类型 **约束限制：** 不涉及 **取值范围：** - ipaddr: 源站IP - domain: 源站域名 - obs_bucket: OBS桶域名  **默认取值：** 不涉及
	OriginType SourcesRequestBodyOriginType `json:"origin_type"`

	// **参数解释：** OBS回源鉴权开关，适用于OBS桶源站场景，如果桶需要带鉴权访问，则需要打开OBS回源鉴权开关。如此，CDN节点回源时会携带鉴权信息，鉴权信息与OBS桶不匹配时，回源失败，避免非授权的请求盗刷流量 **约束限制：** 源站类型是“OBS桶域名”时需要传该参数 **取值范围：** - private: 开启回源鉴权开关 - public: 关闭回源鉴权开关  **默认取值：** public: 关闭回源鉴权开关
	ObsBucketType *string `json:"obs_bucket_type,omitempty"`

	// **参数解释：** 设置主/备源站，主源站优先级高于备源站，CDN节点回源时优先回到主源站 >- 主源站优先级默认为70，备源站优先级默认为30。  **约束限制：** 与priority二者至少传一个，同时存在时以priority为准 **取值范围：** - 1: 主源站 - 0: 备源站  **默认取值：** 不涉及
	ActiveStandby *int32 `json:"active_standby,omitempty"`

	// **参数解释：** 源站优先级，数值越大优先级越高，创建域名时可用于替代active_standby指定源站优先级 **约束限制：** - 与active_standby二者至少传一个，同时存在时以priority为准 - 不同优先级数量不超过6个  **取值范围：** 1-1000 **默认取值：** 不涉及
	Priority *int32 `json:"priority,omitempty"`

	// **参数解释：** 权重，值越大回源到该源站的次数越多。多个优先级相同的源站，由权重决定回源到各个源站的比例 **约束限制：** 不涉及 **取值范围：** 1-100 **默认取值：** 50
	Weight *int32 `json:"weight,omitempty"`

	// **参数解释：** 回源HOST **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 加速域名
	HostName *string `json:"host_name,omitempty"`

	// **参数解释：** 是否开启静态网站托管功能，如果您的源站是OBS桶域名且配置了静态网站托管，开启该功能后，客户端即可访问静态网站托管页面 > 开启该功能后将自动给桶源站地址加上website字段，如果地址中已有website字段，该参数不生效 **约束限制：** - 源站类型为obs_bucket时需要传该参数 - 只有当源站OBS桶已经开启静态网站托管时，才可开启，否则将导致业务异常  **取值范围：** - 0: 关闭 - 1: 开启  **默认取值：** 0: 关闭
	EnableObsWebHosting *int32 `json:"enable_obs_web_hosting,omitempty"`

	// **参数解释：** HTTP端口 **约束限制：** 不涉及 **取值范围：** 1-65535 **默认取值：** 80
	HttpPort *int32 `json:"http_port,omitempty"`

	// **参数解释：** HTTPS端口 **约束限制：** 不涉及 **取值范围：** 1-65535 **默认取值：** 443
	HttpsPort *int32 `json:"https_port,omitempty"`
}

func (o SourcesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourcesRequestBody struct{}"
	}

	return strings.Join([]string{"SourcesRequestBody", string(data)}, " ")
}

type SourcesRequestBodyOriginType struct {
	value string
}

type SourcesRequestBodyOriginTypeEnum struct {
	IPADDR     SourcesRequestBodyOriginType
	DOMAIN     SourcesRequestBodyOriginType
	OBS_BUCKET SourcesRequestBodyOriginType
}

func GetSourcesRequestBodyOriginTypeEnum() SourcesRequestBodyOriginTypeEnum {
	return SourcesRequestBodyOriginTypeEnum{
		IPADDR: SourcesRequestBodyOriginType{
			value: "ipaddr",
		},
		DOMAIN: SourcesRequestBodyOriginType{
			value: "domain",
		},
		OBS_BUCKET: SourcesRequestBodyOriginType{
			value: "obs_bucket",
		},
	}
}

func (c SourcesRequestBodyOriginType) Value() string {
	return c.value
}

func (c SourcesRequestBodyOriginType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourcesRequestBodyOriginType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
