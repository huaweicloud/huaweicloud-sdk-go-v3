package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPublicipPoolTypesRequest Request Object
type ShowPublicipPoolTypesRequest struct {

	// - 功能说明：分页查询起始的资源ID，为空时为查询第一页
	Marker *string `json:"marker,omitempty"`

	// - 功能说明：每页返回的资源个数 - 取值范围：1~2000
	Limit *int32 `json:"limit,omitempty"`

	// - 功能说明：查询字段，形式为“fields=id&fields=name&...” - 支持字段：id/name/size/used/project_id/status/billing_info/created_at/updated_at/type/shared/is_common/description/tags/enterprise_project_id/allow_share_bandwidth_types/public_border_group
	Fields *string `json:"fields,omitempty"`

	// - 功能说明：排序字段，形式为“sort_key=id&sort_dir=asc” - 支持字段：id/name/created_at/updated_at/public_border_group
	SortKey *string `json:"sort_key,omitempty"`

	// - 功能说明：排序方向 - 取值范围：   - asc 顺序   - desc 倒序
	SortDir *string `json:"sort_dir,omitempty"`

	// - 功能说明：根据id过滤
	Id *string `json:"id,omitempty"`

	// - 功能说明：根据name过滤
	Name *string `json:"name,omitempty"`

	// - 功能说明：根据size过滤
	Size *int32 `json:"size,omitempty"`

	// - 功能说明：根据status过滤
	Status *string `json:"status,omitempty"`

	// - 功能说明：根据type过滤
	Type *string `json:"type,omitempty"`

	// - 功能说明：根据description过滤
	Description *string `json:"description,omitempty"`

	// - 功能说明：根据public_border_group过滤
	PublicBorderGroup *string `json:"public_border_group,omitempty"`
}

func (o ShowPublicipPoolTypesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublicipPoolTypesRequest struct{}"
	}

	return strings.Join([]string{"ShowPublicipPoolTypesRequest", string(data)}, " ")
}
