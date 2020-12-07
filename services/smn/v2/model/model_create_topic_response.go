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

// Response Object
type CreateTopicResponse struct {
	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`
	// Topic的唯一的资源标识，可通过查询主题列表获取该标识。
	TopicUrn       *string `json:"topic_urn,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateTopicResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateTopicResponse", string(data)}, " ")
}
