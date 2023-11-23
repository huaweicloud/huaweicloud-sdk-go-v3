package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRules 灰度发布规则
type CreateRules struct {

	// 优先级，数字越大，优先级越高。
	Precedence *int32 `json:"precedence,omitempty"`

	Match *CreateMatch `json:"match,omitempty"`

	// 路由规则列表。
	Route *[]CreateRoute `json:"route,omitempty"`
}

func (o CreateRules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRules struct{}"
	}

	return strings.Join([]string{"CreateRules", string(data)}, " ")
}
