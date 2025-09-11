package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadAuditAgentNewResponse Response Object
type DownloadAuditAgentNewResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DownloadAuditAgentNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadAuditAgentNewResponse struct{}"
	}

	return strings.Join([]string{"DownloadAuditAgentNewResponse", string(data)}, " ")
}
