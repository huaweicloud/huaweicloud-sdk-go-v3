package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQueuesResponse Response Object
type ShowQueuesResponse struct {

	// 队列信息列表
	Data           *[]interface{} `json:"data,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowQueuesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQueuesResponse struct{}"
	}

	return strings.Join([]string{"ShowQueuesResponse", string(data)}, " ")
}
