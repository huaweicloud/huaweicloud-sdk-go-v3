package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTopicAttributeByNameResponse Response Object
type DeleteTopicAttributeByNameResponse struct {

	// 请求的唯一标识ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteTopicAttributeByNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTopicAttributeByNameResponse struct{}"
	}

	return strings.Join([]string{"DeleteTopicAttributeByNameResponse", string(data)}, " ")
}
