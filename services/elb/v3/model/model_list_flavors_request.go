package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFlavorsRequest Request Object
type ListFlavorsRequest struct {

	// 上一页最后一条记录的ID。  使用说明： - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。
	Marker *string `json:"marker,omitempty"`

	// 每页返回的个数。
	Limit *int32 `json:"limit,omitempty"`

	// 是否反向查询。  取值： - true：查询上一页。 - false：查询下一页，默认。  使用说明： - 必须与limit一起使用。 - 当page_reverse=true时，若要查询上一页，marker取值为当前页返回值的previous_marker。
	PageReverse *bool `json:"page_reverse,omitempty"`

	// 规格ID。  支持多值查询，查询条件格式：*id=xxx&id=xxx*。
	Id *[]string `json:"id,omitempty"`

	// 规格名称。   支持多值查询，查询条件格式：*name=xxx&name=xxx*。
	Name *[]string `json:"name,omitempty"`

	// 规格类别。  取值： - L4和L7 表示四层和七层flavor。 - L4_elastic和L7_elastic 表示弹性扩缩容实例的下限规格。 - L4_elastic_max和L7_elastic_max 表示弹性扩缩容实例的上限规格。  支持多值查询，查询条件格式：*type=xxx&type=xxx*。
	Type *[]string `json:"type,omitempty"`

	// 是否查询公共规格。true表示公共规格，所有租户可见。false表示私有规格，为当前租户所有。
	Shared *bool `json:"shared,omitempty"`
}

func (o ListFlavorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ListFlavorsRequest", string(data)}, " ")
}
