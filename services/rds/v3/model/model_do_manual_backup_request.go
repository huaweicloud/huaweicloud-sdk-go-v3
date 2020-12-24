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

// Request Object
type DoManualBackupRequest struct {
	XLanguage *string                    `json:"X-Language,omitempty"`
	Body      *DoManualBackupRequestBody `json:"body,omitempty"`
}

func (o DoManualBackupRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DoManualBackupRequest", string(data)}, " ")
}
