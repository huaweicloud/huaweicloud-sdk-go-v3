package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SourceWithPort 域名信息。
type SourceWithPort struct {

	// **参数解释：** 加速域名id。 **约束限制：** 不涉及 **取值范围：** 不涉及
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释：** 源站IP（非内网IP）或者域名。 **约束限制：** 不涉及 **取值范围：** 不涉及
	IpOrDomain string `json:"ip_or_domain"`

	// 源站类型，ipaddr：源站IP、 domain：源站域名、obs_bucket：OBS桶域名。
	OriginType SourceWithPortOriginType `json:"origin_type"`

	// OBS桶类型。   - private: 私有桶（除桶ACL授权外的其他用户无桶的访问权限）。   - public: 公有桶（任何用户都可以对桶内对象进行读操作）。
	ObsBucketType *string `json:"obs_bucket_type,omitempty"`

	// 主备状态（1代表主源站；0代表备源站）。
	ActiveStandby int32 `json:"active_standby"`

	// **参数解释：** 是否开启静态网站托管功能，如果您的源站是OBS桶域名且配置了静态网站托管，开启该功能后，客户端即可访问静态网站托管页面 > 开启该功能后将自动给桶源站地址加上website字段，如果地址中已有website字段，该参数不生效 **约束限制：** - 源站类型为obs_bucket时需要传该参数 - 只有当源站OBS桶已经开启静态网站托管时，才可开启，否则将导致业务异常  **取值范围：** - 0: 关闭 - 1: 开启  **默认取值：** 0: 关闭
	EnableObsWebHosting *int32 `json:"enable_obs_web_hosting,omitempty"`

	// **参数解释：** HTTP端口，默认80 **约束限制：** 不涉及 **取值范围：** 不涉及
	HttpPort *int32 `json:"http_port,omitempty"`

	// **参数解释：** HTTPS端口，默认443 **约束限制：** 不涉及 **取值范围：** 不涉及
	HttpsPort *int32 `json:"https_port,omitempty"`

	// **参数解释：** 权重，值越大回源到该源站的次数越多。多个优先级相同的源站，由权重决定回源到各个源站的比例 **约束限制：** 不涉及 **取值范围：** 1-100 **默认取值：** 50
	Weight *int32 `json:"weight,omitempty"`
}

func (o SourceWithPort) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourceWithPort struct{}"
	}

	return strings.Join([]string{"SourceWithPort", string(data)}, " ")
}

type SourceWithPortOriginType struct {
	value string
}

type SourceWithPortOriginTypeEnum struct {
	IPADDR     SourceWithPortOriginType
	DOMAIN     SourceWithPortOriginType
	OBS_BUCKET SourceWithPortOriginType
}

func GetSourceWithPortOriginTypeEnum() SourceWithPortOriginTypeEnum {
	return SourceWithPortOriginTypeEnum{
		IPADDR: SourceWithPortOriginType{
			value: "ipaddr",
		},
		DOMAIN: SourceWithPortOriginType{
			value: "domain",
		},
		OBS_BUCKET: SourceWithPortOriginType{
			value: "obs_bucket",
		},
	}
}

func (c SourceWithPortOriginType) Value() string {
	return c.value
}

func (c SourceWithPortOriginType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceWithPortOriginType) UnmarshalJSON(b []byte) error {
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
