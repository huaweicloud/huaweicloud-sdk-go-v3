package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSensitiveMaskRuleResponse Response Object
type CreateSensitiveMaskRuleResponse struct {

	// 状态  - SUCCESS: 成功  - FAILED: 失败
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSensitiveMaskRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSensitiveMaskRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateSensitiveMaskRuleResponse", string(data)}, " ")
}
