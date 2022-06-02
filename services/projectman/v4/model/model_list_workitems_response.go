package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListWorkitemsResponse struct {

	// 工作项
	WorkItems *[]Workitems `json:"work_items,omitempty"`

	// 总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListWorkitemsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkitemsResponse struct{}"
	}

	return strings.Join([]string{"ListWorkitemsResponse", string(data)}, " ")
}
