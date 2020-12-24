/*
 * SMN
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
type ListVersionsRequest struct {
}

func (o ListVersionsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListVersionsRequest", string(data)}, " ")
}
