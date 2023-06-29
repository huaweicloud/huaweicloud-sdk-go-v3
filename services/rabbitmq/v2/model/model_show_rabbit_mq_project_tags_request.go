package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRabbitMqProjectTagsRequest Request Object
type ShowRabbitMqProjectTagsRequest struct {
}

func (o ShowRabbitMqProjectTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRabbitMqProjectTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowRabbitMqProjectTagsRequest", string(data)}, " ")
}
