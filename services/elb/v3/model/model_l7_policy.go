package model

import (
	"encoding/json"

	"strings"
)

// policy对象。
type L7Policy struct {
	// 转发策略的转发动作；取值：REDIRECT_TO_POOL：转发到后端云服务器组；REDIRECT_TO_LISTENER：重定向到监听器

	Action string `json:"action"`
	// 转发策略的管理状态；该字段为预留字段，暂未启用。默认为true。

	AdminStateUp bool `json:"admin_state_up"`
	// 转发策略描述信息

	Description string `json:"description"`
	// 转发策略ID

	Id string `json:"id"`
	// 转发策略对应的监听器ID。当action为REDIRECT_TO_POOL时，只支持创建在PROTOCOL为HTTP或TERMINATED_HTTPS的listener上。 当action为REDIRECT_TO_LISTENER时，只支持创建在PROTOCOL为HTTP的listener上。

	ListenerId string `json:"listener_id"`
	// 转发策略名称

	Name string `json:"name"`
	// 转发策略的优先级，从1递增，最高100。该字段为预留字段，暂未启用。

	Position int32 `json:"position"`
	// 转发策略所在的项目ID。

	ProjectId string `json:"project_id"`
	// provisioning状态，可以为ACTIVE、PENDING_CREATE 或者ERROR。 默认为ACTIVE。

	ProvisioningStatus string `json:"provisioning_status"`
	// 转发到的listener的ID，当action为REDIRECT_TO_LISTENER时生效。使用说明：只支持protocol为TERMINATED_HTTPS的listener。不能指定为其他loadbalancer下的listener。当action为REDIRECT_TO_POOL时，不可指定。

	RedirectListenerId string `json:"redirect_listener_id"`
	// 转发到pool的ID。当action为REDIRECT_TO_POOL时生效。使用说明：指定的pool不能是listener的default_pool。不能是其他listener的l7policy使用的pool。当action为REDIRECT_TO_LISTENER时，不可指定。

	RedirectPoolId string `json:"redirect_pool_id"`
	// 转发到的url。该字段未启用。

	RedirectUrl string `json:"redirect_url"`
	// 转发策略关联的转发规则列表

	Rules []RuleRef `json:"rules"`
}

func (o L7Policy) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "L7Policy struct{}"
	}

	return strings.Join([]string{"L7Policy", string(data)}, " ")
}
