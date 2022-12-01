package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateOrderFormResponse struct {

	// 技能列表
	Data *[]SkillInfo `json:"data,omitempty"`

	// 总数量
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateOrderFormResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrderFormResponse struct{}"
	}

	return strings.Join([]string{"CreateOrderFormResponse", string(data)}, " ")
}
