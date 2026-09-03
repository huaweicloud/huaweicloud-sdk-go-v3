package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnsubscribeInstanceReportNewRequestBody 取消订阅实例报告请求体
type UnsubscribeInstanceReportNewRequestBody struct {

	// 订阅ID
	SubscribeId string `json:"subscribe_id"`
}

func (o UnsubscribeInstanceReportNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnsubscribeInstanceReportNewRequestBody struct{}"
	}

	return strings.Join([]string{"UnsubscribeInstanceReportNewRequestBody", string(data)}, " ")
}
