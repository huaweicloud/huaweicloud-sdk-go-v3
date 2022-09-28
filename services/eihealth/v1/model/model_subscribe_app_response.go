package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SubscribeAppResponse struct {

	// 应用id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SubscribeAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribeAppResponse struct{}"
	}

	return strings.Join([]string{"SubscribeAppResponse", string(data)}, " ")
}
