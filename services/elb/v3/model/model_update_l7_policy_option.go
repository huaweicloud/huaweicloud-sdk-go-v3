package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateL7PolicyOption 更新七层转发策略的请求参数。
type UpdateL7PolicyOption struct {

	// 参数解释：转发策略的管理状态。  约束限制：只支持设置为true。
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 参数解释：转发策略描述信息。
	Description *string `json:"description,omitempty"`

	// 参数解释：转发策略名称。
	Name *string `json:"name,omitempty"`

	// 参数解释：转发到的listener的ID。  约束限制： - 当action为REDIRECT_TO_LISTENER时不能更新为空或null。 - 只支持protocol为HTTPS/TERMINATED_HTTPS的listener。 - 不能指定为其他loadbalancer下的listener。 - 当action为REDIRECT_TO_POOL时，创建或更新时不能传入该参数。
	RedirectListenerId *string `json:"redirect_listener_id,omitempty"`

	// 参数解释：转发到pool的ID。  约束限制： - 指定的pool不能是任何listener的default_pool。不能是其他listener的l7policy使用的pool。 - 当action为REDIRECT_TO_POOL时生效，但不能更新为空或null。 - 当action为REDIRECT_TO_LISTENER时，传入会报错。
	RedirectPoolId *string `json:"redirect_pool_id,omitempty"`

	// 参数解释：转发到多个主机组列表。一个policy最多配置5个pool。
	RedirectPoolsConfig *[]UpdateRedirectPoolsConfig `json:"redirect_pools_config,omitempty"`

	RedirectPoolsStickySessionConfig *UpdateRedirectPoolsStickySessionConfig `json:"redirect_pools_sticky_session_config,omitempty"`

	RedirectUrlConfig *UpdateRedirectUrlConfig `json:"redirect_url_config,omitempty"`

	FixedResponseConfig *UpdateFixtedResponseConfig `json:"fixed_response_config,omitempty"`

	RedirectPoolsExtendConfig *UpdateRedirectPoolsExtendConfig `json:"redirect_pools_extend_config,omitempty"`

	// 参数解释：转发策略关联的转发规则对象。  约束限制： - rules列表中最多含有10个rule规则 （若rule中包含conditions字段，一条condition算一个规则）， 且列表中type为HOST_NAME，PATH，METHOD，SOURCE_IP的rule不能重复，至多指定一条。 - 仅支持全量替换。
	Rules *[]CreateRuleOption `json:"rules,omitempty"`

	// 参数解释：转发策略的优先级。数字越小表示优先级越高。  约束限制： - 同一个监听器下不同转发策略之间不允许重复的优先级数值。 - 当关联的监听器的高级转发策略功能（enhance_l7policy_enable）开启后才会生效，未开启传入该字段会报错。 - 当关联的监听器的高级转发策略功能（enhance_l7policy_enable）未开启，按原有policy的排序逻辑，自动排序。 不同域名优先级独立。相同域名下，按path的compare_type排序， 精确>前缀>正则，匹配类型相同时，path的长度越长优先级越高。 若policy下只有域名rule，没有路径rule，默认path为前缀匹配/。 [- 共享型负载均衡器下的转发策略不支持该字段。](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,fcs,dt,hk_tm)  取值范围： - 当action为REDIRECT_TO_LISTENER时，支持指定为0-10000。 - 其它action取值，支持指定为1-10000。  默认取值： - 若关联的监听器的高级转发策略功能（enhance_l7policy_enable）未开启，且不传入该字段，则新创建的转发策略的优先级的值为1。 - 当action为REDIRECT_TO_LISTENER时，则新创建的转发策略的优先级的值为0。 - 其它action取值，新创建的转发策略的优先级的值为同一监听器下已有转发策略的优先级的最大值+1。   + 若监听器下没有转发策略，则新建的转发策略的优先级为1。   + 若当前已有转发策略的优先级的最大值是10000，则新创建的转发策略会因超出取值范围10000而失败。此时可通过传入指定priority，或调整原有policy的优先级来避免错误。  [不支持该字段，请勿使用。](tag:hcso_dt)  [荷兰region不支持该字段，请勿使用。](tag:dt,dt_test)
	Priority *int32 `json:"priority,omitempty"`
}

func (o UpdateL7PolicyOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateL7PolicyOption struct{}"
	}

	return strings.Join([]string{"UpdateL7PolicyOption", string(data)}, " ")
}
