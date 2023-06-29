package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribeImageResponse Response Object
type SubscribeImageResponse struct {

	// 镜像id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SubscribeImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribeImageResponse struct{}"
	}

	return strings.Join([]string{"SubscribeImageResponse", string(data)}, " ")
}
