package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLineGroupsRequest Request Object
type ListLineGroupsRequest struct {

	// 线路分组ID。 模糊匹配。
	LineId *string `json:"line_id,omitempty"`

	// 线路分组名称。 模糊匹配。
	Name *string `json:"name,omitempty"`

	// 分页查询时配置每页返回的资源个数。  取值范围：0~500  取值一般为10，20，50。默认值为500。
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询起始偏移量，表示从偏移量的下一个资源开始查询。  取值范围：0~2147483647  默认值为0。  当设置marker不为空时，以marker为分页起始标识，offset不生效。。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListLineGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLineGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListLineGroupsRequest", string(data)}, " ")
}
