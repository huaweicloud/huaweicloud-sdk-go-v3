package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKafkaProjectTagsResponse Response Object
type ShowKafkaProjectTagsResponse struct {

	// **参数解释**： 标签列表。
	Tags           *[]TagMultyValueEntity `json:"tags,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowKafkaProjectTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKafkaProjectTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowKafkaProjectTagsResponse", string(data)}, " ")
}
