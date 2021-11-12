package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 保护实例变更规格请求体
type ResizeProtectedInstanceRequestBody struct {
	Resize *ResizeProtectedInstanceRequestParams `json:"resize"`
}

func (o ResizeProtectedInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeProtectedInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"ResizeProtectedInstanceRequestBody", string(data)}, " ")
}
