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
type DeleteTopicRequest struct {
	TopicUrn string `json:"topic_urn"`
}

func (o DeleteTopicRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteTopicRequest", string(data)}, " ")
}
