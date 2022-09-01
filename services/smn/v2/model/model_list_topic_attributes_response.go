package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListTopicAttributesResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty" xml:"request_id"`

	Attributes     *TopicAttribute `json:"attributes,omitempty" xml:"attributes"`
	HttpStatusCode int             `json:"-"`
}

func (o ListTopicAttributesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopicAttributesResponse struct{}"
	}

	return strings.Join([]string{"ListTopicAttributesResponse", string(data)}, " ")
}
