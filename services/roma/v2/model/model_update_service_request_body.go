package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateServiceRequestBody struct {
	// 服务名称，支持中文,英文大小写，数字，下划线和中划线,长度2-64

	ServiceName string `json:"service_name"`
	// 服务描述，长度0-200

	Description *string `json:"description,omitempty"`
	// 服务状态 0-启用 1-停用

	Status int32 `json:"status"`
}

func (o UpdateServiceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServiceRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateServiceRequestBody", string(data)}, " ")
}
