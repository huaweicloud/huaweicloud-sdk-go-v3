package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDefaultTopicResponse Response Object
type UpdateDefaultTopicResponse struct {

	// 返回消息
	Msg *string `json:"msg,omitempty"`

	// 返回状态，如'200','400'
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDefaultTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDefaultTopicResponse struct{}"
	}

	return strings.Join([]string{"UpdateDefaultTopicResponse", string(data)}, " ")
}
