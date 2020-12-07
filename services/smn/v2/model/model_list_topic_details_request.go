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
type ListTopicDetailsRequest struct {
	TopicUrn string `json:"topic_urn"`
}

func (o ListTopicDetailsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListTopicDetailsRequest", string(data)}, " ")
}
