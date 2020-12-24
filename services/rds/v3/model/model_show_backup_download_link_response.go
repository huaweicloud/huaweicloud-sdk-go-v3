/*
 * RDS
 *
 * API v3
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowBackupDownloadLinkResponse struct {
	// 备份文件信息。
	Files *[]GetBackupDownloadLinkFiles `json:"files,omitempty"`
	// OBS桶名。
	Bucket         *string `json:"bucket,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowBackupDownloadLinkResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowBackupDownloadLinkResponse", string(data)}, " ")
}
