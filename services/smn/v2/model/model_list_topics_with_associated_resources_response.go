package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTopicsWithAssociatedResourcesResponse Response Object
type ListTopicsWithAssociatedResourcesResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 返回的Topic个数。该参数不受offset和limit影响，即返回的是您账户下所有的Topic个数。
	TopicCount *int32 `json:"topic_count,omitempty"`

	// Topic及关联资源信息结构体数组。
	Topics         *[]ListTopicsWithAssociatedResourcesItem `json:"topics,omitempty"`
	HttpStatusCode int                                      `json:"-"`
}

func (o ListTopicsWithAssociatedResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopicsWithAssociatedResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListTopicsWithAssociatedResourcesResponse", string(data)}, " ")
}
