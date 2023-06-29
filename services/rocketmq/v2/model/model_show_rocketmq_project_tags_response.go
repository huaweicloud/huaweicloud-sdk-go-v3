package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRocketmqProjectTagsResponse Response Object
type ShowRocketmqProjectTagsResponse struct {

	// 标签列表
	Tags           *[]TagMultyValueEntity `json:"tags,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowRocketmqProjectTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRocketmqProjectTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowRocketmqProjectTagsResponse", string(data)}, " ")
}
