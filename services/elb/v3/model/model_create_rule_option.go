package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建L7转发规则器请求
type CreateRuleOption struct {
	// 转发规则的管理状态，默认为true。  不支持该字段，请勿使用。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 转发规则的匹配方式。取值： - type为HOST_NAME时仅支持EQUAL_TO，支持通配符*。 - type为PATH时可以为Perl类型的REGEX，STARTS_WITH，EQUAL_TO。 [- type为METHOD、SOURCE_IP时，仅支持EQUAL_TO。 - type为HEADER、QUERY_STRING，仅支持EQUAL_TO，支持通配符*、？。](tag:hc,hk,ocb,tlf,ctc,hcso,sbc,g42,tm,cmcc,hk-g42)

	CompareType string `json:"compare_type"`
	// 匹配项的名称，比如转发规则匹配类型是请求头匹配，则key表示请求头参数的名称。  不支持该字段，请勿使用。

	Key *string `json:"key,omitempty"`
	// 匹配项的值，比如转发规则匹配类型是域名匹配，则value表示域名的值。[仅当conditions空时该字段生效。](tag:hc,hk,ocb,tlf,ctc,hcso,sbc,g42,tm,cmcc,hk-g42)  当type为HOST_NAME时，字符串只能包含英文字母、数字、\"-\"、\".\"或\"*\"，必须以字母、数字或\"*\"开头。若域名中包含\"*\"，则\"*\"只能出现在开头且必须以*.开始。当*开头时表示通配0~任一个字符。  当type为PATH时，当转发规则的compare_type为STARTS_WITH、EQUAL_TO时，字符串只能包含英文字母、数字、_~';@^-%#&$.*+?,=!:|/()[]{}，且必须以\"/\"开头。  [当type为METHOD、SOURCE_IP、HEADER,QUERY_STRING时，该字段无意义，使用conditions来指定key/value。](tag:hc,hk,ocb,tlf,ctc,hcso,sbc,g42,tm,cmcc,hk-g42)

	Value string `json:"value"`
	// 转发规则所在的项目ID。

	ProjectId *string `json:"project_id,omitempty"`
	// 转发规则类别。取值： - HOST_NAME：匹配域名 - PATH：匹配请求路径 [- METHOD：匹配请求方法 - HEADER：匹配请求头 - QUERY_STRING：匹配请求查询参数 - SOURCE_IP：匹配请求源IP地址 使用说明：  - 一个l7policy下创建的l7rule的HOST_NAME，PATH，METHOD，SOURCE_IP不能重复。HEADER、QUERY_STRING支持重复的rule配置。](tag:hc,hk,ocb,tlf,ctc,hcso,sbc,g42,tm,cmcc,hk-g42) [使用说明：  - 一个l7policy下创建的l7rule的HOST_NAME，PATH不能重复。](tag:dt,dt_test,hcso_dt)

	Type string `json:"type"`
	// 是否反向匹配。取值：true、false，默认false。  不支持该字段，请勿使用。

	Invert *bool `json:"invert,omitempty"`
	// 转发规则的匹配条件。当监听器的高级转发策略功能（enhance_l7policy_enable）开启后才会生效。 配置了conditions后，字段key、字段value的值无意义。 若指定了conditions，该rule的条件数为conditions列表长度。 列表中key必须相同，value不允许重复。  [不支持该字段，请勿使用。](tag:dt,dt_test,hcso_dt)

	Conditions *[]CreateRuleCondition `json:"conditions,omitempty"`
}

func (o CreateRuleOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRuleOption struct{}"
	}

	return strings.Join([]string{"CreateRuleOption", string(data)}, " ")
}
