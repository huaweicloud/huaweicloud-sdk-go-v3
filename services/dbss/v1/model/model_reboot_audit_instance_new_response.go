package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RebootAuditInstanceNewResponse Response Object
type RebootAuditInstanceNewResponse struct {

	// 操作结果  - success: 成功  - failed: 失败
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RebootAuditInstanceNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebootAuditInstanceNewResponse struct{}"
	}

	return strings.Join([]string{"RebootAuditInstanceNewResponse", string(data)}, " ")
}
