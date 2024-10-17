package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadAuditAgentResponse Response Object
type DownloadAuditAgentResponse struct {

	// 响应状态
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DownloadAuditAgentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadAuditAgentResponse struct{}"
	}

	return strings.Join([]string{"DownloadAuditAgentResponse", string(data)}, " ")
}
