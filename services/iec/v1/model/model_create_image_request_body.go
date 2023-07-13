package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateImageRequestBody 创建边缘私有镜像参数
type CreateImageRequestBody struct {

	// 边缘私有镜像名称。
	Name string `json:"name"`

	// 边缘实例名称。
	InstanceId string `json:"instance_id"`
}

func (o CreateImageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageRequestBody struct{}"
	}

	return strings.Join([]string{"CreateImageRequestBody", string(data)}, " ")
}
