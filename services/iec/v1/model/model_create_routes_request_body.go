package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建路由请求体
type CreateRoutesRequestBody struct {

	// 待创建的路由列表
	Routes []RouteOption `json:"routes"`
}

func (o CreateRoutesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRoutesRequestBody struct{}"
	}

	return strings.Join([]string{"CreateRoutesRequestBody", string(data)}, " ")
}
