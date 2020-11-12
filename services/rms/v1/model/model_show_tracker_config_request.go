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
type ShowTrackerConfigRequest struct {
}

func (o ShowTrackerConfigRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowTrackerConfigRequest", string(data)}, " ")
}
