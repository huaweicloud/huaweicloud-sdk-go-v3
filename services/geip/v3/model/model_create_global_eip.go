package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGlobalEip 创建全域弹性公网IP响应体
type CreateGlobalEip struct {

	// ID
	Id string `json:"id"`

	// 资源名称
	Name string `json:"name"`
}

func (o CreateGlobalEip) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGlobalEip struct{}"
	}

	return strings.Join([]string{"CreateGlobalEip", string(data)}, " ")
}
