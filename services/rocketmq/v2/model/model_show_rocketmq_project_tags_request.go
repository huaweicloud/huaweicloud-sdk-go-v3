package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRocketmqProjectTagsRequest Request Object
type ShowRocketmqProjectTagsRequest struct {
}

func (o ShowRocketmqProjectTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRocketmqProjectTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowRocketmqProjectTagsRequest", string(data)}, " ")
}
