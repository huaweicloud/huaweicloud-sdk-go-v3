package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopCbhRequestBody 关闭云堡垒机实例请求对象。
type StopCbhRequestBody struct {

	// 云堡垒机实例ID，使用UUID格式。
	InstanceId string `json:"instance_id"`
}

func (o StopCbhRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopCbhRequestBody struct{}"
	}

	return strings.Join([]string{"StopCbhRequestBody", string(data)}, " ")
}
