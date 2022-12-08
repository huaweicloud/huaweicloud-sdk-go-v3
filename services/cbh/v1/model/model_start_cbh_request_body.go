package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 启动CBH实例请求对象
type StartCbhRequestBody struct {

	// 实例的server id
	InstanceId string `json:"instance_id"`
}

func (o StartCbhRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartCbhRequestBody struct{}"
	}

	return strings.Join([]string{"StartCbhRequestBody", string(data)}, " ")
}
