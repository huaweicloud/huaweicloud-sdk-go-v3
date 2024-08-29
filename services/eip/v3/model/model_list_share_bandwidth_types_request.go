package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListShareBandwidthTypesRequest Request Object
type ListShareBandwidthTypesRequest struct {

	// 形式为\\\"fields=id&fields=bandwidth_type&...\\\"，支持字段：id/bandwidth_type/name_en/name_zh/created_at/update_at/public_border_group/description
	Fields *[]string `json:"fields,omitempty"`

	// 支持带宽类型的id
	Id *string `json:"id,omitempty"`

	// 带宽支持类型
	BandwidthType *string `json:"bandwidth_type,omitempty"`

	// 带宽类型英文表述
	NameEn *string `json:"name_en,omitempty"`

	// 带宽类型中文表述
	NameZh *string `json:"name_zh,omitempty"`

	// 带宽类型所处位置，中心站点or边缘站点
	PublicBorderGroup *string `json:"public_border_group,omitempty"`

	// 排序，形式为\"sort_key=id&sort_dir=asc\"  支持字段：id/bandwidth_type/public_border_group
	SortKey *string `json:"sort_key,omitempty"`

	// 排序方向  取值范围：asc、desc
	SortDir *string `json:"sort_dir,omitempty"`

	// 每页返回的个数取值范围：0~[2000]，其中2000为局点差异项，具体取值由局点决定
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询起始的资源ID，为空时为查询第一页
	Marker *string `json:"marker,omitempty"`

	// 分页查询起始的资源序号
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListShareBandwidthTypesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListShareBandwidthTypesRequest struct{}"
	}

	return strings.Join([]string{"ListShareBandwidthTypesRequest", string(data)}, " ")
}
