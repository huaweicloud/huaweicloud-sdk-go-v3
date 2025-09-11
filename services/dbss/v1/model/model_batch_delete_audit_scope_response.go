package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteAuditScopeResponse Response Object
type BatchDeleteAuditScopeResponse struct {

	// 状态  - SUCCESS:成功  - FAILED:失败
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchDeleteAuditScopeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAuditScopeResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteAuditScopeResponse", string(data)}, " ")
}
