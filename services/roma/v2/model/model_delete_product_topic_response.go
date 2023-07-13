package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteProductTopicResponse Response Object
type DeleteProductTopicResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteProductTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProductTopicResponse struct{}"
	}

	return strings.Join([]string{"DeleteProductTopicResponse", string(data)}, " ")
}
