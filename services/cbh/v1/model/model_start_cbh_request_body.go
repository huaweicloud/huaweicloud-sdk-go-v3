package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 启动云堡垒机实例请求对象。
type StartCbhRequestBody struct {

	// 云堡垒机实例的ID，使用UUID格式。
	InstanceId string `json:"instance_id"`
}

func (o StartCbhRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartCbhRequestBody struct{}"
	}

	return strings.Join([]string{"StartCbhRequestBody", string(data)}, " ")
}
