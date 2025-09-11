package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditSensitiveMasksNewResponse Response Object
type ListAuditSensitiveMasksNewResponse struct {

	// 规则列表
	Rules *[]SensitiveMaskResponseRules `json:"rules,omitempty"`

	// 总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAuditSensitiveMasksNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditSensitiveMasksNewResponse struct{}"
	}

	return strings.Join([]string{"ListAuditSensitiveMasksNewResponse", string(data)}, " ")
}
