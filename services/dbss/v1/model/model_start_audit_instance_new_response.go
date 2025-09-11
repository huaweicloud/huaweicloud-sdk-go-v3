package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartAuditInstanceNewResponse Response Object
type StartAuditInstanceNewResponse struct {

	// 操作结果  - success: 成功  - failed: 失败
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartAuditInstanceNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartAuditInstanceNewResponse struct{}"
	}

	return strings.Join([]string{"StartAuditInstanceNewResponse", string(data)}, " ")
}
