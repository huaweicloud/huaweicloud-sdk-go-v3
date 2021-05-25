package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListTopicDetailsRequest struct {
	// Topic的唯一的资源标识，可通过[查询主题列表](https://support.huaweicloud.com/api-smn/smn_api_51004.html)获取该标识。

	TopicUrn string `json:"topic_urn"`
}

func (o ListTopicDetailsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListTopicDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListTopicDetailsRequest", string(data)}, " ")
}
