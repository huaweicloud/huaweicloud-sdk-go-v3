package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImportTrafficRequestBody struct {

	// 手机路由类型 direct：默认路由 routing：路由到编码容器
	TrafficType string `json:"traffic_type"`

	// 手机id列表 一次调用最多支持100个
	PhoneIds []string `json:"phone_ids"`
}

func (o ImportTrafficRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportTrafficRequestBody struct{}"
	}

	return strings.Join([]string{"ImportTrafficRequestBody", string(data)}, " ")
}
