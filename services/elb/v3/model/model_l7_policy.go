package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type L7Policy struct {

	// **参数解释**：转发策略的转发动作。  **取值范围**： - REDIRECT_TO_POOL：转发到后端服务器组； - REDIRECT_TO_LISTENER：重定向到监听器； - REDIRECT_TO_URL：重定向到URL； - FIXED_RESPONSE：返回固定响应体。  [不支持REDIRECT_TO_URL和FIXED_RESPONSE](tag:hcso_dt)
	Action string `json:"action"`

	// **参数解释**：转发策略的管理状态。  **取值范围**：只支持设置为true。
	AdminStateUp bool `json:"admin_state_up"`

	// **参数解释**：转发策略描述信息。  **取值范围**：不涉及
	Description string `json:"description"`

	// **参数解释**：转发策略ID。  **取值范围**：不涉及
	Id string `json:"id"`

	// **参数解释**：转发策略所属的监听器ID。  **取值范围**：不涉及
	ListenerId string `json:"listener_id"`

	// **参数解释**：转发策略名称  **取值范围**：不涉及
	Name string `json:"name"`

	// **参数解释**：转发策略的优先级，不支持更新。  不支持该字段，请勿使用。
	Position int32 `json:"position"`

	// **参数解释**：转发策略的优先级。数字越小表示优先级越高。  **取值范围**： - 当action为REDIRECT_TO_LISTENER时，支持指定为0-10000。 - 其它action取值，支持指定为1-10000。  [不支持该字段，请勿使用。](tag:hcso_dt)  [荷兰region不支持该字段，请勿使用。](tag:dt)
	Priority *int32 `json:"priority,omitempty"`

	// **参数解释**：转发策略所在的项目ID。  **取值范围**：不涉及
	ProjectId string `json:"project_id"`

	// **参数解释**：转发策略的配置状态。  **取值范围**： - ACTIVE: 默认值，表示正常。 [- ERROR: 表示当前策略与同一监听器下的其他策略存在相同的规则配置。](tag:hws,hws_hk,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,srg,fcs)
	ProvisioningStatus string `json:"provisioning_status"`

	// **参数解释**：转发到pool的ID。  **取值范围**：不涉及
	RedirectPoolId string `json:"redirect_pool_id"`

	// **参数解释**：转发到的listener的ID。  **取值范围**：不涉及
	RedirectListenerId string `json:"redirect_listener_id"`

	// **参数解释**：转发到的url。  **取值范围**：必须满足格式: protocol://host:port/path?query。  不支持该字段，请勿使用。
	RedirectUrl string `json:"redirect_url"`

	// **参数解释**：转发策略关联的转发规则列表。
	Rules []RuleRef `json:"rules"`

	RedirectUrlConfig *RedirectUrlConfig `json:"redirect_url_config"`

	// **参数解释**：转发策略服务器组的权重配置。当监听器的高级转发策略功能（enhance_l7policy_enable）开启后才会生效，未开启传入该字段会报错。
	RedirectPoolsConfig *[]RedirectPoolsConfig `json:"redirect_pools_config,omitempty"`

	RedirectPoolsStickySessionConfig *RedirectPoolsStickySessionConfig `json:"redirect_pools_sticky_session_config,omitempty"`

	RedirectPoolsExtendConfig *RedirectPoolsExtendConfig `json:"redirect_pools_extend_config,omitempty"`

	FixedResponseConfig *FixtedResponseConfig `json:"fixed_response_config"`

	// **参数解释**：创建时间。  **取值范围**：格式：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。  [注意：独享型实例的历史数据以及共享型实例下的资源，不返回该字段。 ](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,srg,fcs,dt,hk_tm)
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释**：更新时间。  **取值范围**：格式：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。  [注意：独享型实例的历史数据以及共享型实例下的资源，不返回该字段。 ](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,srg,fcs,dt,hk_tm)
	UpdatedAt *string `json:"updated_at,omitempty"`

	// **参数解释**：企业项目ID。  **取值范围**：不涉及
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o L7Policy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "L7Policy struct{}"
	}

	return strings.Join([]string{"L7Policy", string(data)}, " ")
}
