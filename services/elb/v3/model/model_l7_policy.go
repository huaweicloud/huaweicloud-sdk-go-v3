package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// L7Policy policy对象。
type L7Policy struct {

	// 参数解释：转发策略的转发动作。  约束限制： - REDIRECT_TO_LISTENER的优先级最高，配置了以后，该监听器下的其他policy会失效。 - 当action为REDIRECT_TO_POOL时， 只支持创建在PROTOCOL为HTTP、HTTPS、TERMINATED_HTTPS的listener上。 - 当action为REDIRECT_TO_LISTENER时，只支持创建在PROTOCOL为HTTP的listener上。  取值范围： - REDIRECT_TO_POOL：转发到后端服务器组； - REDIRECT_TO_LISTENER：重定向到监听器； - REDIRECT_TO_URL：重定向到URL； - FIXED_RESPONSE：返回固定响应体。  [不支持REDIRECT_TO_URL和FIXED_RESPONSE](tag:hcso_dt)
	Action string `json:"action"`

	// 参数解释：转发策略的管理状态。  约束限制：只支持设置为true。
	AdminStateUp bool `json:"admin_state_up"`

	// 参数解释：转发策略描述信息。
	Description string `json:"description"`

	// 参数解释：转发策略ID。
	Id string `json:"id"`

	// 参数解释：转发策略所属的监听器ID。
	ListenerId string `json:"listener_id"`

	// 参数解释：转发策略名称
	Name string `json:"name"`

	// 参数解释：转发策略的优先级，不支持更新。  不支持该字段，请勿使用。
	Position int32 `json:"position"`

	// 参数解释：转发策略的优先级。数字越小表示优先级越高。  约束限制： - 同一个监听器下不同转发策略之间不允许重复的优先级数值。 - 当关联的监听器的高级转发策略功能（enhance_l7policy_enable）开启后才会生效，未开启传入该字段会报错。 - 当关联的监听器的高级转发策略功能（enhance_l7policy_enable）未开启，按原有policy的排序逻辑，自动排序。 不同域名优先级独立。相同域名下，按path的compare_type排序， 精确>前缀>正则，匹配类型相同时，path的长度越长优先级越高。 若policy下只有域名rule，没有路径rule，默认path为前缀匹配/。 [- 共享型负载均衡器下的转发策略不支持该字段。](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,fcs,dt,hk_tm)  取值范围： - 当action为REDIRECT_TO_LISTENER时，支持指定为0-10000。 - 其它action取值，支持指定为1-10000。  默认取值： - 若关联的监听器的高级转发策略功能（enhance_l7policy_enable）未开启，且不传入该字段，则新创建的转发策略的优先级的值为1。 - 当action为REDIRECT_TO_LISTENER时，则新创建的转发策略的优先级的值为0。 - 其它action取值，新创建的转发策略的优先级的值为同一监听器下已有转发策略的优先级的最大值+1。   + 若监听器下没有转发策略，则新建的转发策略的优先级为1。   + 若当前已有转发策略的优先级的最大值是10000，则新创建的转发策略会因超出取值范围10000而失败。此时可通过传入指定priority，或调整原有policy的优先级来避免错误。  [不支持该字段，请勿使用。](tag:hcso_dt)  [荷兰region不支持该字段，请勿使用。](tag:dt,dt_test)
	Priority *int32 `json:"priority,omitempty"`

	// 参数解释：转发策略所在的项目ID。
	ProjectId string `json:"project_id"`

	// 参数解释：转发策略的配置状态。  取值范围： - ACTIVE: 默认值，表示正常。 [- ERROR: 表示当前策略与同一监听器下的其他策略存在相同的规则配置。 ](tag:hws,hws_hk,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,fcs)
	ProvisioningStatus string `json:"provisioning_status"`

	// 参数解释：转发到pool的ID。  约束限制：当action为REDIRECT_TO_POOL时生效。
	RedirectPoolId string `json:"redirect_pool_id"`

	// 参数解释：转发到的listener的ID。  约束限制： - 当action为REDIRECT_TO_LISTENER时必选。 - 只支持protocol为HTTPS/TERMINATED_HTTPS的listener。 - 不能指定为其他loadbalancer下的listener。 - 当action为REDIRECT_TO_POOL时，创建或更新时不能传入该参数。
	RedirectListenerId string `json:"redirect_listener_id"`

	// 参数解释：转发到的url。  取值范围：必须满足格式: protocol://host:port/path?query。  不支持该字段，请勿使用。
	RedirectUrl string `json:"redirect_url"`

	// 参数解释：转发策略关联的转发规则列表。
	Rules []RuleRef `json:"rules"`

	RedirectUrlConfig *RedirectUrlConfig `json:"redirect_url_config"`

	// 参数解释：转发到多个主机组列表。  约束限制：一个policy最多配置5个pool。
	RedirectPoolsConfig *[]RedirectPoolsConfig `json:"redirect_pools_config,omitempty"`

	RedirectPoolsStickySessionConfig *RedirectPoolsStickySessionConfig `json:"redirect_pools_sticky_session_config,omitempty"`

	RedirectPoolsExtendConfig *RedirectPoolsExtendConfig `json:"redirect_pools_extend_config,omitempty"`

	FixedResponseConfig *FixtedResponseConfig `json:"fixed_response_config"`

	// 参数解释：创建时间。  取值范围：格式：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。  [注意：独享型实例的历史数据以及共享型实例下的资源，不返回该字段。 ](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,fcs,dt,hk_tm)
	CreatedAt *string `json:"created_at,omitempty"`

	// 参数解释：更新时间。  取值范围：格式：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。  [注意：独享型实例的历史数据以及共享型实例下的资源，不返回该字段。 ](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,fcs,dt,hk_tm)
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o L7Policy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "L7Policy struct{}"
	}

	return strings.Join([]string{"L7Policy", string(data)}, " ")
}
