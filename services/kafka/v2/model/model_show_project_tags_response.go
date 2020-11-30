/*
 * Kafka
 *
 * Kafka Document API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowProjectTagsResponse struct {
	// 标签列表
	Tags           *[]ShowProjectTagsRespTags `json:"tags,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ShowProjectTagsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowProjectTagsResponse", string(data)}, " ")
}
