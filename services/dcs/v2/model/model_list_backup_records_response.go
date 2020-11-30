/*
 * DCS
 *
 * DCS V2版本API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListBackupRecordsResponse struct {
	// 返回记录数。
	TotalNum *int32 `json:"total_num,omitempty"`
	// 备份信息的详情数组。
	BackupRecordResponse *[]BackupRecordResponse `json:"backup_record_response,omitempty"`
	HttpStatusCode       int                     `json:"-"`
}

func (o ListBackupRecordsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListBackupRecordsResponse", string(data)}, " ")
}
