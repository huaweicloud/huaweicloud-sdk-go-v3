package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PrivilegesResponseV5Result **参数解释**：  请求返回的结果，角色及权限信息。 **取值范围**： 不涉及。
type PrivilegesResponseV5Result struct {

	// **参数解释**：  操作权限的列表。 **取值范围**： 英文字符串，使用英文逗号分隔。
	Operations *string `json:"operations,omitempty"`

	// **参数解释**：  操作权限的序列号。 **取值范围**： 数字。
	OperationsIndex *[]int32 `json:"operations_index,omitempty"`

	// **参数解释**：  角色ID。 **取值范围**： 32位英文、数字随机字符串。
	RoleId *string `json:"role_id,omitempty"`

	// **参数解释**：  角色的英文名称。 **取值范围**：   Project manager，Product manager，Test manager，Operation manager，System engineer，Committer，Developer，Tester，Participant，Viewer及自定义角色的英文名称。
	RoleName *string `json:"role_name,omitempty"`

	// **参数解释**：  角色的中文名称。 **取值范围**： 项目经理，产品经理，测试经理，运维经理，系统工程师，Committer，开发人员，测试人员，参与者，浏览者及自定义角色的中文名称。
	RoleChineseName *string `json:"role_chinese_name,omitempty"`

	// **参数解释**：  项目ID。 **取值范围**： 32位英文、数字随机字符串。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**：  地域服务ID。 **取值范围**： 32位英文、数字随机字符串。
	AreaServiceId *string `json:"area_service_id,omitempty"`

	// **参数解释**： 授权对象路径。 **取值范围**： 英文、数字、斜线（/）、星号（*）字符串
	GrantedObjectPath *string `json:"granted_object_path,omitempty"`

	// **参数解释**： 授权对象类型ID。 **取值范围**： 32位英文、数字随机字符串。
	GrantedObjectTypeId *string `json:"granted_object_type_id,omitempty"`
}

func (o PrivilegesResponseV5Result) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivilegesResponseV5Result struct{}"
	}

	return strings.Join([]string{"PrivilegesResponseV5Result", string(data)}, " ")
}
