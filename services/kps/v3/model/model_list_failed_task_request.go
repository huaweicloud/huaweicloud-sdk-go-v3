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
type ListFailedTaskRequest struct {
}

func (o ListFailedTaskRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListFailedTaskRequest", string(data)}, " ")
}
