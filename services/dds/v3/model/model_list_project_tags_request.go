/*
 * DDS
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
}

func (o ListProjectTagsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListProjectTagsRequest", string(data)}, " ")
}
