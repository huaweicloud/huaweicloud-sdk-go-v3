package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteListenerTagsRequest struct {

	// 监听器ID
	ListenerId string `json:"listener_id"`

	// 待删除标签的key值。
	Key string `json:"key"`
}

func (o DeleteListenerTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteListenerTagsRequest struct{}"
	}

	return strings.Join([]string{"DeleteListenerTagsRequest", string(data)}, " ")
}
