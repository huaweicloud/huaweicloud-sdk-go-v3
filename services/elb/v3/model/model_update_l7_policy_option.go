package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新七层转发策略的请求参数。
type UpdateL7PolicyOption struct {

	// 转发策略的管理状态，默认为true。  不支持该字段，请勿使用。
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 转发策略描述信息。
	Description *string `json:"description,omitempty"`

	// 转发策略名称。
	Name *string `json:"name,omitempty"`

	// 转发到的listener的ID。  使用说明： - 当action为REDIRECT_TO_LISTENER时不能更新为空或null。 - 只支持protocol为HTTPS/TERMINATED_HTTPS的listener。 - 不能指定为其他loadbalancer下的listener。 - 当action为REDIRECT_TO_POOL时，创建或更新时不能传入该参数。
	RedirectListenerId *string `json:"redirect_listener_id,omitempty"`

	// 转发到pool的ID。  使用说明： - 指定的pool不能是listener的default_pool。不能是其他listener的l7policy使用的pool。 - 当action为REDIRECT_TO_POOL时为必选字段，不能更新为空或null。 当action为REDIRECT_TO_LISTENER时，不可指定。
	RedirectPoolId *string `json:"redirect_pool_id,omitempty"`

	RedirectUrlConfig *UpdateRedirectUrlConfig `json:"redirect_url_config,omitempty"`

	FixedResponseConfig *UpdateFixtedResponseConfig `json:"fixed_response_config,omitempty"`

	// 转发策略关联的转发规则对象。 详细参考表l7rule字段说明。rules列表中最多含有10个rule规则 （若rule中包含conditions字段，一条condition算一个规则）， 且列表中type为HOST_NAME，PATH，METHOD，SOURCE_IP的rule不能重复，至多指定一条。
	Rules *[]CreateRuleOption `json:"rules,omitempty"`

	// 转发策略的优先级。数字越小表示优先级越高，同一监听器下不允许重复。  当监听器的高级转发策略功能（enhance_l7policy_enable）开启后才会生效，未开启传入该字段会报错。  当action为REDIRECT_TO_LISTENER时，仅支持指定为0，优先级最高。  当关联的listener没有开启enhance_l7policy_enable，按原有policy的排序逻辑，自动排序。 各域名之间优先级独立，相同域名下，按path的compare_type排序， 精确>前缀>正则，匹配类型相同时，path的长度越长优先级越高。 若policy下只有域名rule，没有路径rule，默认path为前缀匹配/。  当关联的listener开启了enhance_l7policy_enable，且不传该字段， 则新创建的转发策略的优先级的值为：同一监听器下已有转发策略的优先级的最大值+1。 因此，若当前已有转发策略的优先级的最大值是10000，新创建会因超出取值范围10000而失败。 此时可通过传入指定priority，或调整原有policy的优先级来避免错误。 若监听器下没有转发策略，则新建的转发策略的优先级为1。  [共享型负载均衡器下的转发策略不支持该字段。 ](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,fcs,dt,hk_tm)  [不支持该字段，请勿使用。](tag:hcso_dt)
	Priority *int32 `json:"priority,omitempty"`

	// 转发到的后端主机组的配置。当action为REDIRECT_TO_POOL时生效。  使用说明： - 当action为REDIRECT_TO_POOL时redirect_pool_id和redirect_pools_config 必须指定一个，两个都指定时按redirect_pools_config生效。 - 当action为REDIRECT_TO_LISTENER时，不可指定。  只支持全量覆盖。
	RedirectPoolsConfig *[]CreateRedirectPoolsConfig `json:"redirect_pools_config,omitempty"`
}

func (o UpdateL7PolicyOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateL7PolicyOption struct{}"
	}

	return strings.Join([]string{"UpdateL7PolicyOption", string(data)}, " ")
}
