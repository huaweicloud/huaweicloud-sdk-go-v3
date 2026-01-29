package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DomainVo 项目空间对象
type DomainVo struct {

	// **参数解释：**  项目空间的唯一Id。 **取值范围：**  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：**  项目空间名称。 **取值范围：**  不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：**  项目空间名称。 **取值范围：**  不涉及。
	Title *string `json:"title,omitempty"`

	// **参数解释：**  项目空间对应的32位UUId。 **取值范围：**  不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释：**  项目空间的父项目空间Id，仅在项目群中使用。 **取值范围：**  不涉及。
	ParentId *string `json:"parent_id,omitempty"`

	// **参数解释：**  项目空间类型。 **取值范围：**  -  Project：项目 - Group：项目群
	Category *string `json:"category,omitempty"`

	// **参数解释：**  项目空间创建人Id。 **取值范围：**  不涉及。
	CreatedBy *string `json:"created_by,omitempty"`

	// **参数解释：**  项目空间对应的模型Id。 **取值范围：**  不涉及。 - 10001：系统设备类 - 10002：独立软件类 - 10003：自营软件/云服务类
	ModelId *string `json:"model_id,omitempty"`
}

func (o DomainVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainVo struct{}"
	}

	return strings.Join([]string{"DomainVo", string(data)}, " ")
}
