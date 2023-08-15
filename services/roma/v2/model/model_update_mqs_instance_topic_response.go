package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMqsInstanceTopicResponse Response Object
type UpdateMqsInstanceTopicResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateMqsInstanceTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMqsInstanceTopicResponse struct{}"
	}

	return strings.Join([]string{"UpdateMqsInstanceTopicResponse", string(data)}, " ")
}
