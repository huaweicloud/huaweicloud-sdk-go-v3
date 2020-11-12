/*
 * RMS
 *
 * Resource Manager Api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateTrackerConfigResponse struct {
}

func (o CreateTrackerConfigResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateTrackerConfigResponse", string(data)}, " ")
}
