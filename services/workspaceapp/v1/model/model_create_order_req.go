package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrderReq 创建订单请求体。
type CreateOrderReq struct {

	// 包周期资源对象。
	Resources []Resource `json:"resources"`

	// createApps【添加云应用】。
	Type string `json:"type"`
}

func (o CreateOrderReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrderReq struct{}"
	}

	return strings.Join([]string{"CreateOrderReq", string(data)}, " ")
}
