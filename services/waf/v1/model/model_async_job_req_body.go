package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AsyncJobReqBody **参数解释：** 下发自定义导出攻击事件的异步任务的请求体，用于配置攻击事件日志的导出条件和字段 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type AsyncJobReqBody struct {

	// **参数解释：** 查询日志的时间范围，如1week（1周）、1month（1个月） **约束限制：** 不涉及 **取值范围：** - yesterday - today - 3days - 1week - 1month  **默认取值：** 不涉及
	Recent string `json:"recent"`

	// **参数解释：** 查询指定事件ID列表的日志，仅导出包含这些ID的攻击事件 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Ids *[]string `json:"ids,omitempty"`

	// **参数解释：** 查询不包含事件ID列表的日志，排除这些ID对应的攻击事件 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Nids *[]string `json:"nids,omitempty"`

	// **参数解释：** 查询指定攻击类型列表的日志，仅导出这些类型的攻击事件（如SQL注入、XSS等） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Attacks *[]string `json:"attacks,omitempty"`

	// **参数解释：** 查询不包含攻击类型列表的日志，排除这些类型的攻击事件 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Nattacks *[]string `json:"nattacks,omitempty"`

	// **参数解释：** 查询指定命中的规则id列表的日志，仅导出命中这些规则的攻击事件 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Rules *[]string `json:"rules,omitempty"`

	// **参数解释：** 查询不包含命中的规则id列表的日志，排除命中这些规则的攻击事件 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Nrules *[]string `json:"nrules,omitempty"`

	// **参数解释：** 查询指定源ip(模糊查询)，导出包含该IP（模糊匹配）的攻击事件 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Sip *string `json:"sip,omitempty"`

	// **参数解释：** 查询指定攻击的url(模糊查询)，导出包含该URL（模糊匹配）的攻击事件 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Url *string `json:"url,omitempty"`

	// **参数解释：** 查询指定源ip列表的日志，仅导出这些IP对应的攻击事件 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Sips *[]string `json:"sips,omitempty"`

	// **参数解释：** 查询不包含源ip列表的日志，排除这些IP对应的攻击事件 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Nsips *[]string `json:"nsips,omitempty"`

	// **参数解释：** 查询指定攻击的url链接列表的日志，仅导出这些URL对应的攻击事件 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Urls *[]string `json:"urls,omitempty"`

	// **参数解释：** 查询不包含攻击的url链接列表的日志，排除这些URL对应的攻击事件 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Nurls *[]string `json:"nurls,omitempty"`

	// **参数解释：** 查询指定防护动作列表的日志，仅导出执行这些动作的攻击事件（如block-拦截、log-日志） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Actions *[]string `json:"actions,omitempty"`

	// **参数解释：** 查询不包含防护动作列表的日志，排除执行这些动作的攻击事件 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Nactions *[]string `json:"nactions,omitempty"`

	// **参数解释：** 查询指定host列表的日志，仅导出这些域名对应的攻击事件 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Hosts *[]string `json:"hosts,omitempty"`

	// **参数解释：** 查询指定instance列表的日志，仅导出这些实例对应的攻击事件 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Instances *[]string `json:"instances,omitempty"`

	// **参数解释：** 需要导出的列名称，指定攻击事件日志中需要导出的字段（如时间、事件ID、攻击类型等） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	DumpColumns []string `json:"dump_columns"`
}

func (o AsyncJobReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AsyncJobReqBody struct{}"
	}

	return strings.Join([]string{"AsyncJobReqBody", string(data)}, " ")
}
