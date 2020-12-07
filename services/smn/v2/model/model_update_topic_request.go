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
type UpdateTopicRequest struct {
	TopicUrn string                  `json:"topic_urn"`
	Body     *UpdateTopicRequestBody `json:"body,omitempty"`
}

func (o UpdateTopicRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateTopicRequest", string(data)}, " ")
}
