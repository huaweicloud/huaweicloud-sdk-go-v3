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
type DeleteTopicAttributesRequest struct {
	TopicUrn string `json:"topic_urn"`
}

func (o DeleteTopicAttributesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteTopicAttributesRequest", string(data)}, " ")
}
