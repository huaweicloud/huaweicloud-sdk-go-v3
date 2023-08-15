package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTopicResponse Response Object
type DeleteTopicResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTopicResponse struct{}"
	}

	return strings.Join([]string{"DeleteTopicResponse", string(data)}, " ")
}
