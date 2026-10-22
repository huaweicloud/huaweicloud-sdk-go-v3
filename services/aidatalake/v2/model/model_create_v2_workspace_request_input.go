package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateV2WorkspaceRequestInput **参数解释**：创建工作空间的请求体。
type CreateV2WorkspaceRequestInput struct {

	// **参数解释**：工作空间名称。 **约束限制**：不涉及。 **取值范围**：长度为4~32的英文字母、数字、下划线和中划线的组合。 **默认取值**：不涉及。
	Name string `json:"name"`

	// **参数解释**：用户输入的描述。 **约束限制**：不涉及。 **取值范围**：不允许输入感叹号、小于号、大于号、等号、与号、双引号、单引号、艾特符、井号、百分号、分号、左括号、右括号、竖线、星号、反斜杠。长度为0~255个字符。 **默认取值**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：LakeFormation服务的实例ID，可在LakeFormation控制台实例列表中查看，或通过调用LakeFormation服务的查询实例列表接口获取。 **约束限制**：不涉及。 **取值范围**：长度为32~36的英文字母、数字和中划线的组合。 **默认取值**：不涉及。
	MetastoreId string `json:"metastore_id"`

	// **参数解释**：企业项目ID，获取方法，请参见[获取企业项目ID](aidatalake_02_0042.xml)。 **约束限制**：不涉及。 **取值范围**：长度为0~128个英文字母、数字、下划线和中划线的组合。 **默认取值**：0。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**：标签列表，每个元素为一个V2ResourceTag对象，包含key（标签键）和value（标签值）字段。 **约束限制**：最多20个标签。
	Tags *[]V2ResourceTag `json:"tags,omitempty"`

	// **参数解释**：是否跨租户访问。 **取值范围**：   - true：允许跨租户访问。   - false：不允许跨租户访问。
	CrossTenantAccess *bool `json:"cross_tenant_access,omitempty"`

	// **参数解释**：目标租户Domain ID。 **取值范围**：长度为1~64的英文字母、数字和中划线的组合。
	TargetDomainId *string `json:"target_domain_id,omitempty"`

	// **参数解释**：目标租户Project ID。 **取值范围**：长度为1~64的英文字母、数字和中划线的组合。
	TargetProjectId *string `json:"target_project_id,omitempty"`
}

func (o CreateV2WorkspaceRequestInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateV2WorkspaceRequestInput struct{}"
	}

	return strings.Join([]string{"CreateV2WorkspaceRequestInput", string(data)}, " ")
}
