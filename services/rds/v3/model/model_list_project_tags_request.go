/*
 * RDS
 *
 * API v3
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListProjectTagsRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ListProjectTagsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListProjectTagsRequest", string(data)}, " ")
}
