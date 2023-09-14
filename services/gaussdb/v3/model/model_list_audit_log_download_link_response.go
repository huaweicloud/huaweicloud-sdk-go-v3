package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditLogDownloadLinkResponse Response Object
type ListAuditLogDownloadLinkResponse struct {

	// 获取到的全量SQL文件信息。
	Files          *[]FileInfo `json:"files,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListAuditLogDownloadLinkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditLogDownloadLinkResponse struct{}"
	}

	return strings.Join([]string{"ListAuditLogDownloadLinkResponse", string(data)}, " ")
}
