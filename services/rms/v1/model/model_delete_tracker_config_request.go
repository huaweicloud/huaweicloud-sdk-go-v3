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

// Request Object
type DeleteTrackerConfigRequest struct {
}

func (o DeleteTrackerConfigRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteTrackerConfigRequest", string(data)}, " ")
}
