package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRuleRequest Request Object
type ListRuleRequest struct {

	// 分页获取列表时，页的大小，默认为-1
	Limit *int32 `json:"limit,omitempty"`

	// 分页获取列表时，起始偏移量，默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 分页获取列表时，排序参数，支持 create_at 和 update_at
	OrderBy *string `json:"order_by,omitempty"`

	// 分页获取列表时，排序方向，支持 desc 和 asc
	Order *string `json:"order,omitempty"`
}

func (o ListRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRuleRequest struct{}"
	}

	return strings.Join([]string{"ListRuleRequest", string(data)}, " ")
}
