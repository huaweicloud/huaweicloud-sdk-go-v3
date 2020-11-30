/*
 * DMS
 *
 * DMS Document API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowQueueTagsResponse struct {
	// 标签列表
	Tags           *[]BatchCreateOrDeleteTagReqTags `json:"tags,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o ShowQueueTagsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowQueueTagsResponse", string(data)}, " ")
}
