package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateOrDeleteRabbitMqTagResponse Response Object
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
