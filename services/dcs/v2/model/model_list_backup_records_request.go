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

// Request Object
type ListBackupRecordsRequest struct {
	InstanceId string  `json:"instance_id"`
	BeginTime  *string `json:"begin_time,omitempty"`
	EndTime    *string `json:"end_time,omitempty"`
	Limit      *int32  `json:"limit,omitempty"`
	Offset     *int32  `json:"offset,omitempty"`
}

func (o ListBackupRecordsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListBackupRecordsRequest", string(data)}, " ")
}
