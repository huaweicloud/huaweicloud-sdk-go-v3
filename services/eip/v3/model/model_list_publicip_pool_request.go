package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPublicipPoolRequest struct {

	// 分页查询起始的资源ID，为空时为查询第一页
	Marker *string `json:"marker,omitempty" xml:"marker"`

	// 每页返回的个数取值范围：0~[2000]，其中2000为局点差异项，具体取值由局点决定
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 显示，形式为\"fields=id&fields=name&...\"  支持字段：id/name/size/used/project_id/status/billing_info/created_at/updated_at/type/shared/is_common/description/tags/enterprise_project_id/allow_share_bandwidth_types/public_border_group
	Fields *string `json:"fields,omitempty" xml:"fields"`

	// 排序，形式为\"sort_key=id&sort_dir=asc\"  支持字段：id/name/created_at/updated_at/public_border_group
	SortKey *string `json:"sort_key,omitempty" xml:"sort_key"`

	// 排序方向  取值范围：asc、desc
	SortDir *string `json:"sort_dir,omitempty" xml:"sort_dir"`

	// 根据id过滤
	Id *string `json:"id,omitempty" xml:"id"`

	// 根据name过滤
	Name *string `json:"name,omitempty" xml:"name"`

	// 根据size过滤
	Size *int32 `json:"size,omitempty" xml:"size"`

	// 根据status过滤
	Status *string `json:"status,omitempty" xml:"status"`

	// 根据type过滤
	Type *string `json:"type,omitempty" xml:"type"`

	// 根据description过滤
	Description *string `json:"description,omitempty" xml:"description"`

	// 根据public_border_group过滤
	PublicBorderGroup *string `json:"public_border_group,omitempty" xml:"public_border_group"`
}

func (o ListPublicipPoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublicipPoolRequest struct{}"
	}

	return strings.Join([]string{"ListPublicipPoolRequest", string(data)}, " ")
}
