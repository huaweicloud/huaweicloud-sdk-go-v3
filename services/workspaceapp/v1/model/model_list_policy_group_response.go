package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPolicyGroupResponse Response Object
type ListPolicyGroupResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 策略组列表，返回列表条目数量上限为分页的最大上限值。
	Items          *[]PolicyGroup `json:"items,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListPolicyGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyGroupResponse struct{}"
	}

	return strings.Join([]string{"ListPolicyGroupResponse", string(data)}, " ")
}
