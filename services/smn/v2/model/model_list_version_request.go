/*
 * smn
 *
 * SMN Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListVersionRequest struct {
	ApiVersion string `json:"api_version"`
}

func (o ListVersionRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListVersionRequest", string(data)}, " ")
}
