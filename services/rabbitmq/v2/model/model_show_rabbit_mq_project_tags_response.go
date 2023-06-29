package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRabbitMqProjectTagsResponse Response Object
type ShowRabbitMqProjectTagsResponse struct {

	// 标签列表
	Tags           *[]TagMultyValueEntity `json:"tags,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowRabbitMqProjectTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRabbitMqProjectTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowRabbitMqProjectTagsResponse", string(data)}, " ")
}
