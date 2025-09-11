package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAuditScopeResponse Response Object
type DeleteAuditScopeResponse struct {

	// 状态  - SUCCESS:成功  - FAILED:失败
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAuditScopeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAuditScopeResponse struct{}"
	}

	return strings.Join([]string{"DeleteAuditScopeResponse", string(data)}, " ")
}
