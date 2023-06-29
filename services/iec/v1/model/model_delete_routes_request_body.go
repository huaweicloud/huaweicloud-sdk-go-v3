package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRoutesRequestBody 删除路由请求体
type DeleteRoutesRequestBody struct {

	// 待删除的路由信息
	Routes []DeleteRouteOption `json:"routes"`
}

func (o DeleteRoutesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRoutesRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteRoutesRequestBody", string(data)}, " ")
}
