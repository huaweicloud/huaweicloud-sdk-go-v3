package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPolicyTemplateResponse Response Object
type ListPolicyTemplateResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 策略组列表，返回列表条目数量上限为分页的最大上限值。
	Items          *[]PolicyGroup `json:"items,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListPolicyTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyTemplateResponse struct{}"
	}

	return strings.Join([]string{"ListPolicyTemplateResponse", string(data)}, " ")
}
