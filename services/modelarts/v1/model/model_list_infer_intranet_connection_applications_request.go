package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInferIntranetConnectionApplicationsRequest Request Object
type ListInferIntranetConnectionApplicationsRequest struct {

	// **参数解释：** 内网访问场景。 **约束限制：** 不涉及。 **取值范围：** - POOL：用户资源池接入场景 - VPC：用户VPC网络接入场景 **默认取值：** 不涉及。
	Scene *string `json:"scene,omitempty"`

	// **参数解释：** 内网接入状态，支持列表查询。 **约束限制：** 不涉及。 **取值范围：** - APPROVING：审批中 - REJECTED：拒绝 - CONNECTING：接入中 - CONNECTED：已接入 - CANCELED：已取消 - FAILED：失败 - DELETING：删除中 **默认取值：** 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释：** 内网接入ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 服务ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ServiceId *string `json:"service_id,omitempty"`

	// **参数解释：** 服务名称。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ServiceName *string `json:"service_name,omitempty"`

	// **参数解释：** VPC网络的id。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数解释：** VPC网络的名称。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	VpcName *string `json:"vpc_name,omitempty"`

	// **参数解释：** 资源池的id。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	PoolId *string `json:"pool_id,omitempty"`

	// **参数解释：** 分页列表查询的偏移量。 **约束限制：** offset必须是limit的整数倍。 **取值范围：** 不涉及。 **默认取值：** 0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 指定返回的最大条目数。 **约束限制：** 不涉及。 **取值范围：** [1,500] **默认取值：** 10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 工作空间ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释：** 排序方式 **约束限制：** 不涉及。 **取值范围：** - ASC: 递增排序。 - DESC: 递减排序。 **默认取值：** DESC。
	SortDir *string `json:"sort_dir,omitempty"`

	// **参数解释：** 排序字段，多个字段以\",\"分隔，支持create_at, update_at，默认值update_at。 **约束限制：** 不涉及。 **取值范围：** - create_at：按创建时间排序。 - update_at：按更新时间排序。 **默认取值：** update_at。
	SortKey *string `json:"sort_key,omitempty"`
}

func (o ListInferIntranetConnectionApplicationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInferIntranetConnectionApplicationsRequest struct{}"
	}

	return strings.Join([]string{"ListInferIntranetConnectionApplicationsRequest", string(data)}, " ")
}
