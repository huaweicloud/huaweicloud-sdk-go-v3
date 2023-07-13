package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserAssignment 授权
type UserAssignment struct {

	// 目标用户
	Attach string `json:"attach"`

	// 策略ID
	PolicyStatementId string `json:"policy_statement_id"`
}

func (o UserAssignment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserAssignment struct{}"
	}

	return strings.Join([]string{"UserAssignment", string(data)}, " ")
}
