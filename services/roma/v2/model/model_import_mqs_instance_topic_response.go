package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ImportMqsInstanceTopicResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ImportMqsInstanceTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportMqsInstanceTopicResponse struct{}"
	}

	return strings.Join([]string{"ImportMqsInstanceTopicResponse", string(data)}, " ")
}
