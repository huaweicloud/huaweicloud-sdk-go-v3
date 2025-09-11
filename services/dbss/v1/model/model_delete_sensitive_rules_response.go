package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSensitiveRulesResponse Response Object
type DeleteSensitiveRulesResponse struct {

	// 状态  - SUCCESS:成功  - FAILED:失败
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteSensitiveRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSensitiveRulesResponse struct{}"
	}

	return strings.Join([]string{"DeleteSensitiveRulesResponse", string(data)}, " ")
}
