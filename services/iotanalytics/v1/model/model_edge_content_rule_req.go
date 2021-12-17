package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 在Edge实例中要配置转发规则推送数据的资源空间和产品列表
type EdgeContentRuleReq struct {
	// Edge中的资源空间Id

	AppId string `json:"app_id"`
	// Edge中某资源空间Id下的产品列表

	Products []string `json:"products"`
}

func (o EdgeContentRuleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgeContentRuleReq struct{}"
	}

	return strings.Join([]string{"EdgeContentRuleReq", string(data)}, " ")
}
