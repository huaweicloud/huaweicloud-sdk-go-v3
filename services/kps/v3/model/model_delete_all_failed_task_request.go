/*
 * kps
 *
 * kps v3 版本API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteAllFailedTaskRequest struct {
}

func (o DeleteAllFailedTaskRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteAllFailedTaskRequest", string(data)}, " ")
}
