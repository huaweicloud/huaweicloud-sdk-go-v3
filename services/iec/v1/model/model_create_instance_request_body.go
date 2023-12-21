package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceRequestBody 创建边缘实例请求对象
type CreateInstanceRequestBody struct {
	Server *CreateInstanceOption `json:"server"`
}

func (o CreateInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateInstanceRequestBody", string(data)}, " ")
}
