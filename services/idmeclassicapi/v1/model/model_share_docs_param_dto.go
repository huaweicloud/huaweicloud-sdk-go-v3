package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShareDocsParamDto struct {

	// **参数解释：**  结构化文档ID，系统生成的文档主键唯一标识，用于指定待分享的目标文档。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	StructuredDocId *string `json:"structured_doc_id,omitempty"`

	// **参数解释：**  被分享用户ID，用于指定接收文档分享的目标用户。  **约束限制：**  不涉及。  **取值范围：**  - all：表示分享给所有用户。 - 其他值：指定具体用户的ID。  **默认取值：**  不涉及。
	SharedUserId *string `json:"shared_user_id,omitempty"`

	// **参数解释：**  被分享用户名，用于指定接收文档分享的目标用户名称。  **约束限制：**  不涉及。  **取值范围：**  - all：表示分享给所有用户。 - 其他值：指定具体用户的名称。  **默认取值：**  不涉及。
	SharedUserName *string `json:"shared_user_name,omitempty"`

	// **参数解释：**  分享用户ID，用于标识执行本次分享操作的用户。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	ShareUserId *string `json:"share_user_id,omitempty"`

	// **参数解释：**  分享用户名，用于标识执行本次分享操作的用户名称。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	ShareUserName *string `json:"share_user_name,omitempty"`

	// **参数解释：**  认证类型，用于指定被分享用户对文档的访问权限级别。  **约束限制：**  不涉及。  **取值范围：**  - read：只读权限，被分享用户仅可查看文档内容，适用于文档发布、信息同步等场景。 - write：读写权限，被分享用户可查看和编辑文档内容，适用于协同设计、工艺评审等场景。  **默认取值：**  不涉及。
	AuthType *string `json:"auth_type,omitempty"`

	// **参数解释：**  更新者账号，用于记录执行本次分享操作的用户信息。 若不指定，默认使用当前调用者账号。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  当前调用者账号。
	Modifier *string `json:"modifier,omitempty"`
}

func (o ShareDocsParamDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShareDocsParamDto struct{}"
	}

	return strings.Join([]string{"ShareDocsParamDto", string(data)}, " ")
}
