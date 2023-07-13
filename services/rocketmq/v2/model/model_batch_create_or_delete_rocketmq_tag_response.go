package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateOrDeleteRocketmqTagResponse Response Object
type BatchCreateOrDeleteRocketmqTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateOrDeleteRocketmqTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateOrDeleteRocketmqTagResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateOrDeleteRocketmqTagResponse", string(data)}, " ")
}
