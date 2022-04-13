package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchCreateOrDeleteRabbitMqTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateOrDeleteRabbitMqTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateOrDeleteRabbitMqTagResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateOrDeleteRabbitMqTagResponse", string(data)}, " ")
}
