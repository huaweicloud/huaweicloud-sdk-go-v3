package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTopicResponse Response Object
type UpdateTopicResponse struct {

	// 请求的唯一标识ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTopicResponse struct{}"
	}

	return strings.Join([]string{"UpdateTopicResponse", string(data)}, " ")
}
