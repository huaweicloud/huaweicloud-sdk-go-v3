package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// L7Rule L7转发规则
type L7Rule struct {

	// 参数解释：转发规则的管理状，固定为true。  不支持该字段，请勿使用。
	AdminStateUp bool `json:"admin_state_up"`

	// 参数解释：转发规则的匹配方式。  取值范围：type为HOST_NAME时可以为EQUAL_TO。type为PATH时可以为REGEX，STARTS_WITH，EQUAL_TO。
	CompareType string `json:"compare_type"`

	// 参数解释：匹配内容的键值。  [约束限制：type为HOST_NAME和PATH时，该字段不生效。 ](tag:hws,hws_hk,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,fcs,dt)  [不支持该字段，请勿使用。](tag:hcso_dt)
	Key string `json:"key"`

	// 参数解释：转发规则所在的项目ID。
	ProjectId string `json:"project_id"`

	// 参数解释：转发规则类别。  约束限制： - 一个l7policy下创建的l7rule的HOST_NAME，PATH，METHOD，SOURCE_IP不能重复。 HEADER、QUERY_STRING支持重复的rule配置。  取值范围： - HOST_NAME：匹配域名。 - PATH：匹配请求路径。 - METHOD：匹配请求方法。 - HEADER：匹配请求头。 - QUERY_STRING：匹配请求查询参数。 - SOURCE_IP：匹配请求源IP地址。 - COOKIE: 匹配cookie信息。  [只支持取值为HOST_NAME，PATH。](tag:hcso_dt)
	Type L7RuleType `json:"type"`

	// 参数解释：匹配内容的值。  约束限制：仅当conditions空时该字段生效。  取值范围： - 当type为HOST_NAME时，字符串只能包含英文字母、数字、-.*，必须以字母、数字或*开头。若域名中包含*，则*只能出现在开头且必须以*.开始。当*开头时表示通配0~任一个字符。  - 当type为PATH时，当转发规则的compare_type为STARTS_WITH、EQUAL_TO时，字符串只能包含英文字母、数字、_~';@^-%#&$.*+?,=!:|\\/()\\[\\]{}，且必须以/开头。  - 当type为METHOD、SOURCE_IP、HEADER, QUERY_STRING时，该字段无意义，使用condition_pair来指定key，value。
	Value string `json:"value"`

	// 参数解释：provisioning状态。该字段无效，默认为ACTIVE。  取值范围：ACTIVE、PENDING_CREATE 或者ERROR。
	ProvisioningStatus string `json:"provisioning_status"`

	// 参数解释：是否反向匹配。  约束限制：固定为false。该字段能更新但不会生效。
	Invert bool `json:"invert"`

	// 参数解释：规则ID。
	Id string `json:"id"`

	// 参数解释：转发规则的匹配条件。  约束限制： - 当监听器的高级转发策略功能（enhance_l7policy_enable）开启后才会生效。 - 若转发规则配置了conditions，字段key、字段value的值无意义。 - 同一个rule内的conditions列表中所有key必须相同，value不允许重复。  [不支持该字段，请勿使用。](tag:hcso_dt)  [荷兰region不支持该字段，请勿使用。](tag:dt,dt_test)
	Conditions []RuleCondition `json:"conditions"`

	// 参数解释：创建时间。  取值范围：格式：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。  [注意：独享型实例的历史数据以及共享型实例下的资源，不返回该字段。 ](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,fcs,dt,hk_tm)
	CreatedAt *string `json:"created_at,omitempty"`

	// 参数解释：更新时间。  取值范围：格式：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。  [注意：独享型实例的历史数据以及共享型实例下的资源，不返回该字段。 ](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,fcs,dt,hk_tm)
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o L7Rule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "L7Rule struct{}"
	}

	return strings.Join([]string{"L7Rule", string(data)}, " ")
}

type L7RuleType struct {
	value string
}

type L7RuleTypeEnum struct {
	HOST_NAME    L7RuleType
	PATH         L7RuleType
	METHOD       L7RuleType
	HEADER       L7RuleType
	QUERY_STRING L7RuleType
	SOURCE_IP    L7RuleType
}

func GetL7RuleTypeEnum() L7RuleTypeEnum {
	return L7RuleTypeEnum{
		HOST_NAME: L7RuleType{
			value: "HOST_NAME",
		},
		PATH: L7RuleType{
			value: "PATH",
		},
		METHOD: L7RuleType{
			value: "METHOD",
		},
		HEADER: L7RuleType{
			value: "HEADER",
		},
		QUERY_STRING: L7RuleType{
			value: "QUERY_STRING",
		},
		SOURCE_IP: L7RuleType{
			value: "SOURCE_IP",
		},
	}
}

func (c L7RuleType) Value() string {
	return c.value
}

func (c L7RuleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *L7RuleType) UnmarshalJSON(b []byte) error {
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
