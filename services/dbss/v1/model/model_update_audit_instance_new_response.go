package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAuditInstanceNewResponse Response Object
type UpdateAuditInstanceNewResponse struct {

	// 操作结果  - success: 成功  - failed: 失败
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAuditInstanceNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAuditInstanceNewResponse struct{}"
	}

	return strings.Join([]string{"UpdateAuditInstanceNewResponse", string(data)}, " ")
}
