package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGlobalEipRequestBodyGlobalEip 更新全域弹性公网IP请求体
type UpdateGlobalEipRequestBodyGlobalEip struct {

	// 资源名称
	Name *string `json:"name,omitempty"`

	// 用户自定义的资源描述
	Description *string `json:"description,omitempty"`
}

func (o UpdateGlobalEipRequestBodyGlobalEip) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGlobalEipRequestBodyGlobalEip struct{}"
	}

	return strings.Join([]string{"UpdateGlobalEipRequestBodyGlobalEip", string(data)}, " ")
}
