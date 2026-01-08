package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityPolicy **参数解释**：自定义安全策略信息。
type SecurityPolicy struct {

	// **参数解释**：自定义安全策略的ID。  **取值范围**：不涉及
	Id string `json:"id"`

	// **参数解释**：自定义安全策略的项目ID。  **取值范围**：不涉及
	ProjectId string `json:"project_id"`

	// **参数解释**：自定义安全策略的名称。  **取值范围**：不涉及
	Name string `json:"name"`

	// **参数解释**：自定义安全策略的描述。  **取值范围**：不涉及
	Description string `json:"description"`

	// **参数解释**：关联的监听器。
	Listeners []ListenerRef `json:"listeners"`

	// **参数解释**：自定义安全策略的TLS协议列表。  **取值范围**：不涉及
	Protocols []string `json:"protocols"`

	// **参数解释**：自定义安全策略的加密套件列表。  **取值范围**：不涉及
	Ciphers []string `json:"ciphers"`

	// **参数解释**：资源所属的企业项目ID。  **取值范围**： - \"0\"：表示资源属于default企业项目。 - UUID格式的字符串，表示非默认企业项目。  [不支持该字段，请勿使用。](tag:dt,hcso_dt)
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**：创建时间。  **取值范围**：格式：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。
	CreatedAt string `json:"created_at"`

	// **参数解释**：更新时间。  **取值范围**：格式：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。
	UpdatedAt string `json:"updated_at"`
}

func (o SecurityPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityPolicy struct{}"
	}

	return strings.Join([]string{"SecurityPolicy", string(data)}, " ")
}
