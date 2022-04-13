package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
