package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAccessClientRequestBody 申请接入服务的请求信息
type UpdateAccessClientRequestBody struct {

	// 客户端名称
	Name string `json:"name"`
}

func (o UpdateAccessClientRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAccessClientRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAccessClientRequestBody", string(data)}, " ")
}
