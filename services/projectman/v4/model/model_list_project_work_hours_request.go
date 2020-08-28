/*
 * ProjectMan
 *
 * devcloud projectman api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListProjectWorkHoursRequest struct {
	Body *ListProjectWorkHoursRequestBody `json:"body,omitempty"`
}

func (o ListProjectWorkHoursRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListProjectWorkHoursRequest", string(data)}, " ")
}
