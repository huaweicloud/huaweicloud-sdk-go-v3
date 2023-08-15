package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateProtectedInstanceRequestBody 创建保护实例请求体
type CreateProtectedInstanceRequestBody struct {
	ProtectedInstance *CreateProtectedInstanceRequestParams `json:"protected_instance"`
}

func (o CreateProtectedInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProtectedInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateProtectedInstanceRequestBody", string(data)}, " ")
}
