/*
 * smn
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
type ListTopicAttributesRequest struct {
	TopicUrn string `json:"topic_urn"`
	Name     string `json:"name"`
}

func (o ListTopicAttributesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListTopicAttributesRequest", string(data)}, " ")
}
