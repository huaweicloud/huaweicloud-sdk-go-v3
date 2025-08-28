package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIpGroupOption **参数解释**：创建IP地址组的请求参数。  **约束限制**：不涉及
type CreateIpGroupOption struct {

	// **参数解释**：项目ID。获取方式请参见[获取项目ID](elb_fl_0008.xml)。  **约束限制**：不涉及  **取值范围**：长度为32个字符，由小写字母和数字组成。  **默认取值**：不涉及  > 该字段实际无效，最终使用url中的project_id。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**：IP地址组的描述。  **约束限制**：不涉及  **取值范围**：长度为0-255个字符。  **默认取值**：不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**：IP地址组的名称。  **约束限制**：不涉及  **取值范围**：长度为0-255个字符。  **默认取值**：不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**：IP地址组中的IP列表，支持IPv4和IPv6类型地址。  **约束限制**：不涉及
	IpList []CreateIpGroupIpOption `json:"ip_list"`

	// **参数解释**：资源所属的企业项目ID。创建时不传则资源属于default企业项目，返回enterprise_project_id=\"0\"。  **约束限制**：不能传入空字符串\"\"、\"0\"或不存在的企业项目ID。  **取值范围**：不涉及  **默认取值**：\"0\"  [不支持该字段，请勿使用。](tag:dt,hcso_dt)
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreateIpGroupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpGroupOption struct{}"
	}

	return strings.Join([]string{"CreateIpGroupOption", string(data)}, " ")
}
