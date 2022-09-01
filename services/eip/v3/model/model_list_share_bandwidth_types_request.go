package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListShareBandwidthTypesRequest struct {

	// 形式为\\\"fields=id&fields=bandwidth_type&...\\\"，支持字段：id/bandwidth_type/name_en/name_zh/created_at/update_at/public_border_group/description
	Fields *string `json:"fields,omitempty" xml:"fields"`

	// 支持带宽类型的id
	Id *string `json:"id,omitempty" xml:"id"`

	// 带宽支持类型
	BandwidthType *string `json:"bandwidth_type,omitempty" xml:"bandwidth_type"`

	// 带宽类型英文表述
	NameEn *string `json:"name_en,omitempty" xml:"name_en"`

	// 带宽类型中文表述
	NameZh *string `json:"name_zh,omitempty" xml:"name_zh"`

	// 带宽类型所处位置，中心站点or边缘站点
	PublicBorderGroup *string `json:"public_border_group,omitempty" xml:"public_border_group"`

	// 排序，形式为\"sort_key=id&sort_dir=asc\"  支持字段：id/bandwidth_type/public_border_group
	SortKey *string `json:"sort_key,omitempty" xml:"sort_key"`

	// 排序方向  取值范围：asc、desc
	SortDir *string `json:"sort_dir,omitempty" xml:"sort_dir"`

	// 每页返回的个数取值范围：0~[2000]，其中2000为局点差异项，具体取值由局点决定
	Limit *int32 `json:"limit,omitempty" xml:"limit"`
}

func (o ListShareBandwidthTypesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListShareBandwidthTypesRequest struct{}"
	}

	return strings.Join([]string{"ListShareBandwidthTypesRequest", string(data)}, " ")
}
