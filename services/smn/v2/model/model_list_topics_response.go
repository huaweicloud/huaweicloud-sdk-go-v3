package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListTopicsResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty" xml:"request_id"`

	// 返回的Topic个数。该参数不受offset和limit影响，即返回的是您账户下所有的Topic个数。
	TopicCount *int32 `json:"topic_count,omitempty" xml:"topic_count"`

	// Topic结构体数组。
	Topics         *[]ListTopicsItem `json:"topics,omitempty" xml:"topics"`
	HttpStatusCode int               `json:"-"`
}

func (o ListTopicsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopicsResponse struct{}"
	}

	return strings.Join([]string{"ListTopicsResponse", string(data)}, " ")
}
