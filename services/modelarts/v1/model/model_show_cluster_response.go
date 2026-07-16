package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterResponse Response Object
type ShowClusterResponse struct {

	// **参数解释**： 逻辑资源池ID。 **取值范围**： 不涉及。
	LogicClusterId *string `json:"logic_cluster_id,omitempty"`

	// **参数解释**：资源池状态。 **取值范围**：枚举类型，取值如下： - PENDING：等待中。 - INITIALIZING：初始化中。 - INITIALIZE_FAILED：初始化失败。 - ACTIVE：可用。 - DELETING：删除中。 - DELETED：已删除。 - DELETE_FAILED：删除失败。 - MIGRATING：迁移中。
	Status *string `json:"status,omitempty"`

	// **参数解释**：资源池ID。 **取值范围**：不涉及。
	PoolId *string `json:"pool_id,omitempty"`

	// **参数解释**：资源池类型。  **取值范围**：枚举类型，取值如下： - MANAGED：公共池。 - DEDICATED：专属池。
	Type *string `json:"type,omitempty"`

	// **参数解释**：资源类别。 **取值范围**：枚举类型，取值如下： - GPU - CPU - ASCEND
	ResourceCategories *string `json:"resource_categories,omitempty"`

	// **参数解释**：工作空间ID。获取方法请参见[[查询工作空间列表](ListWorkspace.xml)](tag:hc,hk)。未创建工作空间时默认值为“0”，存在创建并使用的工作空间，以实际取值为准。 **取值范围**：不涉及。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**：实例创建的时间，UTC毫秒。 **取值范围**：不涉及。
	CreateAt *string `json:"create_at,omitempty"`

	// **参数解释**：实例最后更新的时间，UTC毫秒。 **取值范围**：不涉及。
	UpdateAt *string `json:"update_at,omitempty"`

	// **参数解释**：用户项目ID，获取方法请参见[获取项目ID和名称](modelarts_03_0147.xml)。 **取值范围**：不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**：账号ID。 **取值范围**：不涉及。
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释**：资源池规格。
	Flavors *[]Flavor `json:"flavors,omitempty"`

	// **参数解释**：资源池是否允许实例以root启动。
	IsAllowRoot    *bool `json:"is_allow_root,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterResponse struct{}"
	}

	return strings.Join([]string{"ShowClusterResponse", string(data)}, " ")
}
