package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppGroupResponse Response Object
type ListAppGroupResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 应用组列表。
	Items          *[]AppGroup `json:"items,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListAppGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppGroupResponse struct{}"
	}

	return strings.Join([]string{"ListAppGroupResponse", string(data)}, " ")
}
