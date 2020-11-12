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
type DeleteTrackerConfigResponse struct {
}

func (o DeleteTrackerConfigResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteTrackerConfigResponse", string(data)}, " ")
}
