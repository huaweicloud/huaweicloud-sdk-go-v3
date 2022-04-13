package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListCustomLineRequest struct {
	// 解析线路ID。

	LineId *string `json:"line_id,omitempty"`
	// 解析线路名称。

	Name *string `json:"name,omitempty"`
	// 每页返回的资源个数。取值范围为0~100。

	Limit *int32 `json:"limit,omitempty"`
	// 分页查询起始偏移量，表示从偏移量的下一个资源开始查询。

	Offset *int32 `json:"offset,omitempty"`
	// 是否查询详细信息。  取值范围：  true：是，查询详细信息。 false：否，不查询详细信息。 默认为true。

	ShowDetail *bool `json:"show_detail,omitempty"`
}

func (o ListCustomLineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomLineRequest struct{}"
	}

	return strings.Join([]string{"ListCustomLineRequest", string(data)}, " ")
}
