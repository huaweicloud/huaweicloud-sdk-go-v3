package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAssetOperationRecordsRequest Request Object
type ListAssetOperationRecordsRequest struct {

	// 工作空间ID。[获取方法请参见[查询工作空间列表](ListWorkspace.xml)。](tag:hc)未创建工作空间时默认值为“0”，存在创建并使用的工作空间，以实际取值为准。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**：限制量，单次查询总量，必须由数字组成。 **约束限制**：不涉及。 **取值范围**：1-1000。 **默认取值**：100。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：偏移量，查询起始偏移，必须由数字组成。 **约束限制**：不涉及。 **取值范围**0-100000000。 **默认取值**：0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：排序方向。 **约束限制**：不涉及。 **取值范围**： 枚举值： - ASC: 升序 - DESC: 降序。 **默认取值**：DESC。
	SortDir *string `json:"sort_dir,omitempty"`

	// **参数解释**：排序字段。 **约束限制**：不涉及。 **取值范围**：枚举值： - created_at: 创建时间 - updated_at: 更新时间。 **默认取值**：created_at。
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释**：操作记录ID。 **约束限制**：1-64位字符串。 **取值范围**：不涉及。 **默认取值**：不涉及。
	RecordId *string `json:"record_id,omitempty"`

	// **参数解释**：作业类型。 **约束限制**：1-40位字符串。 **取值范围**：不涉及。 **默认取值**：不涉及。
	JobType *string `json:"job_type,omitempty"`

	// **参数解释**：关联任务ID。 **约束限制**：1-64位字符串。 **取值范围**：不涉及。 **默认取值**：不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**：关联任务名称。 **约束限制**：1-128位字符串。 **取值范围**：不涉及。 **默认取值**：不涉及。
	JobName *string `json:"job_name,omitempty"`

	// **参数解释**：关联任务状态。 **约束限制**：1-40位字符串。 **取值范围**：不涉及。 **默认取值**：不涉及。
	JobStatus *string `json:"job_status,omitempty"`

	// **参数解释**：操作用户ID。 **约束限制**：1-36位字符串。 **取值范围**：不涉及。 **默认取值**：不涉及。
	UserId *string `json:"user_id,omitempty"`

	// **参数解释**：操作用户的用户名。 **约束限制**：1-128位字符串。 **取值范围**：不涉及。 **默认取值**：不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**：资产ID。 **约束限制**：1-64位字符串。 **取值范围**：不涉及。 **默认取值**：不涉及。
	AssetId *string `json:"asset_id,omitempty"`

	// **参数解释**：资产code。 **约束限制**：1-64位字符串。 **取值范围**：不涉及。 **默认取值**：不涉及。
	AssetCode *string `json:"asset_code,omitempty"`

	// **参数解释**：任务是否删除。 **约束限制**：不涉及。 **取值范围**：true,false。 **默认取值**：不涉及。
	IsDeleted *bool `json:"is_deleted,omitempty"`
}

func (o ListAssetOperationRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssetOperationRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListAssetOperationRecordsRequest", string(data)}, " ")
}
