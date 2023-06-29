package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAuditLogResponse Response Object
type UpdateAuditLogResponse struct {

	// 开启/关闭全量SQL操作结果。
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAuditLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAuditLogResponse struct{}"
	}

	return strings.Join([]string{"UpdateAuditLogResponse", string(data)}, " ")
}
