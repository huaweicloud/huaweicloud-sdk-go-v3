package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceRequestBody 更新边缘实例请求体
type UpdateInstanceRequestBody struct {
	Server *UpdateInstanceOption `json:"server"`
}

func (o UpdateInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateInstanceRequestBody", string(data)}, " ")
}
