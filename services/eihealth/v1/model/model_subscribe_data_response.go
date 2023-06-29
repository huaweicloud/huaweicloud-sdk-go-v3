package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribeDataResponse Response Object
type SubscribeDataResponse struct {

	// 数据作业ID
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SubscribeDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribeDataResponse struct{}"
	}

	return strings.Join([]string{"SubscribeDataResponse", string(data)}, " ")
}
