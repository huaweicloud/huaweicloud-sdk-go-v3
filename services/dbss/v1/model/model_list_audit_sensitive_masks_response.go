package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditSensitiveMasksResponse Response Object
type ListAuditSensitiveMasksResponse struct {

	// 规则列表
	Rules *[]SensitiveMaskResponseRules `json:"rules,omitempty"`

	// 总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAuditSensitiveMasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditSensitiveMasksResponse struct{}"
	}

	return strings.Join([]string{"ListAuditSensitiveMasksResponse", string(data)}, " ")
}
