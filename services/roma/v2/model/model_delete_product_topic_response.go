package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
