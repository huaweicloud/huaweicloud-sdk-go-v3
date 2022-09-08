package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新路由请求体
type UpdateRoutesRequestBody struct {

	// 待更新的路由信息
	Routes []RouteOption `json:"routes"`
}

func (o UpdateRoutesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRoutesRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateRoutesRequestBody", string(data)}, " ")
}
