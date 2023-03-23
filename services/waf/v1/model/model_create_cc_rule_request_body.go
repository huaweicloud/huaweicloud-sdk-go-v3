package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateCcRuleRequestBody struct {

	// cc规则防护模式，对应console上的mode，现在只支持创建高级cc规则防护模式。   - 0：标准，只支持对域名的防护路径做限制。  - 1：高级，支持对路径、IP、Cookie、Header、Params字段做限制。
	Mode int32 `json:"mode"`

	// cc规则防护规则限速条件，当cc防护规则为高级模式（mode参数值为1）时，该参数必填。
	Conditions []CcCondition `json:"conditions"`

	Action *CreateCcRuleRequestBodyAction `json:"action"`

	// 限速模式：   - ip：IP限速，根据IP区分单个Web访问者。   - cookie：用户限速，根据Cookie键值区分单个Web访问者。   - header：用户限速，根据Header区分单个Web访问者。   - other：根据Referer（自定义请求访问的来源）字段区分单个Web访问者。   - policy: 策略限速   - domain: 域名限速     - url: url限速
	TagType CreateCcRuleRequestBodyTagType `json:"tag_type"`

	// 用户标识，当限速模式为用户限速(cookie或header)时，需要传该参数。   - 选择cookie时，设置cookie字段名，即用户需要根据网站实际情况配置唯一可识别Web访问者的cookie中的某属性变量名。用户标识的cookie，不支持正则，必须完全匹配。例如：如果网站使用cookie中的某个字段name唯一标识用户，那么可以选取name字段来区分Web访问者。   - 选择header时，设置需要防护的自定义HTTP首部，即用户需要根据网站实际情况配置可识别Web访问者的HTTP首部。
	TagIndex *string `json:"tag_index,omitempty"`

	TagCondition *CcrulesListInfoTagCondition `json:"tag_condition,omitempty"`

	// 限制频率，单位为次，范围为1~2147483647
	LimitNum int32 `json:"limit_num"`

	// 限速周期，单位为秒，范围1~3600
	LimitPeriod int32 `json:"limit_period"`

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
}

func (o CreateCcRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCcRuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCcRuleRequestBody", string(data)}, " ")
}

type CreateCcRuleRequestBodyTagType struct {
	value string
}

type CreateCcRuleRequestBodyTagTypeEnum struct {
	IP     CreateCcRuleRequestBodyTagType
	COOKIE CreateCcRuleRequestBodyTagType
	HEADER CreateCcRuleRequestBodyTagType
	OTHER  CreateCcRuleRequestBodyTagType
	POLICY CreateCcRuleRequestBodyTagType
	DOMAIN CreateCcRuleRequestBodyTagType
	URL    CreateCcRuleRequestBodyTagType
}

func GetCreateCcRuleRequestBodyTagTypeEnum() CreateCcRuleRequestBodyTagTypeEnum {
	return CreateCcRuleRequestBodyTagTypeEnum{
		IP: CreateCcRuleRequestBodyTagType{
			value: "ip",
		},
		COOKIE: CreateCcRuleRequestBodyTagType{
			value: "cookie",
		},
		HEADER: CreateCcRuleRequestBodyTagType{
			value: "header",
		},
		OTHER: CreateCcRuleRequestBodyTagType{
			value: "other",
		},
		POLICY: CreateCcRuleRequestBodyTagType{
			value: "policy",
		},
		DOMAIN: CreateCcRuleRequestBodyTagType{
			value: "domain",
		},
		URL: CreateCcRuleRequestBodyTagType{
			value: "url",
		},
	}
}

func (c CreateCcRuleRequestBodyTagType) Value() string {
	return c.value
}

func (c CreateCcRuleRequestBodyTagType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateCcRuleRequestBodyTagType) UnmarshalJSON(b []byte) error {
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
