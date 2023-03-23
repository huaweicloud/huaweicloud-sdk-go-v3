package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type DeleteCcRuleResponse struct {

	// Rule ID.
	Id *string `json:"id,omitempty"`

	// Policy ID.
	Policyid *string `json:"policyid,omitempty"`

	// 当mode值为0时，该参数有返回值。规则应用的URL链接，不包含域名。
	Url *string `json:"url,omitempty"`

	// 路径是否为前缀模式，当防护url以*结束，则为前缀模式。
	Prefix *bool `json:"prefix,omitempty"`

	// cc规则防护模式，对应console上的mode，现在只支持创建高级cc规则防护模式。   - 0：标准，只支持对域名的防护路径做限制。  - 1：高级，支持对路径、IP、Cookie、Header、Params字段做限制。
	Mode float32 `json:"mode,omitempty"`

	// 规则状态，0：关闭，1：开启
	Status *int32 `json:"status,omitempty"`

	// cc规则防护规则限速条件，当cc防护规则为高级模式（mode参数值为1）时，该参数必填。
	Conditions *[]CcCondition `json:"conditions,omitempty"`

	Action *CcrulesListInfoAction `json:"action,omitempty"`

	// 限速模式：   - ip：IP限速，根据IP区分单个Web访问者。   - cookie：用户限速，根据Cookie键值区分单个Web访问者。   - header：用户限速，根据Header区分单个Web访问者。   - other：根据Referer（自定义请求访问的来源）字段区分单个Web访问者。   - policy: 策略限速   - domain: 域名限速     - url: url限速
	TagType *DeleteCcRuleResponseTagType `json:"tag_type,omitempty"`

	// 用户标识，当限速模式为用户限速(cookie或header)时，需要传该参数。   - 选择cookie时，设置cookie字段名，即用户需要根据网站实际情况配置唯一可识别Web访问者的cookie中的某属性变量名。用户标识的cookie，不支持正则，必须完全匹配。例如：如果网站使用cookie中的某个字段name唯一标识用户，那么可以选取name字段来区分Web访问者。   - 选择header时，设置需要防护的自定义HTTP首部，即用户需要根据网站实际情况配置可识别Web访问者的HTTP首部。
	TagIndex *string `json:"tag_index,omitempty"`

	TagCondition *CcrulesListInfoTagCondition `json:"tag_condition,omitempty"`

	// 限制频率，单位为次，范围为1~2147483647
	LimitNum *int32 `json:"limit_num,omitempty"`

	// 限速周期，单位为秒，范围1~3600
	LimitPeriod *int32 `json:"limit_period,omitempty"`

	// 放行频率，单位为次，范围为0~2147483647。只有当防护动作类型为dynamic_block时，才需要传该参数。
	UnlockNum *int32 `json:"unlock_num,omitempty"`

	// 阻断时间，单位为秒，范围为0~65535。当“防护动作”选择“阻断”时，可设置阻断后恢复正常访问页面的时间。
	LockTime *int32 `json:"lock_time,omitempty"`

	// 是否开启域名聚合统计。
	DomainAggregation *bool `json:"domain_aggregation,omitempty"`

	// 是否开启全局计数。
	RegionAggregation *bool `json:"region_aggregation,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// 该参数为预留参数，用于后续功能扩展，当前请用户忽略该参数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 该参数为预留参数，用于后续功能扩展，当前请用户忽略该参数
	Unaggregation *bool `json:"unaggregation,omitempty"`

	// 规则老化时间，该参数为预留参数，用于后续功能扩展，当前请用户忽略该参数
	AgingTime *int32 `json:"aging_time,omitempty"`

	// 规则创建对象，该参数为预留参数，用于后续功能扩展，当前请用户忽略该参数
	Producer *int32 `json:"producer,omitempty"`

	// 创建规则时间戳
	Timestamp      *int64 `json:"timestamp,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o DeleteCcRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCcRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteCcRuleResponse", string(data)}, " ")
}

type DeleteCcRuleResponseTagType struct {
	value string
}

type DeleteCcRuleResponseTagTypeEnum struct {
	IP     DeleteCcRuleResponseTagType
	COOKIE DeleteCcRuleResponseTagType
	HEADER DeleteCcRuleResponseTagType
	OTHER  DeleteCcRuleResponseTagType
	POLICY DeleteCcRuleResponseTagType
	DOMAIN DeleteCcRuleResponseTagType
	URL    DeleteCcRuleResponseTagType
}

func GetDeleteCcRuleResponseTagTypeEnum() DeleteCcRuleResponseTagTypeEnum {
	return DeleteCcRuleResponseTagTypeEnum{
		IP: DeleteCcRuleResponseTagType{
			value: "ip",
		},
		COOKIE: DeleteCcRuleResponseTagType{
			value: "cookie",
		},
		HEADER: DeleteCcRuleResponseTagType{
			value: "header",
		},
		OTHER: DeleteCcRuleResponseTagType{
			value: "other",
		},
		POLICY: DeleteCcRuleResponseTagType{
			value: "policy",
		},
		DOMAIN: DeleteCcRuleResponseTagType{
			value: "domain",
		},
		URL: DeleteCcRuleResponseTagType{
			value: "url",
		},
	}
}

func (c DeleteCcRuleResponseTagType) Value() string {
	return c.value
}

func (c DeleteCcRuleResponseTagType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteCcRuleResponseTagType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
