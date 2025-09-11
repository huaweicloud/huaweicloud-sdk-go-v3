package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSensitiveMaskRuleResponse Response Object
type UpdateSensitiveMaskRuleResponse struct {

	// 状态  - SUCCESS:成功  - FAILED:失败
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateSensitiveMaskRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSensitiveMaskRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateSensitiveMaskRuleResponse", string(data)}, " ")
}
