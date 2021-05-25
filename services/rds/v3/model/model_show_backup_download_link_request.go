package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowBackupDownloadLinkRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 备份ID。

	BackupId string `json:"backup_id"`
}

func (o ShowBackupDownloadLinkRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowBackupDownloadLinkRequest struct{}"
	}

	return strings.Join([]string{"ShowBackupDownloadLinkRequest", string(data)}, " ")
}
