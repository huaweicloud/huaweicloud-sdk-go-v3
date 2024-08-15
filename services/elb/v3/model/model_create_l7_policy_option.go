package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateL7PolicyOption 创建转发策略请求参数。
type CreateL7PolicyOption struct {

	// 参数解释：转发策略的转发动作。  约束限制： - REDIRECT_TO_LISTENER的优先级最高，配置了以后，该监听器下的其他policy会失效。 - 当action为REDIRECT_TO_POOL时， 只支持创建在PROTOCOL为HTTP、HTTPS、TERMINATED_HTTPS的listener上。 - 当action为REDIRECT_TO_LISTENER时，只支持创建在PROTOCOL为HTTP的listener上。  取值范围： - REDIRECT_TO_POOL：转发到后端云服务器组。 - REDIRECT_TO_LISTENER：重定向到监听器。 - REDIRECT_TO_URL：重定向到URL。 - FIXED_RESPONSE：返回固定响应体。  [不支持REDIRECT_TO_URL和FIXED_RESPONSE](tag:hcso_dt)
	Action string `json:"action"`

	// 参数解释：转发策略的管理状态。  约束限制：只支持设置为true。
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 参数解释：转发策略描述信息。
	Description *string `json:"description,omitempty"`

	// 参数解释：转发策略对应的监听器ID。  约束限制： - 当action为REDIRECT_TO_POOL时，只支持创建在PROTOCOL为HTTP或HTTPS的listener上。 - 当action为REDIRECT_TO_LISTENER时，只支持创建在PROTOCOL为HTTP的listener上。
	ListenerId string `json:"listener_id"`

	// 参数解释：转发策略名称。
	Name *string `json:"name,omitempty"`

	// 参数解释：转发策略的优先级。  约束限制：不支持更新。  不支持该字段，请勿使用。
	Position *int32 `json:"position,omitempty"`

	// 参数解释：转发策略的优先级。数字越小表示优先级越高。  约束限制： - 同一个监听器下不同转发策略之间不允许重复的优先级数值。 - 当关联的监听器的高级转发策略功能（enhance_l7policy_enable）开启后才会生效，未开启传入该字段会报错。 - 当关联的监听器的高级转发策略功能（enhance_l7policy_enable）未开启，按原有policy的排序逻辑，自动排序。 不同域名优先级独立。相同域名下，按path的compare_type排序， 精确>前缀>正则，匹配类型相同时，path的长度越长优先级越高。 若policy下只有域名rule，没有路径rule，默认path为前缀匹配/。  [- 共享型负载均衡器下的转发策略不支持该字段。](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,fcs,dt,hk_tm)  取值范围： - 当action为REDIRECT_TO_LISTENER时，支持指定为0-10000。 - 其它action取值，支持指定为1-10000。  默认取值： - 若关联的监听器的高级转发策略功能（enhance_l7policy_enable）未开启，且不传入该字段，则新创建的转发策略的优先级的值为1。 - 当action为REDIRECT_TO_LISTENER时，则新创建的转发策略的优先级的值为0。 - 其它action取值，新创建的转发策略的优先级的值为同一监听器下已有转发策略的优先级的最大值+1。   + 若监听器下没有转发策略，则新建的转发策略的优先级为1。   + 若当前已有转发策略的优先级的最大值是10000，则新创建的转发策略会因超出取值范围10000而失败。此时可通过传入指定priority，或调整原有policy的优先级来避免错误。  [不支持该字段，请勿使用。](tag:hcso_dt)  [荷兰region不支持该字段，请勿使用。](tag:dt,dt_test)
	Priority *int32 `json:"priority,omitempty"`

	// 参数解释：转发策略所在的项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 参数解释：转发到的listener的ID，当action为REDIRECT_TO_LISTENER时必选。  约束限制： - 只支持protocol为HTTPS/TERMINATED_HTTPS的listener。 - 不能指定为其他loadbalancer下的listener。 - 当action为REDIRECT_TO_POOL时，创建或更新时不能传入该参数。 [- 共享型负载均衡器下的转发策略不支持该字段。](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,fcs,dt,hk_tm)
	RedirectListenerId *string `json:"redirect_listener_id,omitempty"`

	// 参数解释：转发到pool的ID。  约束限制： - 当action为REDIRECT_TO_POOL时生效。 - 当action为REDIRECT_TO_LISTENER时，传入会报错。
	RedirectPoolId *string `json:"redirect_pool_id,omitempty"`

	// 参数解释：转发到的url。  约束限制：必须满足格式: protocol://host:port/path?query。  [不支持该字段，请勿使用。](tag:hcso_dt)
	RedirectUrl *string `json:"redirect_url,omitempty"`

	RedirectUrlConfig *CreateRedirectUrlConfig `json:"redirect_url_config,omitempty"`

	// 参数解释：转发到多个主机组列表。  约束限制：一个policy最多配置5个pool。
	RedirectPoolsConfig *[]CreateRedirectPoolsConfig `json:"redirect_pools_config,omitempty"`

	RedirectPoolsStickySessionConfig *CreateRedirectPoolsStickySessionConfig `json:"redirect_pools_sticky_session_config,omitempty"`

	FixedResponseConfig *CreateFixtedResponseConfig `json:"fixed_response_config,omitempty"`

	RedirectPoolsExtendConfig *CreateRedirectPoolsExtendConfig `json:"redirect_pools_extend_config,omitempty"`

	// 参数解释：转发策略关联的转发规则对象。  约束限制： - rules列表中最多含有10个rule规则 （若rule中包含conditions字段，一条condition算一个规则）， 且列表中type为HOST_NAME，PATH，METHOD，SOURCE_IP的rule不能重复，至多指定一条。 - 仅支持全量替换。 - 如果l7policy 是重定向到listener的话，不允许创建l7rule。
	Rules *[]CreateL7PolicyRuleOption `json:"rules,omitempty"`
}

func (o CreateL7PolicyOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateL7PolicyOption struct{}"
	}

	return strings.Join([]string{"CreateL7PolicyOption", string(data)}, " ")
}
