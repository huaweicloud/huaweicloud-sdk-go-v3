package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StructuredDocShareViewDto struct {

	// **参数解释：**  分享权限记录的唯一标识，可用于批量删除分享权限接口定位目标记录。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：**  创建者账号，标识创建该分享记录的用户。  **取值范围：**  不涉及。
	Creator *string `json:"creator,omitempty"`

	// **参数解释：**  更新者账号，标识最后更新该分享记录的用户。  **取值范围：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`

	// **参数解释：**  创建时间。使用UTC+0时间格式，格式为yyyy-MM-ddTHH:mm:ss.SSSZ。  **取值范围：**  不涉及。
	CreateTime *string `json:"createTime,omitempty"`

	// **参数解释：**  最后更新时间。使用UTC+0时间格式，格式为yyyy-MM-ddTHH:mm:ss.SSSZ。  **取值范围：**  不涉及。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// **参数解释：**  系统版本号，用于数据版本控制。  **取值范围：**  不涉及。
	RdmVersion *int32 `json:"rdmVersion,omitempty"`

	// **参数解释：**  软删除标识。  **取值范围：**  - 0：表示未删除。 - 1：表示已删除。
	RdmDeleteFlag *int32 `json:"rdmDeleteFlag,omitempty"`

	// **参数解释：**  扩展类型，标识对象的扩展类别。  **取值范围：**  不涉及。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	Tenant *TenantViewDto `json:"tenant,omitempty"`

	// **参数解释：**  类名，标识对象的Java类名称。  **取值范围：**  不涉及。
	ClassName *string `json:"className,omitempty"`

	StructuredDoc *StructuredDocViewDto `json:"structuredDoc,omitempty"`

	// **参数解释：**  分享用户名，标识执行分享操作的用户名称。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	ShareUserName *string `json:"shareUserName,omitempty"`

	// **参数解释：**  被分享用户名，标识接收文档分享的用户名称。  **约束限制：**  不涉及。  **取值范围：**  - all：表示分享给所有人。 - 其他值：指定具体用户的名称。  **默认取值：**  不涉及。
	SharedUserName *string `json:"sharedUserName,omitempty"`

	// **参数解释：**  被分享用户ID，标识接收文档分享的用户ID。  **约束限制：**  不涉及。  **取值范围：**  - all：表示分享给所有人。 - 其他值：指定具体用户的ID。  **默认取值：**  不涉及。
	SharedUserId *string `json:"sharedUserId,omitempty"`

	// **参数解释：**  认证类型，标识被分享用户对文档的访问权限级别。  **约束限制：**  不涉及。  **取值范围：**  - read：只读权限，被分享用户仅可查看文档内容。 - write：读写权限，被分享用户可查看和编辑文档内容。  **默认取值：**  不涉及。
	AuthType *string `json:"authType,omitempty"`

	// **参数解释：**  分享用户ID，标识执行分享操作的用户ID。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	ShareUserId *string `json:"shareUserId,omitempty"`
}

func (o StructuredDocShareViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StructuredDocShareViewDto struct{}"
	}

	return strings.Join([]string{"StructuredDocShareViewDto", string(data)}, " ")
}
