package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetSensitiveMaskRuleSwitchResponse Response Object
type SetSensitiveMaskRuleSwitchResponse struct {

	// 状态  - SUCCESS: 成功  - FAILED: 失败
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetSensitiveMaskRuleSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetSensitiveMaskRuleSwitchResponse struct{}"
	}

	return strings.Join([]string{"SetSensitiveMaskRuleSwitchResponse", string(data)}, " ")
}
