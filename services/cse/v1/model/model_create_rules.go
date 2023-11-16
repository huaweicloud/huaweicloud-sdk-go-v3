package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRules 灰度发布规则
type CreateRules struct {

	// precedence
	Precedence *int32 `json:"precedence,omitempty"`

	Match *CreateMatch `json:"match,omitempty"`

	// route
	Route *[]CreateRoute `json:"route,omitempty"`
}

func (o CreateRules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRules struct{}"
	}

	return strings.Join([]string{"CreateRules", string(data)}, " ")
}
