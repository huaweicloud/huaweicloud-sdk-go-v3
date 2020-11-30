/*
 * BMS
 *
 * BMS Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowSpecifiedVersionRequest struct {
	ApiVersion string `json:"api_version"`
}

func (o ShowSpecifiedVersionRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowSpecifiedVersionRequest", string(data)}, " ")
}
