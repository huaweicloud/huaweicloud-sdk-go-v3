package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuditLogDownloadLinkResponse Response Object
type ShowAuditLogDownloadLinkResponse struct {

	// **参数解释**：  审计日志下载链接列表，链接有效时间5分钟。  **取值范围**：  不涉及。
	Links          *[]string `json:"links,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowAuditLogDownloadLinkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuditLogDownloadLinkResponse struct{}"
	}

	return strings.Join([]string{"ShowAuditLogDownloadLinkResponse", string(data)}, " ")
}
