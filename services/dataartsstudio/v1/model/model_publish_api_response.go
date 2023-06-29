package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishApiResponse Response Object
type PublishApiResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o PublishApiResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishApiResponse struct{}"
	}

	return strings.Join([]string{"PublishApiResponse", string(data)}, " ")
}
