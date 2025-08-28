package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IpGroup **参数解释**：IP地址组信息。
type IpGroup struct {

	// **参数解释**：IP地址组的ID。  **取值范围**：不涉及
	Id string `json:"id"`

	// **参数解释**：IP地址组的名称。  **取值范围**：不涉及
	Name string `json:"name"`

	// **参数解释**：IP地址组的描述信息。  **取值范围**：不涉及
	Description string `json:"description"`

	// **参数解释**：IP地址组中包含的IP地址列表。[]表示任意IP。
	IpList []IpInfo `json:"ip_list"`

	// **参数解释**：与IP地址组关联的监听器的ID列表。
	Listeners []ListenerRef `json:"listeners"`

	// **参数解释**：IP地址组的项目ID。  **取值范围**：不涉及
	ProjectId string `json:"project_id"`

	// **参数解释**：资源所属的企业项目ID。  **取值范围**： - \"0\"：表示资源属于default企业项目。 - UUID格式的字符串，表示非默认企业项目。  [不支持该字段，请勿使用。](tag:dt,hcso_dt)
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**：IP地址组的创建时间。  **取值范围**：格式：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。
	CreatedAt string `json:"created_at"`

	// **参数解释**：IP地址组的更新时间。  **取值范围**：格式：yyyy-MM-dd'T'HH:mm:ss'Z'，UTC时区。
	UpdatedAt string `json:"updated_at"`
}

func (o IpGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpGroup struct{}"
	}

	return strings.Join([]string{"IpGroup", string(data)}, " ")
}
