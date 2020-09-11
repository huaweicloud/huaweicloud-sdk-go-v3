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

// Response Object
type DeleteAllFailedTaskResponse struct {
}

func (o DeleteAllFailedTaskResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteAllFailedTaskResponse", string(data)}, " ")
}
